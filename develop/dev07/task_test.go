package main

import (
	"fmt"
	"testing"
	"time"
)

func TestOrChan1(t *testing.T) {
	ch1 := make(chan string, 3)
	ch2 := make(chan int, 3)

	ch1 <- "one"
	ch1 <- "two"
	ch1 <- "three"

	ch2 <- 1
	ch2 <- 2
	ch2 <- 3

	sig := func(data ...interface{}) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			for _, v := range data {
				c <- v
			}
			time.Sleep(1 * time.Second)
		}()
		return c
	}

	for v := range or(sig(ch1), sig(ch2)) {
		fmt.Println(v)
	}
}
