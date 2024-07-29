package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
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

	buff, indexBuff := parseFile(args)

	if !args.c {
		buffRes := process(buff, indexBuff, args)
		for _, val := range buffRes {
			fmt.Println(val)
		}
	} else {
		for _, v := range indexBuff {
			fmt.Println(len(v))
		}

	}

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

func parseFile(args flags) ([][]string, [][]int) {
	buff := make([][]string, len(args.pathFile))
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
			buff[i] = append(buff[i], str)
			if compileReg.MatchString(str) {
				indexBuff[i] = append(indexBuff[i], j)
			}
		}
	}

	return buff, indexBuff
}

func process(buff [][]string, indexBuff [][]int, args flags) []string {

	if args.C > 0 {
		args.A = args.C
		args.B = args.C
	}

	buffRes := []string{}

	for i := range indexBuff {

		for j := 0; j < len(indexBuff[i]); j++ {
			if args.B > 0 {
				resFlag, num := BFlag(&buff[i], &indexBuff[i], j, args)
				buffRes = append(buffRes, resFlag...)
				trimSlice(&buff[i], &indexBuff[i], num, j)
			}

			buffRes = append(buffRes, buff[i][indexBuff[i][j]])

			if args.A > 0 {
				resFlag, num := AFlag(&buff[i], &indexBuff[i], j, args)
				buffRes = append(buffRes, resFlag...)
				trimSlice(&buff[i], &indexBuff[i], num, j)
			}
		}
		
	}
	return buffRes
}

func trimSlice(buff *[]string, indexBuff *[]int, num, i int) {
	*buff = (*buff)[num:]
	for j := i; j < len(*indexBuff); j++ {
		(*indexBuff)[j] -= num
	}

}

// AFlag добавляет в результирующий буффер +N строк после совпадения
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

// BFlag добавляет в результирующий буффер +N строк до совпадения
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
			if args.C > 0 || (args.A > 0 && args.B > 0) {
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
