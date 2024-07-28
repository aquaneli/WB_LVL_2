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

	// fmt.Println(indexBuff)

	/* ----- */

	buffRes := []string{}
	if err != nil {
		log.Fatal(err)
	}

	if args.C > 0 {
		args.A = args.C
		args.B = args.C
	}

	count := 0

	for i := 0; i < len(indexBuff); i++ {

		if args.B > 0 {
			resFlag, num := BFlag(&buff, &indexBuff, i, args)
			if resFlag != nil {
				buffRes = append(buffRes, resFlag...)
			}

			buff = buff[num:]

			for j := count; j < len(indexBuff); j++ {
				indexBuff[j] -= num
			}

			// fmt.Println(indexBuff, buff)
		}

		buffRes = append(buffRes, buff[indexBuff[i]])

		//отдельно А работает
		if args.A > 0 {
			resFlag, num := AFlag(&buff, &indexBuff, i, args)
			buffRes = append(buffRes, resFlag...)

			buff = buff[num:]

			for j := count; j < len(indexBuff); j++ {
				indexBuff[j] -= num
			}

			//если index элемент < 0 то элемента уже не существует
			// fmt.Println(indexBuff)
		}
		count++
	}

	/* ----- */

	fmt.Println(indexBuff)
	for _, val := range buffRes {
		fmt.Println(val)
	}

}

func AFlag(buff *[]string, indexBuff *[]int, i int, args flags) ([]string, int) {
	buffRes := []string{}

	//val текущий индекс
	val := (*indexBuff)[i]

	//если не последний индекс
	if i+1 < len(*indexBuff) {

		//если от текщего элемента нам нужно взять до следующего индекса но не доходим до него и причем текущий индекс мы не берем
		if val+args.A < (*indexBuff)[i+1] {

			//взяли срез со следующего элемента после индекса и до args.A+val+1 тут 1 потому что нужно взять след элмент
			buffRes = append(buffRes, (*buff)[val+1:val+args.A+1]...)
			if args.A+val+1 < (*indexBuff)[i+1] && args.B == 0 {
				buffRes = append(buffRes, "--")
			}

			//возвращаем индекс на сколько мы обрезали до куда а +1 потому что мы обрезали и сам индекс
			return buffRes, val + args.A + 1

			//если от текщего элемента нам нужно взять до следующего индекса но доходим до него и причем текущий индекс мы не берем
		} else if val+args.A >= (*indexBuff)[i+1] {

			//взяли срез со следующего элемента после индекса и до (*indexBuff)[i+1] т.е следующего индекса тут +1 нет потому что мы не хотим брать следующий индекс

			buffRes = append(buffRes, (*buff)[val+1:(*indexBuff)[i+1]]...)

			//возвращаем индекс на сколько мы обрезали слайс т.е. тут верну индекс следующего элемента потому что он является крайним индексом
			return buffRes, (*indexBuff)[i+1]

		}

		//если последний индекс
	} else {

		//если мы хотим обрезать от текущего элемента и до val+args.A тут + 1 потоу что нужно взять следующий элемент но элементов осталось больше чем нам нужно
		if val+args.A < len(*buff) {
			buffRes = append(buffRes, (*buff)[val+1:val+args.A+1]...)

			//если мы хотим обрезать от текущего элемента и до val+args.A тут + 1 потоу что нужно взять следующий элемент но но элементов осталось меньше чем нам нужно
		} else if val+args.A >= len(*buff) {
			buffRes = append(buffRes, (*buff)[val+1:]...)
		}

	}
	return buffRes, 0
}

func BFlag(buff *[]string, indexBuff *[]int, i int, args flags) ([]string, int) {
	buffRes := []string{}
	val := (*indexBuff)[i]

	//берем срез для самого первого элемента
	if i == 0 {
		//если мы хотим взять элементов больше чем есть сверху
		if args.B > val {

			//взяли от самго начала и до индекса но не сам индекс
			buffRes = append(buffRes, (*buff)[:val]...)

			//val на сколько элементов нужно сократить  buff т.е val будет указывать на индекс
			return buffRes, val
		} else if args.B <= val {
			buffRes = append(buffRes, (*buff)[val-args.B:val]...)

			//тут берем не val-args.B а просто val так как теперь это будет верхушкой
			return buffRes, val
		}

		//теперь тут берем для всех остальных элементов
	} else {

		if args.B >= val {
			if args.C > 0 {
				buffRes = append(buffRes, (*buff)[:val]...)
				return buffRes, val
			} else {
				buffRes = append(buffRes, (*buff)[1:val]...)
				return buffRes, val
			}

		} else if args.B < val {
			buffRes = append(buffRes, "--")
			buffRes = append(buffRes, (*buff)[val-args.B:val]...)
			return buffRes, val
		}
	}

	return buffRes, val
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
