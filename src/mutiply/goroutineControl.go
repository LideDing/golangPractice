package mutiply

import (
	"fmt"
	"sync"
	"time"
)

func GoroutineControl(totalGoroutine, maxGoroutine int) {
	var wg sync.WaitGroup
	ch := make(chan int, maxGoroutine)

	for i := 0; i < totalGoroutine; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() {
				<-ch
				wg.Done()
			}()
			ch <- i
			fmt.Printf("goroutine %d start\n", i)
			time.Sleep(time.Second * 1)
			fmt.Printf("goroutine %d end\n", i)
		}(i)
	}
	wg.Wait()
}
