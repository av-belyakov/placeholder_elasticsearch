package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	simplelogger "github.com/av-belyakov/simplelogger"

	"placeholder_elasticsearch/confighandler"
	"placeholder_elasticsearch/coremodule"
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/elasticsearchinteractions"
	"placeholder_elasticsearch/eventenrichmentmodule"
	"placeholder_elasticsearch/internal/versionandname"
	"placeholder_elasticsearch/memorytemporarystorage"
	"placeholder_elasticsearch/mongodbinteractions"
	"placeholder_elasticsearch/natsinteractions"
	rules "placeholder_elasticsearch/rulesinteraction"
	"placeholder_elasticsearch/supportingfunctions"
	"placeholder_elasticsearch/zabbixinteractions"
)

const ROOT_DIR = "placeholder_elasticsearch"

func getLoggerSettings(cls []confighandler.LogSet) []simplelogger.Options {
	loggerConf := make([]simplelogger.Options, 0, len(cls))

	for _, v := range cls {
		loggerConf = append(loggerConf, simplelogger.Options{
			MsgTypeName:     v.MsgTypeName,
			WritingToFile:   v.WritingFile,
			PathDirectory:   v.PathDirectory,
			WritingToStdout: v.WritingStdout,
			MaxFileSize:     v.MaxFileSize,
		})
	}

	return loggerConf
}

// loggingHandler обработчик логов
func loggingHandler(
	channelZabbix chan<- zabbixinteractions.MessageSettings,
	sl *simplelogger.SimpleLoggerSettings,
	logging <-chan datamodels.MessageLogging) {
	for msg := range logging {
		_ = sl.WriteLoggingData(msg.MsgData, msg.MsgType)

		if msg.MsgType == "error" || msg.MsgType == "warning" {
			channelZabbix <- zabbixinteractions.MessageSettings{
				EventType: "error",
				Message:   fmt.Sprintf("%s: %s", msg.MsgType, msg.MsgData),
			}
		}

		if msg.MsgType == "info" {
			channelZabbix <- zabbixinteractions.MessageSettings{
				EventType: "info",
				Message:   msg.MsgData,
			}
		}
	}
}

// counterHandler обработчик счетчиков
func counterHandler(
	channelZabbix chan<- zabbixinteractions.MessageSettings,
	storageApp *memorytemporarystorage.CommonStorageTemporary,
	sl *simplelogger.SimpleLoggerSettings,
	counting <-chan datamodels.DataCounterSettings) {

	for data := range counting {
		d, h, m, s := supportingfunctions.GetDifference(storageApp.GetStartTimeDataCounter(), time.Now())
		patternTime := fmt.Sprintf("со старта приложения: дней %d, часов %d, минут %d, секунд %d", d, h, m, s)
		var msg string

		switch data.DataType {
		case "update accepted events":
			storageApp.IncrementAcceptedEvents()
			msg = fmt.Sprintf("принято: %d, %s", storageApp.GetAcceptedEvents(), patternTime)

		case "update processed events":
			storageApp.IncrementProcessedEvents()
			msg = fmt.Sprintf("обработано: %d, %s", storageApp.GetProcessedEvents(), patternTime)

		case "update events meet rules":
			if data.DataMsg == "subject_case" {
				storageApp.IncrementCaseEventsMeetRules()
				msg = fmt.Sprintf("подписка-'subject_case', соответствие правилам: %d, %s", storageApp.GetCaseEventsMeetRules(), patternTime)
			}

			if data.DataMsg == "subject_alert" {
				storageApp.IncrementAlertEventsMeetRules()
				msg = fmt.Sprintf("подписка-'subject_alert', соответствие правилам: %d, %s", storageApp.GetAlertEventsMeetRules(), patternTime)
			}

		case "update count insert MongoDB":
			if data.DataMsg == "subject_case" {
				storageApp.IncrementCaseInsertMongoDB()
				msg = fmt.Sprintf("подписка-'subject_case', добавлено в MongoDB: %d, %s", storageApp.GetCaseInsertMongoDB(), patternTime)
			}

			if data.DataMsg == "subject_alert" {
				storageApp.IncrementAlertInsertMongoDB()
				msg = fmt.Sprintf("подписка-'subject_alert', добавлено в MongoDB: %d, %s", storageApp.GetAlertInsertMongoDB(), patternTime)
			}

		case "update count insert Elasticserach":
			if data.DataMsg == "subject_case" {
				storageApp.IncrementCaseInsertElasticsearch()
				msg = fmt.Sprintf("подписка-'subject_case', добавлено в Elasticsearch: %d, %s", storageApp.GetCaseInsertElasticsearch(), patternTime)
			}

			if data.DataMsg == "subject_alert" {
				storageApp.IncrementAlertInsertElasticsearch()
				msg = fmt.Sprintf("подписка-'subject_alert', добавлено в Elasticsearch: %d, %s", storageApp.GetAlertInsertElasticsearch(), patternTime)
			}
		}

		_ = sl.WriteLoggingData(msg, "debug")

		channelZabbix <- zabbixinteractions.MessageSettings{
			EventType: "info",
			Message:   msg,
		}
	}
}

