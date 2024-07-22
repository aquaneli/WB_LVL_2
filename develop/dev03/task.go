package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type flags struct {
	k int
	n bool
	r bool
	u bool
}

func main() {
	var path string
	_, err := fmt.Scan(&path)
	if err != nil {
		log.Fatal(err)
	}
	data := parsFile(path)

	args := flags{}
	args.k = *flag.Int("k", 0, "указание колонки для сортировки")
	args.n = *flag.Bool("n", false, "сортировать по числовому значению")
	args.r = *flag.Bool("r", false, "сортировать в обратном порядке")
	args.u = *flag.Bool("u", false, "не выводить повторяющиеся строки")

	/* Сортировка по названию месяца */

	for _, val := range data {
		fmt.Println(val)
	}
}

/* Парсинг файла */
func parsFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	data := make([]string, 0, 5)
	for sc.Scan() {
		data = append(data, sc.Text())
	}
	return data
}

/* 1.Дефолтная сортировка без флагов */
func nonFlag(data *[]string) {
	sort.Strings(*data)
}

/* 2.Сортировка по k-й колонке, где колонками в строке являются слова разделенные пробелами */
func KFlag(data *[]string) {
	k := 3
	less := func(i, j int) bool {
		if k < 1 {
			log.Fatal("incorrect value")
		}
		arr1 := strings.Fields((*data)[i])
		arr2 := strings.Fields((*data)[j])

		if k-1 >= len(arr1) {
			return true
		}

		if k-1 >= len(arr2) {
			return false
		}

		return arr1[k-1] < arr2[k-1]
	}
	//метод сортирует в зависимости от того хотим ли мы чтобы i элемент стоял перед j
	sort.Slice(*data, less)
}

/* 3.сортировать по числовому значению */
func NFlag(data *[]string) {
	less := func(i, j int) bool {
		arr1 := strings.Fields((*data)[i])
		arr2 := strings.Fields((*data)[j])

		if len(arr1) == 0 || len(arr2) == 0 {
			return true
		}

		fl1, err1 := strconv.ParseFloat(arr1[0], 32)
		if err1 != nil {
			return true
		}

		fl2, err2 := strconv.ParseFloat(arr2[0], 32)
		if err2 != nil {
			return false
		}

		return fl1 < fl2
	}
	sort.Slice(*data, less)
}

/* 5.Убрать дубликаты строк */
func UFlag(data *[]string) []string {
	m := make(map[string]string, len(*data))
	result := make([]string, 0, len(*data))
	for _, val := range *data {
		_, ok := m[val]
		if !ok {
			m[val] = val
		}
	}
	for _, val := range m {
		result = append(result, val)
	}
	sort.Strings(result)
	return result
}
