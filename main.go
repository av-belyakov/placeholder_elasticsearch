package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/av-belyakov/simplelogger"

	"placeholder_elasticsearch/confighandler"
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/elasticsearchinteractions"

	"placeholder_elasticsearch/coremodule"
	"placeholder_elasticsearch/memorytemporarystorage"
	"placeholder_elasticsearch/mongodbinteractions"
	"placeholder_elasticsearch/natsinteractions"
	rules "placeholder_elasticsearch/rulesinteraction"
	"placeholder_elasticsearch/supportingfunctions"
	"placeholder_elasticsearch/zabbixinteractions"
)

const ROOT_DIR = "placeholder_elasticsearch"

var (
	err             error
	sl              simplelogger.SimpleLoggerSettings
	confApp         confighandler.ConfigApp
	hz              *zabbixinteractions.HandlerZabbixConnection
	warning         string
	storageApp      *memorytemporarystorage.CommonStorageTemporary
	lrcase, lralert *rules.ListRule

	iz       chan string
	logging  chan datamodels.MessageLogging
	counting chan datamodels.DataCounterSettings
)

func getAppName(pf string, nl int) (string, error) {
	var line string

	f, err := os.OpenFile(pf, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return line, err
	}
	defer f.Close()

	num := 1
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		if num == nl {
			return sc.Text(), nil
		}

		num++
	}

	return line, nil
}

func getLoggerSettings(cls []confighandler.LogSet) []simplelogger.MessageTypeSettings {
	loggerConf := make([]simplelogger.MessageTypeSettings, 0, len(cls))

	for _, v := range cls {
		loggerConf = append(loggerConf, simplelogger.MessageTypeSettings{
			MsgTypeName:   v.MsgTypeName,
			WritingFile:   v.WritingFile,
			PathDirectory: v.PathDirectory,
			WritingStdout: v.WritingStdout,
			MaxFileSize:   v.MaxFileSize,
		})
	}

	return loggerConf
}

// loggingHandler обработчик логов
func loggingHandler(
	iz chan<- string,
	sl simplelogger.SimpleLoggerSettings,
	logging <-chan datamodels.MessageLogging) {

	for msg := range logging {
		_ = sl.WriteLoggingData(msg.MsgData, msg.MsgType)

		if msg.MsgType == "error" || msg.MsgType == "info" {
			iz <- msg.MsgData
		}
	}
}

// counterHandler обработчик счетчиков
func counterHandler(
	iz chan<- string,
	storageApp *memorytemporarystorage.CommonStorageTemporary,
	counting <-chan datamodels.DataCounterSettings) {

	for data := range counting {
		switch data.DataType {
		case "update accepted events":
			storageApp.SetAcceptedEventsDataCounter(data.Count)
		case "update processed events":
			storageApp.SetProcessedEventsDataCounter(data.Count)
		case "update events meet rules":
			storageApp.SetEventsMeetRulesDataCounter(data.Count)
		case "events do not meet rules":
			storageApp.SetEventsDoNotMeetRulesDataCounter(data.Count)
		case "update count insert MongoDB":
			storageApp.SetInsertMongoDBDataCounter(data.Count)
		case "update count insert Elasticserach":
			storageApp.SetInsertElasticsearchDataCounter(data.DataMsg, data.Count)
		}

		d, h, m, s := supportingfunctions.GetDifference(storageApp.GetStartTimeDataCounter(), time.Now())

		patternReciveEvents := fmt.Sprintf("принято: %d", storageApp.GetAcceptedEventsDataCounter())
		patternRuleIsOk := fmt.Sprintf("соответствие правилам: %d", storageApp.GetEventsMeetRulesDataCounter())
		patternInsertMongoDB := fmt.Sprintf("добавлено в MongoDB: %d", storageApp.GetInsertMongoDBDataCounter())

		num, _ := storageApp.GetInsertElasticsearchDataCounter(data.DataMsg)
		patternInsertES := fmt.Sprintf("добавлено в Elasticsearch: %d", num)
		patternTime := fmt.Sprintf("со старта приложения: дней %d, часов %d, минут %d, секунд %d", d, h, m, s)
		msg := fmt.Sprintf("подписка-'%s', %s, %s, %s, %s %s", data.DataMsg, patternReciveEvents, patternRuleIsOk, patternInsertMongoDB, patternInsertES, patternTime)

		log.Printf("\t%s\n", msg)

		iz <- msg
	}
}

