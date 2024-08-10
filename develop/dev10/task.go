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
	"time"
)

func main() {
	d, socket := flags()
	conn, err := net.DialTimeout("tcp", socket, *d)
	if err != nil {
		log.Fatal(err)
	}

	sig := make(chan error)
	go WriteData(conn, sig)
	go ReadData(conn, sig)

	time.Sleep(time.Second * 10)

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

func WriteData(c net.Conn, sig chan error) {
	go func() {
		if _, ok := <-sig; !ok {
			os.Stdin.Close()
			fmt.Println("WriteData finished 2")
			fmt.Scanf("\n")
		}
	}()

	r := bufio.NewReader(os.Stdin)
loop:
	for {
		b, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("WriteData finished 1")
				c.Close()
				break loop
			}
			if errors.Unwrap(err) == os.ErrClosed {
				fmt.Println("WriteData finished 2")
				break loop
			}
			log.Fatal(err)
		}
		c.Write(b)
	}
}

func ReadData(c net.Conn, sig chan error) {
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
}
