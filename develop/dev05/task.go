package main

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

	buff, indexBuff := parseFile(args)

	if !args.c {
		process(buff, indexBuff, args)
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

	return buff, indexBuff
}

func process(buff [][]string, indexBuff [][]int, args flags) [][]string {

	if args.C > 0 {
		args.A = args.C
		args.B = args.C
	}

	// resBuff := [][]string{}

	for i := range indexBuff {

		for j := 0; j < len(indexBuff[i]); j++ {

			if args.B > 0 {
				num := BFlag(&buff[i], &indexBuff[i], args, i, j)
				trimSlice(&buff[i], &indexBuff[i], num, j)
			}

			if len(indexBuff) > 1 {

				// if !args.n {
				// 	fmt.Printf("%s:%s\n", args.pathFile[i], buff[i][indexBuff[i][j]])
				// } else {
				// 	fmt.Printf("%s:%d:%s\n", args.pathFile[i], numberStr+1, buff[i][indexBuff[i][j]])
				// }

			} else {

				// if !args.n {
				// 	fmt.Printf("%s\n", buff[i][indexBuff[i][j]])
				// } else {
				// 	fmt.Printf("%d:%s\n", numberStr+1, buff[i][indexBuff[i][j]])
				// }

			}

			if args.A > 0 {
				num := AFlag(&buff[i], &indexBuff[i], args, i, j)
				trimSlice(&buff[i], &indexBuff[i], num, j)
			}
		}

	}
}

func trimSlice(buff *[]string, indexBuff *[]int, num, i int) {
	*buff = (*buff)[num:]
	for j := i; j < len(*indexBuff); j++ {
		(*indexBuff)[j] -= num
	}

}

// AFlag печатает в результирующий буффер +N строк после совпадения
func AFlag(buff *[]string, indexBuff *[]int, args flags, indexFile, j int) int {
	val := (*indexBuff)[j]
	num := 0

	if j+1 < len(*indexBuff) {

		if val+args.A < (*indexBuff)[j+1] {
			printLines((*buff)[val+1:val+args.A+1], args, indexFile)
			// if args.A+val+1 < (*indexBuff)[j+1] {
			// 	fmt.Println("--")
			// }
			num = val + args.A + 1
		} else if val+args.A >= (*indexBuff)[j+1] {
			printLines((*buff)[val+1:(*indexBuff)[j+1]], args, indexFile)
			num = (*indexBuff)[j+1]
		}

	} else {

		if val+args.A < len(*buff) {
			printLines((*buff)[val+1:val+args.A+1], args, indexFile)
		} else if val+args.A >= len(*buff) {
			printLines((*buff)[val+1:], args, indexFile)
		}

	}

	return num
}

// BFlag печатает в результирующий буффер +N строк до совпадения
func BFlag(buff *[]string, indexBuff *[]int, args flags, indexFile, j int) int {
	val := (*indexBuff)[j]

	if j == 0 {

		if args.B > val {
			printLines((*buff)[:val], args, indexFile)

		} else if args.B <= val {
			printLines((*buff)[val-args.B:val], args, indexFile)
		}

	} else {

		if args.B >= val {

			if args.C > 0 || (args.A > 0 && args.B > 0) {
				printLines((*buff)[:val], args, indexFile)
			} else {
				printLines((*buff)[1:val], args, indexFile)
			}

		} else if args.B < val {

			// if val-args.B > 0 {
			// 	fmt.Println("--")
			// }
			printLines((*buff)[val-args.B:val], args, indexFile)

		}

	}

	return val
}

func printLines(buff []string, args flags, indexFile int) {
	for _, v := range buff {
		if len(args.pathFile) > 1 {
			fmt.Printf("%s-%s\n", args.pathFile[indexFile], v)
		} else {
			fmt.Printf("%s\n", v)
		}

	}
}
