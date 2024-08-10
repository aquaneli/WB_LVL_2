package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	d, socket := flags()
	conn, err := net.DialTimeout("tcp", socket, *d)
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	sig := make(chan error)
	go WriteData(conn, sig, &wg)
	go ReadData(conn, sig, &wg)

	wg.Wait()
}

func flags() (*time.Duration, string) {
	d := flag.Duration("timeout", time.Second*10, "server connection timeout")
	flag.Parse()

	if len(flag.Args()) != 2 {
		log.Fatal("usage: ./task [--timeout] ip port")
	}

	socket := flag.Args()[0] + ":" + flag.Args()[1]

	return d, socket
}

func WriteData(c net.Conn, sig chan error, wg *sync.WaitGroup) {
	go func() {
		if _, ok := <-sig; !ok {
			os.Stdin.Close()
			fmt.Println("connection interrupted\npress Enter to end the program")
		}
	}()

	r := bufio.NewReader(os.Stdin)
loop:
	for {
		b, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("WriteData finished 1")
				defer c.Close()
				break loop
			}
			if errors.Unwrap(err) == os.ErrClosed {
				fmt.Println("WriteData finished 2")
				defer c.Close()
				break loop
			}
			log.Fatal(err)
		}
		c.Write(b)
	}
	wg.Done()
}

func ReadData(c net.Conn, sig chan error, wg *sync.WaitGroup) {
	r := bufio.NewReader(c)
loop:
	for {
		b, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("ReadData finished 1")
				close(sig)
				break loop
			}

			if errors.Unwrap(err) == net.ErrClosed {
				fmt.Println("ReadData finished 2")
				break loop
			}

			log.Fatal(err)
		}
		os.Stdout.Write(b)
	}
	wg.Done()
}
