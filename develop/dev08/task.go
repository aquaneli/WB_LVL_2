package main

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	process "github.com/shirou/gopsutil/process"
)

func main() {
	unixShellUtil()
}

// Запуск программы в бесконечном цикле
func unixShellUtil() {
	bs := bufio.NewScanner(os.Stdin)
	for {
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
		dir := strings.Split(currentDir, "/")
		fmt.Printf("util$ %s %c ", dir[len(dir)-1], '%')

		if bs.Scan() {
			cmdAll := strings.Split(bs.Text(), "|")

			for _, str := range cmdAll {
				cmd := strings.Fields(str)
				switch cmd[0] {
				case "cd":
					cd(cmd)
				case "pwd":
					pwd()
				case "echo":
					echo(cmd)
				case "kill":
					kill(cmd)
				case "ps":
					ps()
				case "exit":
					return
				default:
					fmt.Printf("command not found: %s\n", cmd[0])
				}
			}

		} else {
			break
		}

	}
}

// Перемещение по директориям
func cd(cmd []string) {
	if len(cmd) == 1 {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		err = os.Chdir(home)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	} else if len(cmd) == 2 {
		err := os.Chdir(cmd[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	} else {
		fmt.Println("cd: too many arguments")
	}
}

// Вывод пути
func pwd() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(currentDir)
}

// Вывести строку в stdout
func echo(cmd []string) {
	if len(cmd) > 1 {
		fmt.Println(strings.Join(cmd[1:], " "))
	} else {
		fmt.Println()
	}

}

// Убить процесс
func kill(cmd []string) {
	if len(cmd) > 1 {
		arr := make([]int, 0, len(cmd))
		for _, v := range cmd[1:] {
			pid, err := strconv.Atoi(v)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			arr = append(arr, pid)
		}

		for _, v := range arr {
			proc, err := os.FindProcess(v)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			err = proc.Kill()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			fmt.Println("Process killed")
		}

	}
}

// Вывести pid процессы
func ps() {
	processes, err := process.Processes()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Printf("%3s\t%s\n", "PID", "NAME")
	for _, p := range processes {
		name, err := p.Name()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		fmt.Printf("%3d\t%s\t\n", p.Pid, name)
	}
}