// interactionZabbix осуществляет взаимодействие с Zabbix
func interactionZabbix(
	ctx context.Context,
	confApp confighandler.ConfigApp,
	sl *simplelogger.SimpleLoggerSettings,
	channelZabbix <-chan zabbixinteractions.MessageSettings) error {

	connTimeout := time.Duration(7 * time.Second)
	hz, err := zabbixinteractions.NewZabbixConnection(
		ctx,
		zabbixinteractions.SettingsZabbixConnection{
			Port:              confApp.Zabbix.NetworkPort,
			Host:              confApp.Zabbix.NetworkHost,
			NetProto:          "tcp",
			ZabbixHost:        confApp.Zabbix.ZabbixHost,
			ConnectionTimeout: &connTimeout,
		})
	if err != nil {
		return err
	}

	et := make([]zabbixinteractions.EventType, len(confApp.Zabbix.EventTypes))
	for _, v := range confApp.Zabbix.EventTypes {
		et = append(et, zabbixinteractions.EventType{
			IsTransmit: v.IsTransmit,
			EventType:  v.EventType,
			ZabbixKey:  v.ZabbixKey,
			Handshake:  zabbixinteractions.Handshake(v.Handshake),
		})
	}

	if err = hz.Handler(et, channelZabbix); err != nil {
		return err
	}

	go func() {
		for err := range hz.GetChanErr() {
			_, f, l, _ := runtime.Caller(0)
			_ = sl.WriteLoggingData(fmt.Sprintf("zabbix module: '%s' %s:%d", err.Error(), f, l-1), "error")
		}
	}()

	return nil
}

func readListRules(rootDir, dirName, fileName string) (*rules.ListRule, string, error) {
	var (
		err      error
		warning  string
		warnings []string
		lr       *rules.ListRule
	)

	lr, warnings, err = rules.NewListRule(rootDir, dirName, fileName)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		return lr, warning, fmt.Errorf("%v %s:%d", err, f, l-2)
	}

	// проверяем наличие правил Pass или Passany
	if len(lr.GetRulePass()) == 0 && !lr.GetRulePassany() {
		msg := "there are no rules for handling received from NATS or all rules have failed validation"
		_, f, l, _ := runtime.Caller(0)
		return lr, warning, fmt.Errorf(" '%s' %s:%d", msg, f, l-3)
	}

	// если есть какие либо логические ошибки в файле с YAML правилами для обработки сообщений поступающих от NATS
	if len(warnings) > 0 {
		for _, v := range warnings {
			warning += fmt.Sprintln(v)
		}
	}

	return lr, warning, err
}

