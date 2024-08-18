package main

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

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
	go writeData(conn, sig, &wg)
	go readData(conn, sig, &wg)

	wg.Wait()
}

// Парсинг флагов
func flags() (*time.Duration, string) {
	d := flag.Duration("timeout", time.Second*10, "server connection timeout")
	flag.Parse()

	if len(flag.Args()) != 2 {
		log.Fatal("usage: ./task [--timeout] ip port")
	}

	socket := flag.Args()[0] + ":" + flag.Args()[1]

	return d, socket
}

// Считывать данные из stdin и отправлять их на сервер до нажатия ctrl+D
func writeData(c net.Conn, sig chan error, wg *sync.WaitGroup) {
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

// Считать данные с обработанных данных сервера и отправить их в stdout
func readData(c net.Conn, sig chan error, wg *sync.WaitGroup) {
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
