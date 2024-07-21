package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([]byte, 64)
	for {
		_, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Println(string(data))
	}
}
