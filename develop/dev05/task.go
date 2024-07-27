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

	reg := "0"
	compileReg, err := regexp.Compile(reg)

	/* ----- */

	bs := bufio.NewScanner(file)

	buff := []string{}
	indexBuff := []int{}

	for i := 0; bs.Scan(); i++ {
		str := bs.Text()
		buff = append(buff, str)
		if compileReg.MatchString(str) {
			indexBuff = append(indexBuff, i)
		}
	}

	/* ----- */

	buffRes := []string{}
	if err != nil {
		log.Fatal(err)
	}

	for i, val := range indexBuff {
		// fmt.Println(val)

		if args.B > 0 {

			if i-1 >= 0 {

				if val-args.B > indexBuff[i-1] {
					if val-(args.B+1) > indexBuff[i-1] {
						buffRes = append(buffRes, "--")
					}
					buffRes = append(buffRes, buff[val-args.B:val]...)
				} else {
					buffRes = append(buffRes, buff[indexBuff[i-1]+1:val]...)
				}

			} else {

				if args.B > val {
					buffRes = append(buffRes, buff[:val]...)
				} else {
					buffRes = append(buffRes, buff[val-args.B:val]...)
				}

			}

		}

		buffRes = append(buffRes, buff[val])

		if args.A > 0 {

			if i+1 < len(indexBuff) {

				if val+args.A+1 < indexBuff[i+1] {
					buffRes = append(buffRes, buff[val+1:val+args.A+1]...)
					buffRes = append(buffRes, "--")
				} else {
					buffRes = append(buffRes, buff[val+1:indexBuff[i+1]]...)
				}

			} else {

				if val+args.A+1 < len(buff) {
					buffRes = append(buffRes, buff[val+1:val+args.A+1]...)
				} else {
					buffRes = append(buffRes, buff[val+1:]...)
				}

			}

		}

	}

	/* ----- */

	for _, val := range buffRes {
		fmt.Println(val)
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

	// if flag.NFlag() > 1 {
	// 	log.Fatalln("Only 1 flag can be used")
	// }

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
