package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type flags struct {
	k int
	n bool
	r bool
	u bool
}

func main() {

	//Парсинг файла
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	data := make([]string, 0, 5)
	for sc.Scan() {
		data = append(data, sc.Text())
	}

	//Дефолтная сортировка без флагов
	// sort.Strings(data)

	//Сортировка с n-й колонки
	k := 3
	less := func(i, j int) bool {
		arr1 := strings.Fields(data[i])
		arr2 := strings.Fields(data[j])
		
		if k-1 >= len(arr1) {
			return true
		}
		
		if k-1 >= len(arr2) {
			return false
		}

		return arr1[k-1] < arr2[k-1]
	}

	//метод сортирует в зависимости хотим ли мы переставлять i и j элементы
	sort.Slice(data, less)

	for _, val := range data {
		fmt.Println(val)
	}
}
