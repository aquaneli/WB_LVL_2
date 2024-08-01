package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		// sig(2*time.Hour),
		// sig(5*time.Minute),
		sig(1*time.Second),
		sig(2*time.Second),
		sig(5*time.Second),
		// sig(1*time.Hour),
		// sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	if len(channels) == 0 {
		return nil
	}
	//сюда будут передаваться данные из закрытого канала
	doneChan := make(chan interface{})
	wg := sync.WaitGroup{}
	wg.Add(len(channels))

	for _, ch := range channels {
		go func(doneChan chan interface{}, ch <-chan interface{}) {
			defer wg.Done()
			for val := range ch {
				doneChan <- val
			}
		}(doneChan, ch)
	}

	go func(doneChan chan interface{}) {
		wg.Wait()
		close(doneChan)
		fmt.Println("doneChan closed")
	}(doneChan)

	return doneChan
}
