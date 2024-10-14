package natsinteractions

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"

	"placeholder_elasticsearch/confighandler"
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/memorytemporarystorage"
)

func NewClientNATS(
	conf confighandler.AppConfigNATS,
	storageApp *memorytemporarystorage.CommonStorageTemporary,
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) (*ModuleNATS, error) {
	mnats := ModuleNATS{
		chanOutputNATS: make(chan SettingsOutputChan),
		chanInputNATS:  make(chan SettingsInputChan),
	}

	//инициируем хранилище для дескрипторов сообщений NATS
	ns = NewStorageNATS()

	if conf.SubjectCase == "" && conf.SubjectAlert == "" {
		_, f, l, _ := runtime.Caller(0)
		return &mnats, fmt.Errorf("'there is not a single subscription available for NATS in the configuration file' %s:%d", f, l-1)
	}

	subjects := map[string]string{
		"subject_case":  conf.SubjectCase,
		"subject_alert": conf.SubjectAlert,
	}

	nc, err := nats.Connect(
		fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		nats.MaxReconnects(-1),
		nats.ReconnectWait(3*time.Second))
	_, f, l, _ := runtime.Caller(0)
	if err != nil {
		return &mnats, fmt.Errorf("'%s' %s:%d", err.Error(), f, l-4)
	}

	//обработка разрыва соединения с NATS
	nc.SetDisconnectErrHandler(func(c *nats.Conn, err error) {
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("the connection with NATS has been disconnected %s:%d", f, l-4),
			MsgType: "error",
		}
	})

	//обработка переподключения к NATS
	nc.SetReconnectHandler(func(c *nats.Conn) {
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("the connection to NATS has been re-established %s:%d", f, l-4),
			MsgType: "info",
		}
	})

	for k, v := range subjects {
		//не добавляем обработчик если подписка пуста
		if v == "" {
			continue
		}

		nc.Subscribe(v, func(m *nats.Msg) {
			mnats.chanOutputNATS <- SettingsOutputChan{
				MsgId:       ns.setElement(m),
				SubjectType: k,
				Data:        m.Data,
			}

			//счетчик принятых событий
			counting <- datamodels.DataCounterSettings{
				DataType: "update accepted events",
				DataMsg:  k,
				Count:    1,
			}
		})
	}

	go SenderData(ns, mnats.GetDataDeliveryChannel(), logging)

	log.Printf("Connect to NATS with address %s:%d\n", conf.Host, conf.Port)

	return &mnats, nil
}

// SenderData выполняет отправку в NATS данных приходящих в модуль из основной части приложения
func SenderData(ns *natsStorage, chanInput <-chan SettingsInputChan, logging chan<- datamodels.MessageLogging) {

	fmt.Println("func 'SenderData' START")

	for data := range chanInput {
		fmt.Println("func 'SenderData' DATA:", data)

		if data.Command != "send tag" {
			continue
		}
		fmt.Println("func 'SenderData' CHECK 111")

		//получаем дескриптор соединения с NATS для отправки eventId
		ncd, ok := ns.getElement(data.TaskId)
		if !ok {
			_, f, l, _ := runtime.Caller(0)

			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("connection descriptor for task id '%s' not found %s:%d", data.TaskId, f, l-2),
				MsgType: "error",
			}

			continue
		}
		fmt.Println("func 'SenderData' CHECK 222")

		resMsg := datamodels.NewResponseMessage()
		res, err := json.Marshal(resMsg.GetResponseMessageFromMispToTheHave())
		if err != nil {
			_, f, l, _ := runtime.Caller(0)

			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("%s %s:%d", err.Error(), f, l-2),
				MsgType: "error",
			}

			continue
		}
		fmt.Println("func 'SenderData' CHECK 333")
		request, err := json.MarshalIndent(resMsg.GetResponseMessageFromMispToTheHave(), "", " ")
		if err != nil {
			_, f, l, _ := runtime.Caller(0)

			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("%s %s:%d", err.Error(), f, l-2),
				MsgType: "error",
			}
		}
		fmt.Println(string(request))

		//отправляем в NATS пакет с eventId для добавления его в TheHive
		if err := ncd.Respond(res); err != nil {
			_, f, l, _ := runtime.Caller(0)

			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("%s %s:%d", err.Error(), f, l-2),
				MsgType: "error",
			}
		}

		fmt.Println("func 'SenderData' SEND DATA IS SUCCEFULY")

	}
}
