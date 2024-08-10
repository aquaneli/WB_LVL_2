package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	d, socket := flags()
	conn, err := net.DialTimeout("tcp", socket, *d)
	if err != nil {
		log.Fatal(err)
	}

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	defer close(sigChan)

	ctx, cancel := context.WithCancel(context.Background())

	go WriteData(conn, ctx, sigChan)
	go ReadData(conn, ctx, sigChan)

	select {
	case <-sigChan:
		cancel()
	}

	time.Sleep(time.Second * 3)

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

func WriteData(c net.Conn, ctx context.Context, sigChan chan os.Signal) {
	r := bufio.NewReader(os.Stdin)
loop:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteData finished")
			break loop
		default:
			b, err := r.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					fmt.Println("WriteData finished")
					c.Close()
					break loop
				}
				log.Fatal(err)
			}
			c.Write(b)
		}
	}
}

func ReadData(c net.Conn, ctx context.Context, sigChan chan os.Signal) {
	r := bufio.NewReader(c)
loop:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ReadData finished")
			break loop
		default:
			b, err := r.ReadBytes('\n')
			if err != nil {
				if errors.Unwrap(err) == net.ErrClosed {
					fmt.Println("ReadData finished")
					break loop
				}
				log.Fatal(err)
			}
			os.Stdout.Write(b)
		}

	}
}
