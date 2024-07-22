package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	/* 1.Дефолтная сортировка без флагов */
	// sort.Strings(data)

	/* 2.Сортировка по n-й колонке */
	// k := 3
	// less := func(i, j int) bool {
	// 	if k < 1 {
	// 		log.Fatal("incorrect value")
	// 	}
	// 	arr1 := strings.Fields(data[i])
	// 	arr2 := strings.Fields(data[j])

	// 	if k-1 >= len(arr1) {
	// 		return true
	// 	}

	// 	if k-1 >= len(arr2) {
	// 		return false
	// 	}

	// 	return arr1[k-1] < arr2[k-1]
	// }
	//метод сортирует в зависимости хотим ли мы чтобы i элемент стоял перед j
	// sort.Slice(data, less)

	/* 3.сортировать по числовому значению */
	// less := func(i, j int) bool {
	// 	arr1 := strings.Fields(data[i])
	// 	arr2 := strings.Fields(data[j])

	// 	if len(arr1) == 0 || len(arr2) == 0 {
	// 		return true
	// 	}

	// 	fl1, err1 := strconv.ParseFloat(arr1[0], 32)
	// 	if err1 != nil {
	// 		return true
	// 	}

	// 	fl2, err2 := strconv.ParseFloat(arr2[0], 32)
	// 	if err2 != nil {
	// 		return false
	// 	}

	// 	return fl1 < fl2
	// }
	// sort.Slice(data, less)

	/* 5.Убрать дубликаты строк */
	undub := func() []string {
		m := make(map[string]string, len(data))
		result := make([]string, 0, len(data))
		for _, val := range data {
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
	}()

	data = undub

	for _, val := range data {
		fmt.Println(val)
	}

}
