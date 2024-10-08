package testsyncmutex_test

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sync"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type MyTestMutix struct {
	List map[string]string
	mu   sync.RWMutex
}

var _ = Describe("Testsyncmutex", Ordered, func() {
	var (
		f *os.File

		myTestMutix MyTestMutix

		err error
	)

	BeforeAll(func() {
		myTestMutix.List = map[string]string{
			"one":   "any one",
			"two":   "any two",
			"three": "any three",
			"four":  "any four",
			"five":  "any five",
			"six":   "any six",
			"seven": "any seven",
			"eigth": "any eigth",
			"nine":  "any nine",
			"ten":   "any ten",
		}

		go func() {
			f, err = os.Create("proff.out")
			if err != nil {
				log.Fatal("could not create CPU profile: ", err)
			}

			if err = pprof.StartCPUProfile(f); err != nil {
				log.Fatal("could not start CPU profile: ", err)
			}
		}()
	})

	Context("Test 1.", func() {
		It("Is ok", func() {
			count := 10
			var wg sync.WaitGroup
			wg.Add(count)

			var i int
			c := time.Tick(1 * time.Second)

			for range c {
				//for i := 0; i < count; i++ {
				go func(num int) {
					defer wg.Done()

					myTestMutix.mu.RLock()
					for k, v := range myTestMutix.List {
						fmt.Println("гроутина №", num, " - value:", v)

						if k == "ten" {
							delete(myTestMutix.List, k)
						}
					}
					myTestMutix.mu.RUnlock()
				}(i)

				i++

				if i == count {
					break
				}
			}

			wg.Wait()

			Expect(true).Should(BeTrue())
		})
	})

	AfterAll(func() {
		defer f.Close() // error handling omitted for example
		defer pprof.StopCPUProfile()
	})
})
