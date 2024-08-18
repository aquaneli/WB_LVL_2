package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestWrite(t *testing.T) {
	str := make(chan []byte)
	go func() {
		newServerWrite(str)
	}()

	conn, err := net.DialTimeout("tcp", "localhost:8080", time.Second*10)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)
	sig := make(chan error)

	r, w, _ := os.Pipe()
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = r

	go func() {
		defer w.Close()
		w.WriteString("qwe\n")
		w.WriteString("privet\n")
		w.WriteString("good job\n")
		w.WriteString("stop\n")
	}()

	go writeData(conn, sig, &wg)
	wg.Wait()

	time.Sleep(time.Second * 1)
	expect := []string{"qwe\n", "privet\n", "good job\n"}
	i := 0
	for v := range str {
		if !reflect.DeepEqual(string(v), expect[i]) {
			t.Errorf("expected %q but got %q", expect, string(v))
		}
		i++
	}
	time.Sleep(time.Second * 1)
	close(sig)
}

func newServerWrite(str chan []byte) {
	addr, _ := net.ResolveTCPAddr("tcp", ":8080")
	ln, _ := net.ListenTCP("tcp", addr)
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		r := bufio.NewReader(conn)
		for {
			b, _ := r.ReadBytes('\n')
			if string(b) == "stop\n" {
				close(str)
				return
			}
			str <- b
		}
	}
}
