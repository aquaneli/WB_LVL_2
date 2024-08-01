package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	unixShellUtil()
}

func unixShellUtil() {
	bs := bufio.NewScanner(os.Stdin)

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("cd: could not get current directory: %v\n", err)
	}
	fmt.Printf("util$ %s %c ", currentDir, '%')

	for bs.Scan() {

		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatalf("cd: could not get current directory: %v\n", err)
		}
		fmt.Printf("util$ %s %c ", currentDir, '%')

		if bs.Text()[:2] == "cd" {
			cd(strings.Join(strings.Split(bs.Text(), " ")[1:], "\\"))
		}

	}
}

func cd(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println("cd: directory not found: %s\n", dir)
	}
}
