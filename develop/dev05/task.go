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

	bs := bufio.NewScanner(file)

	reg := "и"
	compileReg, err := regexp.Compile(reg)
	if err != nil {
		log.Fatal(err)
	}

	buff := []string{}
	buffBefore := []string{}

	args := parseFlags()

	countAfter := 0
	countBefore := 0

	for bs.Scan() {

		cr := compileReg.MatchString(bs.Text())

		if args.A > 0 {
			AFlag(cr, &buff, &countAfter, bs, args)
		} else if args.B > 0 {
			BFlag(cr, &buff, &buffBefore, &countBefore, bs, args)
		} else if args.C > 0 {

		} else if cr {

			buff = append(buff, bs.Text())

		}

	}

	if args.A > 0 && len(buff) > 0 && buff[len(buff)-1] == "--" && countAfter >= args.A {
		buff = buff[:len(buff)-1]
	}

	for _, val := range buff {
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
		buffBefore = nil
	}
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
