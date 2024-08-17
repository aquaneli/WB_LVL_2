package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type flags struct {
	A        int
	B        int
	C        int
	c        bool
	i        bool
	v        bool
	F        bool
	n        bool
	pattern  string
	pathFile []string
}

func main() {
	args, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}
	buff, indexBuff, buffStr := parseFile(args)
	indexBuffReserve := getCopyIndexBuff(indexBuff)
	resBuff := process(buff, indexBuff, args)
	printBuff(resBuff, indexBuffReserve, buffStr, args)
}

func getCopyIndexBuff(indexBuff [][]int) [][]int {
	indexBuffReserve := make([][]int, len(indexBuff))
	for i := range indexBuff {
		for range indexBuff[i] {
			indexBuffReserve[i] = make([]int, len(indexBuff[i]))
			copy(indexBuffReserve[i], indexBuff[i])
		}
	}
	return indexBuffReserve
}

func parseFlags() (flags, error) {
	AFlag := flag.Int("A", 0, "указание колонки для сортировки")
	BFlag := flag.Int("B", 0, "сортировать по числовому значению")
	CFlag := flag.Int("C", 0, "сортировать в обратном порядке")
	cFlag := flag.Bool("c", false, "не выводить повторяющиеся строки")
	iFlag := flag.Bool("i", false, "сортировать по названию месяца")
	vFlag := flag.Bool("v", false, "игнорировать хвостовые пробелы")
	FFlag := flag.Bool("F", false, "проверять отсортированы ли данные")
	nFlag := flag.Bool("n", false, "сортировать по числовому значению с учетом суффиксов")
	flag.Parse()

	if len(flag.Args()) < 2 {
		return flags{}, errors.New("file not specified\nusage: ./task [pattern] [-A num] [-B num] [-C num] [-c -i -v -F -n] [files...]")
	}

	patternArg := flag.Arg(0)
	path := flag.Args()[1:]

	return flags{
		A:        *AFlag,
		B:        *BFlag,
		C:        *CFlag,
		c:        *cFlag,
		i:        *iFlag,
		v:        *vFlag,
		F:        *FFlag,
		n:        *nFlag,
		pattern:  patternArg,
		pathFile: path,
	}, nil
}

// Вывод всех результирующих строк строк
func printBuff(resBuff, indexBuffReserve [][]int, buffStr [][]string, args flags) {

	if !args.c {

		for i := range resBuff {
			count := 0
			for j, v := range resBuff[i] {
				if len(resBuff) > 1 {
					printManyFiles(indexBuffReserve, buffStr, args, v, i, &count)
				} else {
					printOneFile(indexBuffReserve, buffStr, args, v, i, &count)
				}
				if j+1 < len(resBuff[i]) && (resBuff[i][j+1]-resBuff[i][j] > 1) {
					fmt.Println("--")
				}
			}
			if len(resBuff[i]) > 0 && i < len(resBuff)-1 {
				fmt.Println("--")
			}
		}

	} else {
		printFilesFlagC(indexBuffReserve, args)
	}
}

// Вывод когда файлов много
func printManyFiles(indexBuffReserve [][]int, buffStr [][]string, args flags, v, i int, count *int) {
	if !args.n {
		if *count < len(indexBuffReserve[i]) && v == indexBuffReserve[i][*count] {
			fmt.Printf("%s:%s\n", args.pathFile[i], buffStr[i][v])
			*count++
		} else {
			fmt.Printf("%s-%s\n", args.pathFile[i], buffStr[i][v])
		}
	} else {
		if *count < len(indexBuffReserve[i]) && v == indexBuffReserve[i][*count] {
			fmt.Printf("%s:%d:%s\n", args.pathFile[i], v+1, buffStr[i][v])
			*count++
		} else {
			fmt.Printf("%s-%d-%s\n", args.pathFile[i], v+1, buffStr[i][v])
		}
	}
}

// Вывод когда файл один
func printOneFile(indexBuffReserve [][]int, buffStr [][]string, args flags, v, i int, count *int) {
	if !args.n {
		fmt.Printf("%s\n", buffStr[i][v])
	} else {
		if v == indexBuffReserve[i][*count] {
			fmt.Printf("%d:%s\n", v+1, buffStr[i][v])
			*count++
		} else {
			fmt.Printf("%d-%s\n", v+1, buffStr[i][v])
		}
	}
}

