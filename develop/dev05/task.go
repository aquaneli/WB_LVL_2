package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

type flags struct {
	A int
	B int
	C int
	c int
	i bool
	v bool
	F bool
	n bool
}

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	/* ----- */

	args := parseFlags()

	/* ----- */

	bs := bufio.NewScanner(file)
	buff := []string{}
	for bs.Scan() {
		buff = append(buff, bs.Text())
	}

	/* ----- */

	buffRes := []string{}

	reg := "0.1"
	compileReg, err := regexp.Compile(reg)
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	for i, val := range buff {

		buffRes = append(buffRes, val)
		cr := compileReg.MatchString(val)

		if !cr {
			count++
		}
		if cr {
			if args.B < count {
				count = args.B
			}
			copy(buffRes[len(buffRes)-1-count:], buff[i-count:])
			count = 0
		}

	}

	for _, val := range buffRes {
		fmt.Println(val)
	}

}

func AFlag(cr bool, buff *[]string, countAfter *int, bs *bufio.Scanner, args flags) {
	if cr {
		*buff = append(*buff, (*bs).Text())
		*countAfter = 0
	}

	if !cr && len(*buff) > 0 {
		if *countAfter < args.A {
			*buff = append(*buff, (*bs).Text())
		} else if *countAfter == args.A {
			*buff = append(*buff, "--")
		}
		*countAfter++
	}
}

func BFlag(cr bool, buff *[]string, buffBefore *[]string, countBefore *int, bs *bufio.Scanner, args flags) {
	if !cr {
		*buffBefore = append(*buffBefore, (*bs).Text())
		*countBefore++
	}

	if cr {
		*buffBefore = append(*buffBefore, (*bs).Text())
		if *countBefore > args.B {
			if len(*buff) > 0 {
				*buff = append(*buff, "--")
			}
			*countBefore = args.B
		}

		l := len(*buffBefore) - 1 - *countBefore
		*buff = append(*buff, (*buffBefore)[l:]...)
		*countBefore = 0
		*buffBefore = (*buffBefore)[:0]
	}
}

func CFlag() {

}

func parseFlags() flags {
	AFlag := flag.Int("A", 0, "указание колонки для сортировки")
	BFlag := flag.Int("B", 0, "сортировать по числовому значению")
	CFlag := flag.Int("C", 0, "сортировать в обратном порядке")
	cFlag := flag.Int("c", 0, "не выводить повторяющиеся строки")
	iFlag := flag.Bool("i", false, "сортировать по названию месяца")
	vFlag := flag.Bool("v", false, "игнорировать хвостовые пробелы")
	FFlag := flag.Bool("F", false, "проверять отсортированы ли данные")
	nFlag := flag.Bool("n", false, "сортировать по числовому значению с учетом суффиксов")
	flag.Parse()

	if flag.NFlag() > 1 {
		log.Fatalln("Only 1 flag can be used")
	}

	return flags{
		A: *AFlag,
		B: *BFlag,
		C: *CFlag,
		c: *cFlag,
		i: *iFlag,
		v: *vFlag,
		F: *FFlag,
		n: *nFlag,
	}
}
