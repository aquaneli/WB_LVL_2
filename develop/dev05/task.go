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
	c bool
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

	args := parseFlags()
	buff, indexBuff := parseFile(file, args)
	buffRes := []string{}

	if !args.c {
		buffRes = process(buff, indexBuff, args)
		for _, val := range buffRes {
			fmt.Println(val)
		}
	} else {
		fmt.Println(len(indexBuff))
	}

}

func parseFlags() flags {
	AFlag := flag.Int("A", 0, "указание колонки для сортировки")
	BFlag := flag.Int("B", 0, "сортировать по числовому значению")
	CFlag := flag.Int("C", 0, "сортировать в обратном порядке")
	cFlag := flag.Bool("c", false, "не выводить повторяющиеся строки")
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

func process(buff []string, indexBuff []int, args flags) []string {

	if args.C > 0 {
		args.A = args.C
		args.B = args.C
	}

	buffRes := []string{}

	for i := 0; i < len(indexBuff); i++ {
		if args.B > 0 {
			resFlag, num := BFlag(&buff, &indexBuff, i, args)
			buffRes = append(buffRes, resFlag...)
			trimSlice(&buff, &indexBuff, num, i)
		}

		buffRes = append(buffRes, buff[indexBuff[i]])

		if args.A > 0 {
			resFlag, num := AFlag(&buff, &indexBuff, i, args)
			buffRes = append(buffRes, resFlag...)
			trimSlice(&buff, &indexBuff, num, i)
		}
	}
	return buffRes
}

func parseFile(file *os.File, args flags) ([]string, []int) {
	reg := ""
	if len(os.Args) > 2 {
		reg = os.Args[2]
	}
	if args.i {
		reg = "(?i)" + reg
	}
	compileReg, err := regexp.Compile(reg)

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

	if err != nil {
		log.Fatal(err)
	}

	return buff, indexBuff
}

func trimSlice(buff *[]string, indexBuff *[]int, num, i int) {
	*buff = (*buff)[num:]
	for j := i; j < len(*indexBuff); j++ {
		(*indexBuff)[j] -= num
	}

}

func AFlag(buff *[]string, indexBuff *[]int, i int, args flags) ([]string, int) {
	buffRes := []string{}
	val := (*indexBuff)[i]
	num := 0

	if i+1 < len(*indexBuff) {

		if val+args.A < (*indexBuff)[i+1] {
			buffRes = append(buffRes, (*buff)[val+1:val+args.A+1]...)
			if args.A+val+1 < (*indexBuff)[i+1] && args.B == 0 {
				buffRes = append(buffRes, "--")
			}
			num = val + args.A + 1

		} else if val+args.A >= (*indexBuff)[i+1] {
			buffRes = append(buffRes, (*buff)[val+1:(*indexBuff)[i+1]]...)
			num = (*indexBuff)[i+1]
		}

	} else {

		if val+args.A < len(*buff) {
			buffRes = append(buffRes, (*buff)[val+1:val+args.A+1]...)
		} else if val+args.A >= len(*buff) {
			buffRes = append(buffRes, (*buff)[val+1:]...)
		}

	}
	return buffRes, num
}

func BFlag(buff *[]string, indexBuff *[]int, i int, args flags) ([]string, int) {
	buffRes := []string{}
	val := (*indexBuff)[i]

	if i == 0 {
		if args.B > val {
			buffRes = append(buffRes, (*buff)[:val]...)

		} else if args.B <= val {
			buffRes = append(buffRes, (*buff)[val-args.B:val]...)
		}
	} else {
		if args.B >= val {
			if args.C > 0 {
				buffRes = append(buffRes, (*buff)[:val]...)
			} else {
				buffRes = append(buffRes, (*buff)[1:val]...)
			}
		} else if args.B < val {
			if args.B+1 < val {
				buffRes = append(buffRes, "--")
			}
			buffRes = append(buffRes, (*buff)[val-args.B:val]...)
		}
	}

	return buffRes, val
}
