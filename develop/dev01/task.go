package main

/*
	С помощью библиотеки https://github.com/beevik/ntp мы получаем текущее время ,
	которое будет возвращено нам с сервера и затем выведем его
	Если случится ошибка, ntp.Time неисправленное локальное системное время, то в таком случае
	мы выйдем из программы с ошибкой 1.
*/

import (
	"fmt"
	"os"

	ntp "github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("ntp1.ntp-servers.net")
	if err != nil {
		fmt.Fprint(os.Stderr, "Ошибка:", err)
		os.Exit(1)
	}
	fmt.Println("Точное время:", time)
}
