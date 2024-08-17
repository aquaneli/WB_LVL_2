package main

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

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