// Вывод флага -c
func printFilesFlagC(indexBuffReserve [][]int, args flags) {
	for i, v := range indexBuffReserve {
		if len(indexBuffReserve) > 1 {
			fmt.Printf("%s:%d\n", args.pathFile[i], len(v))
		} else {
			fmt.Printf("%d\n", len(v))
		}

	}
}

// Парсинг файлов
func parseFile(args flags) ([][]int, [][]int, [][]string) {
	buff := make([][]int, len(args.pathFile))
	buffStr := make([][]string, len(args.pathFile))
	indexBuff := make([][]int, len(args.pathFile))

	for i, val := range args.pathFile {
		file, err := os.Open(val)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()

		if args.i {
			args.pattern = "(?i)" + args.pattern
		}

		compileReg, err := regexp.Compile(args.pattern)
		if err != nil {
			log.Fatal(err)
		}

		bs := bufio.NewScanner(file)

		for j := 0; bs.Scan(); j++ {
			str := bs.Text()
			buff[i] = append(buff[i], j)
			buffStr[i] = append(buffStr[i], str)
			cmp := true

			if args.F && strings.Compare(str, args.pattern) != 0 {
				cmp = false
			} else if !args.F && !compileReg.MatchString(str) {
				cmp = false
			}

			if args.v && !cmp || !args.v && cmp {
				indexBuff[i] = append(indexBuff[i], j)
			}

		}
	}

	return buff, indexBuff, buffStr
}

// Обработка строк
func process(buff [][]int, indexBuff [][]int, args flags) [][]int {

	if args.C > 0 {
		args.A = args.C
		args.B = args.C
	}

	resBuff := make([][]int, len(buff))

	for i := range indexBuff {
		for j := 0; j < len(indexBuff[i]); j++ {

			if args.B > 0 {
				res, num := BFlag(&buff[i], &indexBuff[i], args, i, j)
				resBuff[i] = append(resBuff[i], res...)
				trimSlice(&buff[i], &indexBuff[i], num, j)
			}

			resBuff[i] = append(resBuff[i], buff[i][indexBuff[i][j]])

			if args.A > 0 {
				res, num := AFlag(&buff[i], &indexBuff[i], args, i, j)
				resBuff[i] = append(resBuff[i], res...)
				trimSlice(&buff[i], &indexBuff[i], num, j)
			}

		}

	}
	return resBuff
}

func trimSlice(buff *[]int, indexBuff *[]int, num, i int) {
	*buff = (*buff)[num:]
	for j := i; j < len(*indexBuff); j++ {
		(*indexBuff)[j] -= num
	}
}

// AFlag печатает в результирующий буффер +N строк после совпадения
func AFlag(buff *[]int, indexBuff *[]int, args flags, indexFile, j int) ([]int, int) {
	val := (*indexBuff)[j]
	resBuff := []int{}
	num := 0

	if j+1 < len(*indexBuff) {

		if val+args.A < (*indexBuff)[j+1] {
			resBuff = append(resBuff, (*buff)[val+1:val+args.A+1]...)
			num = val + args.A + 1

		} else if val+args.A >= (*indexBuff)[j+1] {
			resBuff = append(resBuff, (*buff)[val+1:(*indexBuff)[j+1]]...)
			num = (*indexBuff)[j+1]
		}

	} else {

		if val+args.A < len(*buff) {
			resBuff = append(resBuff, (*buff)[val+1:val+args.A+1]...)

		} else if val+args.A >= len(*buff) {
			resBuff = append(resBuff, (*buff)[val+1:]...)
		}

	}

	return resBuff, num
}

// BFlag печатает в результирующий буффер +N строк до совпадения
func BFlag(buff *[]int, indexBuff *[]int, args flags, indexFile, j int) ([]int, int) {
	val := (*indexBuff)[j]
	resBuff := []int{}

	if j == 0 {

		if args.B > val {
			resBuff = append(resBuff, (*buff)[:val]...)

		} else if args.B <= val {
			resBuff = append(resBuff, (*buff)[val-args.B:val]...)
		}

	} else {

		if args.B >= val {
			if args.C > 0 || (args.A > 0 && args.B > 0) {
				resBuff = append(resBuff, (*buff)[:val]...)

			} else {
				resBuff = append(resBuff, (*buff)[1:val]...)
			}

		} else if args.B < val {
			resBuff = append(resBuff, (*buff)[val-args.B:val]...)

		}

	}

	return resBuff, val
}
