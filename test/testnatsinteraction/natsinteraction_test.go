package testnatsinteraction_test

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/supportingfunctions"
	//"placeholder_elasticsearch/test/testnatsinteraction"
)

// ModuleNATS инициализированный модуль
// ChanOutputMISP - канал для отправки полученных данных из модуля
type ModuleNATS struct {
	chanOutputNATS chan SettingsOutputChan
}

type SettingsOutputChan struct {
	MsgId   string
	MsgType string
	Data    []byte
}

type SettingsInputChan struct {
	Command, EventId, TaskId string
}

func (mnats ModuleNATS) GetDataReceptionChannel() <-chan SettingsOutputChan /*[]byte*/ {
	return mnats.chanOutputNATS
}

func (mnats ModuleNATS) SendingDataOutput(data SettingsOutputChan) {
	mnats.chanOutputNATS <- data
}

var _ = Describe("Natsinteraction", Ordered, func() {
	var (
		errConnect error

		logging = make(chan datamodels.MessageLogging)
		done    = make(chan struct{})
		mnats   ModuleNATS
	)

	const (
		Host = "nats.cloud.gcm"
		Port = 4222
	)

	newClient := func(
		logging chan<- datamodels.MessageLogging,
		done chan<- struct{},
	) (*ModuleNATS, error) {
		nc, err := nats.Connect(
			fmt.Sprintf("%s:%d", Host, Port),
			nats.MaxReconnects(-1),
			nats.ReconnectWait(3*time.Second))
		_, f, l, _ := runtime.Caller(0)
		if err != nil {
			return &mnats, fmt.Errorf("'%v' %s:%d", err, f, l-4)
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

		nc.Subscribe("main_caseupdate", func(m *nats.Msg) {
			mnats.chanOutputNATS <- SettingsOutputChan{
				MsgId:   uuid.NewString(),
				MsgType: "caseupdate",
				Data:    m.Data,
			}
		})

		nc.Subscribe("main_alertupdate", func(msg *nats.Msg) {
			mnats.chanOutputNATS <- SettingsOutputChan{
				MsgId:   uuid.NewString(),
				MsgType: "alertupdate",
				Data:    msg.Data,
			}

			logging <- datamodels.MessageLogging{
				MsgType: "STOP TEST",
			}
		})

		log.Printf("Connect to NATS with address %s:%d\n", Host, Port)

		return &mnats, nil
	}

	BeforeAll(func() {
		mnats.chanOutputNATS = make(chan SettingsOutputChan)

		go func() {
			fmt.Println("___ Logging START")
			defer fmt.Println("___ Logging STOP")

			for log := range logging {
				if log.MsgType == "error" {
					fmt.Println("----", log, "----")
				}

				if log.MsgType == "STOP TEST" {
					return
				}
			}
		}()

		_, errConnect = newClient(logging, done)
	})

	Context("Тест 1. Подключение к NATS", func() {
		It("Должно быть успешно установлено подключение к NATS", func() {
			Expect(errConnect).ShouldNot(HaveOccurred())
		})

		It("Должно быть получено сообщение типа 'alertupdate'", func() {
			var str string
			var err error

			for data := range mnats.chanOutputNATS {
				if data.MsgType == "alertupdate" {
					done <- struct{}{}

					str, err = supportingfunctions.NewReadReflectJSONSprint(data.Data)

					break
				}
			}

			Expect(err).ShouldNot(HaveOccurred())
			fmt.Println(str)
			Expect(len(str)).ShouldNot(Equal(0))
		})
	})
})
