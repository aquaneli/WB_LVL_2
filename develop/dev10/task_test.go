package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"testing"
	"time"
)

func TestWrite(t *testing.T) {
	res := []string{}
	go func() {
		res = newServer()
	}()

	conn, err := net.DialTimeout("tcp", "localhost:8080", time.Second*10)
	if err != nil {
		log.Fatal(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	sig := make(chan error)

	r, w := io.Pipe()
	os.Stdin = r

	WriteData(conn, sig, &wg)

	wg.Wait()

}

func newServer() []string {
	addr, _ := net.ResolveTCPAddr("tcp", ":8080")
	ln, _ := net.ListenTCP("tcp", addr)
	res := []string{}
	defer ln.Close()

	for {

		conn, _ := ln.Accept()
		r := bufio.NewReader(conn)

		for {
			b, _ := r.ReadBytes('\n')
			if string(b) == "stop" {
				return res
			}
			// fmt.Println(string(b))
			res = append(res, string(b))
		}

	}

}
