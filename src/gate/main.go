// gate project main.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	exit_chan := make(chan int)

	wg.Add(10)
	j := 0
	for i := 0; i < 10; i++ {
		go func(chan int) {
			j = j + 1
			select {
			case <-exit_chan:
				fmt.Printf("this is %d\n", j)

			}
			if j == 10 {
				time.Sleep(time.Duration(10) * time.Second)
			}
			wg.Done()
		}(exit_chan)
	}

	close(exit_chan)
	wg.Wait()
}
