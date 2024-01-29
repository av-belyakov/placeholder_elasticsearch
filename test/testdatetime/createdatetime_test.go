package testdatetime_test

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Createdatetime", func() {
	Context("Test 1. Create current date time", func() {
		It("Должно быть спешно полученно текущее время в соответсвующем формате", func() {
			currentDatetime := time.Now()

			fmt.Printf("time.Now(): %s\n", currentDatetime.String())
			fmt.Printf("Date Time UTC: %s\n", currentDatetime.UTC())
			fmt.Printf("Date Time RFC3339: %s\n", currentDatetime.Format(time.RFC3339))
			fmt.Printf("Date Time RFC3339: %s\n", currentDatetime.Format(time.RFC3339))
			fmt.Printf("Unixtime millise: %d\n", time.Now().UnixMilli())
			fmt.Printf("UnixMilli(1647277028643).Format(time.RFC3339): %s\n", time.UnixMilli(1647277028643).Format(time.RFC3339))
			fmt.Println("now 1:", time.Now().String(), " now 2:", time.UnixMilli(time.Now().UnixMilli()).Format(time.RFC3339))
			fmt.Println("empty:", time.UnixMilli(0).Format(time.RFC3339))

			Expect(true).Should(BeTrue())
		})
	})
})