func main() {
	ctx, ctxCancel := signal.NotifyContext(context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	//инициализируем модуль чтения конфигурационного файла
	confApp, err := confighandler.NewConfig(ROOT_DIR)
	if err != nil {
		log.Fatalf("error module 'confighandler': %v", err)
	}

	//инициализируем модуль логирования
	sl, err := simplelogger.NewSimpleLogger(ctx, ROOT_DIR, getLoggerSettings(confApp.GetListLogs()))
	if err != nil {
		log.Fatalf("error module 'simplelogger': %v", err)
	}

	// инициализируем модуль чтения правил обработки Cases поступающих через NATS
	lrcase, warning, err := readListRules(ROOT_DIR, confApp.AppConfigRulesProcMsg.Directory, confApp.AppConfigRulesProcMsg.FileCase)
	if err != nil {
		_ = sl.WriteLoggingData(fmt.Sprint(err), "error")

		log.Fatalf("error module 'rulesinteraction': %v\n", err)
	}

	if warning != "" {
		_ = sl.WriteLoggingData(warning, "warning")
	}

	// инициализируем модуль чтения правил обработки Alerts поступающих через NATS
	lralert, warning, err := readListRules(ROOT_DIR, confApp.AppConfigRulesProcMsg.Directory, confApp.AppConfigRulesProcMsg.FileAlert)
	if err != nil {
		_ = sl.WriteLoggingData(fmt.Sprint(err), "error")

		log.Fatalf("error module 'rulesinteraction': %v\n", err)
	}

	if warning != "" {
		_ = sl.WriteLoggingData(warning, "warning")
	}

	//взаимодействие с Zabbix
	channelZabbix := make(chan zabbixinteractions.MessageSettings)
	if err := interactionZabbix(ctx, confApp, sl, channelZabbix); err != nil {
		_, f, l, _ := runtime.Caller(0)
		_ = sl.WriteLoggingData(fmt.Sprintf(" '%s' %s:%d", err.Error(), f, l-3), "error")

		log.Fatalf("error module 'zabbixinteraction': %v\n", err)
	}

	var appName string
	if an, err := supportingfunctions.GetAppName("README.md", 1); err != nil {
		_, f, l, _ := runtime.Caller(0)
		_ = sl.WriteLoggingData(fmt.Sprintf(" '%s' %s:%d", err, f, l-2), "warning")
	} else {
		appName = an
	}

	appStatus := "production"
	envValue, ok := os.LookupEnv("GO_PHELASTIC_MAIN")
	if ok && envValue == "development" {
		appStatus = envValue
	}

	natsConf := confApp.GetAppNATS()
	if natsConf.SubjectAlert == "" && natsConf.SubjectCase == "" {
		log.Fatalln("The application has been stopped. At least one subscription must be set for NATS")
	}

	log.Printf("%s application, version %s is running. Application status is '%s'\n", versionandname.GetName(), versionandname.GetVersion(), appStatus)

	//инициализируем модуль временного хранения информации
	storageApp := memorytemporarystorage.NewTemporaryStorage()

	//добавляем время инициализации счетчика хранения
	storageApp.SetStartTimeDataCounter(time.Now())

	// логирование данных
	logging := make(chan datamodels.MessageLogging)
	go loggingHandler(channelZabbix, sl, logging)

	//вывод данных счетчика
	counting := make(chan datamodels.DataCounterSettings)
	go counterHandler(channelZabbix, storageApp, sl, counting)

	//инициализация модуля для взаимодействия с NATS (Данный модуль обязателен для взаимодействия)
	natsModule, err := natsinteractions.NewClientNATS(*confApp.GetAppNATS(), storageApp, logging, counting)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		_ = sl.WriteLoggingData(fmt.Sprintf(" '%s' %s:%d", err, f, l-2), "error")

		log.Fatalf("error module 'natsclient': %v\n", err)
	}

	// инициализация модуля для взаимодействия с СУБД MongoDB
	mongodbModule, err := mongodbinteractions.HandlerMongoDB(*confApp.GetAppMongoDB(), logging, counting)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		_ = sl.WriteLoggingData(fmt.Sprintf(" '%s' %s:%d", err, f, l-2), "error")

		log.Fatalf("error module 'mongodbclient': %v\n", err)
	}

	//инициализация модуля применяемого для обогащения кейсов
	eventenrichmentModule, err := eventenrichmentmodule.NewEventEnrichmentModule(ctx, confApp.NCIRCC, confApp.ZabbixJsonRPC, logging)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		_ = sl.WriteLoggingData(fmt.Sprintf(" '%s' %s:%d", err, f, l-2), "error")

		log.Fatalf("error module 'eventenrichmentmodule': %v\n", err)
	}

	//инициализация модуля для взаимодействия с ElasticSearch
	esModule, err := elasticsearchinteractions.HandlerElasticSearch(*confApp.GetAppES(), logging, counting)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		_ = sl.WriteLoggingData(fmt.Sprintf(" '%s' %s:%d", err, f, l-2), "error")
	}

	logging <- datamodels.MessageLogging{
		MsgData: "application '" + appName + "' is started",
		MsgType: "info",
	}

	go func() {
		sigChan := make(chan os.Signal, 1)
		osCall := <-sigChan
		log.Printf("system call:%+v", osCall)

		close(counting)
		close(logging)
		close(channelZabbix)

		ctxCancel()
	}()

	core := coremodule.NewCoreHandler(storageApp, logging, counting)
	core.CoreHandler(ctx, lrcase, lralert, natsModule, esModule, mongodbModule, eventenrichmentModule)
}
