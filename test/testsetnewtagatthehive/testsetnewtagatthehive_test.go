package testsetnewtagatthehive_test

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/supportingfunctions"
)

var (
	ns   *natsStorage
	once sync.Once
)

type natsStorage struct {
	mutex   sync.Mutex
	storage map[string]messageDescriptors
}

type messageDescriptors struct {
	timeCreate int64
	msgNats    *nats.Msg
}

func NewStorageNATS() *natsStorage {
	once.Do(func() {
		ns = &natsStorage{storage: make(map[string]messageDescriptors)}

		go checkLiveTime(ns)
	})

	return ns
}

func checkLiveTime(ns *natsStorage) {
	for range time.Tick(5 * time.Second) {
		go func() {
			ns.mutex.Lock()
			defer ns.mutex.Unlock()

			for k, v := range ns.storage {
				if time.Now().Unix() > (v.timeCreate + 360) {
					ns.deleteElement(k)
				}
			}
		}()
	}
}

func (ns *natsStorage) getElement(id string) (*nats.Msg, bool) {
	if elem, ok := ns.storage[id]; ok {
		return elem.msgNats, ok
	}

	return nil, false
}

func (ns *natsStorage) setElement(m *nats.Msg) string {
	id := uuid.New().String()

	ns.mutex.Lock()
	defer ns.mutex.Unlock()

	ns.storage[id] = messageDescriptors{
		timeCreate: time.Now().Unix(),
		msgNats:    m,
	}

	return id
}

func (ns *natsStorage) deleteElement(id string) {
	delete(ns.storage, id)
}

type receivedElement struct {
	Event eventElement `json:"event"`
}

type eventElement struct {
	ObjectType string        `json:"objectType"`
	Object     objectElement `json:"object"`
}

type objectElement struct {
	CaseId string `json:"caseId"`
}

// ModuleNATS инициализированный модуль
// ChanOutputMISP - канал для отправки полученных данных из модуля
type ModuleNATS struct {
	chanOutputNATS chan SettingsOutputChan
	chanInputNATS  chan SettingsInputChan
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

var _ = Describe("Testsetnewtagatthehive", Ordered, func() {
	var (
		errConnect error

		logging = make(chan datamodels.MessageLogging)
		done    = make(chan struct{})
		mnats   ModuleNATS
	)

	const (
		MyNumTestCase = "39100"
		Host          = "nats.cloud.gcm"
		Port          = 4222
	)

	newClient := func(
		logging chan<- datamodels.MessageLogging,
		done chan<- struct{},
	) (*ModuleNATS, error) {
		storage := NewStorageNATS()

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

		go func() {
			for msg := range mnats.chanInputNATS {
				if msg.Command != "set tag" {
					continue
				}

				nm, ok := storage.getElement(msg.TaskId)
				Expect(ok).Should(BeTrue())

				resMsg := datamodels.NewResponseMessage()
				res, err := json.Marshal(resMsg.GetResponseMessageFromMispToTheHave())
				Expect(err).ShouldNot(HaveOccurred())

				err = nm.Respond(res)
				Expect(err).ShouldNot(HaveOccurred())

				fmt.Println("sended tag to TheHive-------->")
				req, err := json.MarshalIndent(res, "", " ")
				Expect(err).ShouldNot(HaveOccurred())

				fmt.Println(string(req))
			}
		}()

		nc.Subscribe("main_caseupdate", func(msg *nats.Msg) {
			fmt.Println("MAIN CASEUPDATE")

			str, err := supportingfunctions.NewReadReflectJSONSprint(msg.Data)
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Println(str)

			mnats.chanOutputNATS <- SettingsOutputChan{
				MsgId:   storage.setElement(msg),
				MsgType: "caseupdate",
				Data:    msg.Data,
			}
		})

		//nc.Subscribe("main_alertupdate", func(msg *nats.Msg) {
		nc.Subscribe("main_alertupdate", func(msg *nats.Msg) {
			fmt.Println("MAIN ALERT")

			//str, err := supportingfunctions.NewReadReflectJSONSprint(msg.Data)
			//Expect(err).ShouldNot(HaveOccurred())
			//fmt.Println(str)

			mnats.chanOutputNATS <- SettingsOutputChan{
				MsgId:   storage.setElement(msg),
				MsgType: "alertupdate",
				Data:    msg.Data,
			}
		})

		log.Printf("Connect to NATS with address %s:%d\n", Host, Port)
		done <- struct{}{}

		return &mnats, nil
	}

	BeforeAll(func() {
		mnats.chanOutputNATS = make(chan SettingsOutputChan)
		mnats.chanInputNATS = make(chan SettingsInputChan)

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
			go func() {
				for v := range done {
					fmt.Println("SIGNAL DONE", v)
				}
			}()

			for data := range mnats.chanOutputNATS {
				fmt.Println("chanOutputNATS msgType =", data.MsgType)

				if data.MsgType == "caseupdate" {
					c := receivedElement{}
					err := json.Unmarshal(data.Data, &c)
					Expect(err).ShouldNot(HaveOccurred())

					fmt.Printf("Received CASE:%+v", c)
					if c.Event.ObjectType == "case" && c.Event.Object.CaseId == MyNumTestCase {
						mnats.chanInputNATS <- SettingsInputChan{
							Command: "set tag",
							TaskId:  data.MsgId,
						}

						done <- struct{}{}
						close(mnats.chanInputNATS)

						break
					}
				}

				if data.MsgType == "alertupdate" {
					c := receivedElement{}
					err := json.Unmarshal(data.Data, &c)
					Expect(err).ShouldNot(HaveOccurred())

					fmt.Printf("Received ALERT:%+v", c)
				}
			}

			Expect(true).ShouldNot(BeTrue())
		})
	})
})
