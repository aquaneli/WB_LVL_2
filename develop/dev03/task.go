package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type flags struct {
	k    *int
	n    *bool
	r    *bool
	u    *bool
	M    *bool
	b    *bool
	c    *bool
	h    *bool
	path []string
}

func main() {
	args := parseFlags()
	data, err := parseFiles(&args)
	if err != nil {
		log.Fatalln(err)
	}
	res := sortStrings(data, args)

	if !*args.c {
		for _, val := range res {
			fmt.Println(val)
		}
	}
}

func sortStrings(data []string, args flags) []string {
	if *args.k > 0 {
		kFlag(&data, args)
	} else if *args.n {
		nFlag(&data)
	} else if *args.u {
		uFlag(&data)
	} else if *args.M {
		mFlag(&data)
	} else if *args.b {
		bFlag(&data)
	} else if *args.c {
		result := cFlag(&data)
		if !result {
			return nil
		}
	} else if *args.h {
		hFlag(&data)
	} else {
		nonFlag(&data)
	}

	if *args.r {
		rFlag(&data)
	}
	return data
}

/* Парсинг флагов */
func parseFlags() flags {
	args := flags{}
	args.k = flag.Int("k", 0, "указание колонки для сортировки")
	args.n = flag.Bool("n", false, "сортировать по числовому значению")
	args.r = flag.Bool("r", false, "сортировать в обратном порядке")
	args.u = flag.Bool("u", false, "не выводить повторяющиеся строки")
	args.M = flag.Bool("M", false, "сортировать по названию месяца")
	args.b = flag.Bool("b", false, "игнорировать хвостовые пробелы")
	args.c = flag.Bool("c", false, "проверять отсортированы ли данные")
	args.h = flag.Bool("h", false, "сортировать по числовому значению с учетом суффиксов")
	flag.Parse()
	return args
}

/* Парсинг файлов */
func parseFiles(args *flags) ([]string, error) {
	data := make([]string, 0, 5)
	ok := false
	pathFiles(args)
	for _, val := range args.path {
		scanFile(val, &data)
		ok = true
	}
	if !ok {
		return []string{}, errors.New("file not specified")
	}
	return data, nil
}

/* Находим все файлы */
func pathFiles(args *flags) {
	for _, val := range os.Args[1:] {
		if val[0] != '-' {
			args.path = append(args.path, val)
		}
	}
}

/* Сканирование файлов */
func scanFile(path string, data *[]string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		*data = append(*data, sc.Text())
	}
}

/* 1.Дефолтная сортировка без флагов */
func nonFlag(data *[]string) {
	sort.Strings(*data)
}

/* 2.Сортировка по k-й колонке, где колонками в строке являются слова разделенные пробелами */
func kFlag(data *[]string, args flags) {
	less := func(i, j int) bool {
		if *args.k < 1 {
			log.Fatal("incorrect value")
		}
		arr1 := strings.Fields((*data)[i])
		arr2 := strings.Fields((*data)[j])
		if *args.k-1 >= len(arr1) {
			return true
		}

		if *args.k-1 >= len(arr2) {
			return false
		}

		return arr1[*args.k-1] < arr2[*args.k-1]
	}
	// метод сортирует в зависимости от того хотим ли мы чтобы i элемент стоял перед j
	sort.Slice(*data, less)
}

/* 3.сортировать по числовому значению */
func nFlag(data *[]string) {
	less := func(i, j int) bool {
		arr1 := strings.Fields((*data)[i])
		arr2 := strings.Fields((*data)[j])
		if len(arr1) == 0 || len(arr2) == 0 {
			return true
		}
		num1 := parseNum(arr1[0])
		num2 := parseNum(arr2[0])
		return num1 < num2
	}
	sort.Slice(*data, less)
}

func parseNum(arr string) float64 {
	i := 0
	for checkPoint := 0; i < len(arr) && (unicode.IsDigit(rune(arr[i])) || arr[i] == '.'); {
		if arr[i] == '.' {
			checkPoint++
		}
		if checkPoint > 1 {
			return 0
		}
		i++
	}
	res, _ := strconv.ParseFloat(arr[:i], 64)
	return res
}

/* 4. делает слайс в обратном порядке */
func rFlag(data *[]string) {
	last := len(*data) - 1
	for i := 0; i < len(*data)/2; i++ {
		(*data)[i], (*data)[last-i] = (*data)[last-i], (*data)[i]
	}
}

/* 5.Убрать дубликаты строк */
func uFlag(data *[]string) {
	m := make(map[string]string, len(*data))
	result := make([]string, 0, len(*data))
	for _, val := range *data {
		_, ok := m[val]
		if !ok {
			m[val] = val
			result = append(result, val)
		}
	}
	*data = result
}

/* 6.Сортировать по названию месяца */
func mFlag(data *[]string) {
	less := func(i, j int) bool {
		arr1 := strings.Fields((*data)[i])
		arr2 := strings.Fields((*data)[j])
		return parseMonth(arr1[0]) < parseMonth(arr2[0])
	}
	sort.Slice(*data, less)
}

func parseMonth(arr string) int {
	switch arr {
	case "January":
		return 1
	case "February":
		return 2
	case "March":
		return 3
	case "April":
		return 4
	case "May":
		return 5
	case "June":
		return 6
	case "July":
		return 7
	case "August":
		return 8
	case "September":
		return 9
	case "October":
		return 10
	case "November":
		return 11
	case "December":
		return 12
	}
	return 0 // Значение по умолчанию, если месяц не найден
}

/* 7.Игнорировать хвостовые пробелы */
func bFlag(data *[]string) {
	less := func(i, j int) bool {
		return strings.TrimSpace((*data)[i]) < strings.TrimSpace((*data)[j])
	}
	sort.Slice(*data, less)
}

/* 8.Проверять отсортированы ли данные */
func cFlag(data *[]string) bool {
	for i := 0; i < len(*data)-1; i++ {
		if (*data)[i] > (*data)[i+1] {
			fmt.Printf("disorder: %s\n", (*data)[i+1])
			return false
		}
	}
	return true
}

/* 9.ортировать по числовому значению с учетом суффиксов */
func hFlag(data *[]string) {
	less := func(i, j int) bool {
		arr1 := strings.Fields((*data)[i])
		arr2 := strings.Fields((*data)[j])
		num1, r1 := searchSuffix(arr1[0])
		num2, r2 := searchSuffix(arr2[0])
		if r1 == r2 {
			return num1 < num2
		}
		return r1 < r2
	}
	sort.Slice(*data, less)
}

/* Проверка есть ли суффик и корректное ли число перед ним */
func searchSuffix(arr string) (float64, int) {
	i := 0
	for checkPoint := 0; i < len(arr) && (unicode.IsDigit(rune(arr[i])) || arr[i] == '.'); {
		if arr[i] == '.' {
			checkPoint++
		}
		if checkPoint > 1 {
			return 0, 0
		}
		i++
	}

	// если числовой части нет, возвращаем 0
	if i == 0 {
		return 0, 0
	}
	num, err := strconv.ParseFloat(arr[:i], 64)
	if err != nil {
		return 0, 0
	}
	return num, parseSuffix(arr, i)
}

/* Обработка суффиксов */
func parseSuffix(arr string, i int) int {
	var suffix rune
	if len(arr) > i {
		suffix = rune(arr[i])
	}
	switch suffix {
	case 'K':
		return 1
	case 'M':
		return 2
	case 'G':
		return 3
	case 'T':
		return 4
	case 'P':
		return 5
	case 'E':
		return 6
	case 'Z':
		return 7
	case 'Y':
		return 8
	}
	return 0
}
