package main

import (
	"fmt"
	"testing"
	"time"
)

func TestOrChan1(t *testing.T) {
	sig := func(data ...interface{}) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			for _, v := range data {
				c <- v
			}
			time.Sleep(time.Second * 1)
		}()
		return c
	}

	res := or(sig("one", 1), sig("two", 2), sig("three", 3))

	for v := range res {
		fmt.Println(v)
	}

	if <-res != nil {
		t.Errorf("The channel is not closed")
	}
}
