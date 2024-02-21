package supportingfunctions

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateChannelDuplication(t *testing.T) {
	inputChan := make(chan struct{})
	chanCount := 2

	channels := CreateChannelDuplication[struct{}](inputChan, chanCount)
	assert.Equal(t, len(channels), 2)

	fmt.Println(channels)

	var num int64
	var wg sync.WaitGroup
	for i := 0; i < chanCount; i++ {
		wg.Add(1)

		fmt.Println(i)

		go func(i int) {
			defer wg.Done()

			<-channels[i]
			atomic.AddInt64(&num, 1)
		}(i)
	}

	inputChan <- struct{}{}
	close(inputChan)
	wg.Wait()

	assert.Equal(t, num, int64(chanCount))
}
