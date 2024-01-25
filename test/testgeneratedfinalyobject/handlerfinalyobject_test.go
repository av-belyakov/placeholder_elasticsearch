package testgeneratedfinalyobject_test

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/coremodule"
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/supportingfunctions"
)

func readFileJson(fpath, fname string) ([]byte, error) {
	var newResult []byte

	rootPath, err := supportingfunctions.GetRootPath("placeholder_misp")
	if err != nil {
		return newResult, err
	}

	fmt.Println("func 'readFileJson', path = ", path.Join(rootPath, fpath, fname))

	f, err := os.OpenFile(path.Join(rootPath, fpath, fname), os.O_RDONLY, os.ModePerm)
	if err != nil {
		return newResult, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		newResult = append(newResult, sc.Bytes()...)
	}

	return newResult, nil
}

var _ = Describe("Handlerfinalyobject", Ordered, func() {
	var (
		errReadFile error
		exampleByte []byte
		decodeJson  *coremodule.DecodeJsonMessageSettings

		logging  chan datamodels.MessageLogging
		counting chan datamodels.DataCounterSettings
	)

	BeforeAll(func() {
		logging = make(chan datamodels.MessageLogging)
		counting = make(chan datamodels.DataCounterSettings)

		go func() {
			fmt.Println("___ Logging START")
			defer fmt.Println("___ Logging STOP")

			for log := range logging {
				fmt.Println("----", log, "----")
			}
		}()

		//вывод данных счетчика
		go func() {
			dc := storageApp.GetDataCounter()
			d, h, m, s := supportingfunctions.GetDifference(dc.StartTime, time.Now())

			fmt.Printf("\tСОБЫТИЙ принятых/обработанных: %d/%d, соответствие/не соответствие правилам: %d/%d, время со старта приложения: дней %d, часов %d, минут %d, секунд %d\n", dc.AcceptedEvents, dc.ProcessedEvents, dc.EventsMeetRules, dc.EventsDoNotMeetRules, d, h, m, s)

			for d := range counting {
				switch d.DataType {
				case "update accepted events":
					storageApp.SetAcceptedEventsDataCounter(d.Count)
				case "update processed events":
					storageApp.SetProcessedEventsDataCounter(d.Count)
				case "update events meet rules":
					storageApp.SetEventsMeetRulesDataCounter(d.Count)
				case "events do not meet rules":
					storageApp.SetEventsDoNotMeetRulesDataCounter(d.Count)
				}

				dc := storageApp.GetDataCounter()
				d, h, m, s := supportingfunctions.GetDifference(dc.StartTime, time.Now())

				fmt.Printf("\tСОБЫТИЙ принятых/обработанных: %d/%d, соответствие/не соответствие правилам: %d/%d, время со старта приложения: дней %d, часов %d, минут %d, секунд %d\n", dc.AcceptedEvents, dc.ProcessedEvents, dc.EventsMeetRules, dc.EventsDoNotMeetRules, d, h, m, s)
			}
		}()

		exampleByte, errReadFile = readFileJson("testing/test_json", "example_caseId_33705.json")

		decodeJson = coremodule.NewDecodeJsonMessageSettings(listRule, logging, counting)
	})

	Context("Тест 1. Проверка чтения тестового JSON файла", func() {
		It("При чтении файла не должно быть ошибок", func() {
			Expect(errReadFile).ShouldNot(HaveOccurred())
		})
	})

	/*
		Context("Тест 2. Проверка преобразования тестового JSON файла в бинарный вид", func ()  {
			It("При преобразовании файла не должно быть ошибок", func ()  {

			})
		})

		Context("", func ()  {
			It("", func ()  {

			})
		})
	*/
})