// interactionZabbix осуществляет взаимодействие с Zabbix
func interactionZabbix(
	ctx context.Context,
	confApp confighandler.ConfigApp,
	hz *zabbixinteractions.HandlerZabbixConnection,
	sl simplelogger.SimpleLoggerSettings,
	iz <-chan string) {
	co := confApp.GetCommonApp()
	t := time.Tick(time.Duration(co.Zabbix.TimeInterval) * time.Minute)

	for {
		select {
		case <-ctx.Done():
			return

		case <-t:
			if !co.Zabbix.IsTransmit {
				continue
			}

			if _, err := hz.SendData([]string{co.Zabbix.Handshake}); err != nil {
				_, f, l, _ := runtime.Caller(0)
				_ = sl.WriteLoggingData(fmt.Sprintf(" '%v' %s:%d", err, f, l-1), "error")
			}

		case msg := <-iz:
			if !co.Zabbix.IsTransmit {
				continue
			}

			if _, err := hz.SendData([]string{msg}); err != nil {
				_, f, l, _ := runtime.Caller(0)
				_ = sl.WriteLoggingData(fmt.Sprintf(" '%v' %s:%d", err, f, l-1), "error")
			}
		}
	}
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

func init() {
	iz = make(chan string)
	logging = make(chan datamodels.MessageLogging)
	counting = make(chan datamodels.DataCounterSettings)

	//инициализируем модуль чтения конфигурационного файла
	confApp, err = confighandler.NewConfig(ROOT_DIR)
	if err != nil {
		log.Fatalf("error module 'confighandler': %v", err)
	}

	//инициализируем модуль логирования
	sl, err = simplelogger.NewSimpleLogger(ROOT_DIR, getLoggerSettings(confApp.GetListLogs()))
	if err != nil {
		log.Fatalf("error module 'simplelogger': %v", err)
	}

	// инициализируем модуль чтения правил обработки Cases поступающих через NATS
	lrcase, warning, err = readListRules(ROOT_DIR, confApp.AppConfigRulesProcMsg.Directory, confApp.AppConfigRulesProcMsg.FileCase)
	if err != nil {
		_ = sl.WriteLoggingData(fmt.Sprint(err), "error")

		log.Fatalf("error module 'rulesinteraction': %v\n", err)
	}

	// инициализируем модуль чтения правил обработки Alerts поступающих через NATS
	lrcase, warning, err = readListRules(ROOT_DIR, confApp.AppConfigRulesProcMsg.Directory, confApp.AppConfigRulesProcMsg.FileAlert)
	if err != nil {
		_ = sl.WriteLoggingData(fmt.Sprint(err), "error")

		log.Fatalf("error module 'rulesinteraction': %v\n", err)
	}

	if warning != "" {
		_ = sl.WriteLoggingData(warning, "warning")
	}

	commOpt := confApp.GetCommonApp()
	host := fmt.Sprintf("%s:%d", commOpt.Zabbix.NetworkHost, commOpt.Zabbix.NetworkPort)

	//инициализируем модуль связи с Zabbix
	hz = zabbixinteractions.NewHandlerZabbixConnection(host, commOpt.Zabbix.ZabbixHost, commOpt.Zabbix.ZabbixKey)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			_ = sl.WriteLoggingData(fmt.Sprintf("stop 'main' function, %v", err), "error")
		}
	}()

	var appName string
	appStatus := "production"
	if an, err := getAppName("README.md", 1); err != nil {
		_, f, l, _ := runtime.Caller(0)
		_ = sl.WriteLoggingData(fmt.Sprintf(" '%s' %s:%d", err, f, l-2), "warning")
	} else {
		appName = an
	}

	envValue, ok := os.LookupEnv("GO_PHELASTIC_MAIN")
	if ok && envValue == "development" {
		appStatus = envValue
	}

	natsConf := confApp.GetAppNATS()
	if natsConf.SubjectAlert == "" && natsConf.SubjectCase == "" {
		log.Fatalln("The application has been stopped. At least one subscription must be set for NATS")
	}

	appVersion := supportingfunctions.GetAppVersion(appName)
	log.Printf("Placeholder_Elasticsearch application, version %s is running. Application status is '%s'\n", appVersion, appStatus)

	//инициализируем модуль временного хранения информации
	storageApp = memorytemporarystorage.NewTemporaryStorage()

	//добавляем время инициализации счетчика хранения
	storageApp.SetStartTimeDataCounter(time.Now())

	//взаимодействие с Zabbix
	ctxZabbix, closeZabbix := context.WithCancel(context.Background())
	go interactionZabbix(ctxZabbix, confApp, hz, sl, iz)

	//вывод данных счетчика
	go counterHandler(iz, storageApp, counting)

	// логирование данных
	go loggingHandler(iz, sl, logging)

	//инициализация модуля для взаимодействия с NATS (Данный модуль обязателен для взаимодействия)
	natsModule, err := natsinteractions.NewClientNATS(*confApp.GetAppNATS(), storageApp, logging, counting)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		_ = sl.WriteLoggingData(fmt.Sprintf(" '%s' %s:%d", err, f, l-2), "error")

		log.Fatal(err)
	}

	// инициализация модуля для взаимодействия с СУБД MongoDB
	mongodbModule, err := mongodbinteractions.HandlerMongoDB(*confApp.GetAppMongoDB(), logging, counting)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		_ = sl.WriteLoggingData(fmt.Sprintf(" '%s' %s:%d", err, f, l-2), "error")

		log.Fatal(err)
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

	ctxCoreHandler, closeCoreHandler := context.WithCancel(context.Background())
	defer func() {
		close(counting)
		close(logging)

		closeZabbix()
		closeCoreHandler()
	}()
	core := coremodule.NewCoreHandler(storageApp, logging, counting)
	core.CoreHandler(ctxCoreHandler, lrcase, lralert, natsModule, esModule, mongodbModule)
}
