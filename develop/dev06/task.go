package main

import (
	"bufio"
	"errors"
	"flag"
	"log"
	"os"
	"regexp"
	"strconv"
)

type cutArgs struct {
	args flags
	path []string
}

type flags struct {
	f []int
	d string
	s bool
}

func main() {
	cA, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}
	parseFiles(cA)

}

func parseFiles(cA *cutArgs) error {
	str := make([][]string, len(cA.path))
	for i, val := range cA.path {
		file, err := os.Open(val)
		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			str[i] = append(str[i])
		}
	}
	return nil
}

func parseFlags() (*cutArgs, error) {
	fFlag := flag.String("f", "", "выбрать поля (колонки)")
	dFlag := flag.String("d", "\t", "использовать другой разделитель")
	sFlag := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	if len(flag.Args()) < 2 {
		return nil, errors.New("file not specified\nusage: ")
	}

	numCols, err := searchReg(*fFlag)
	if err != nil {
		return nil, err
	}

	return &cutArgs{
		args: flags{
			f: numCols,
			d: *dFlag,
			s: *sFlag},
		path: os.Args[1:],
	}, nil
}

func searchReg(strReg string) ([]int, error) {
	var result []int
	seen := make(map[int]struct{})

	// Проверка на допустимые символы (только цифры, запятые и дефисы)
	if matched, _ := regexp.MatchString(`^[\d,-]+$`, strReg); !matched {
		return nil, errors.New("invalid characters in input string")
	}

	// Регулярное выражение для обработки диапазонов и отдельных чисел
	re := regexp.MustCompile(`(\d+)(?:-(\d+))?`)

	// Находим все совпадения
	matches := re.FindAllStringSubmatch(strReg, -1)

	for _, match := range matches {
		// Начальное значение диапазона или одиночное число
		start, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, err
		}

		if match[2] != "" {
			// Если это диапазон
			end, err := strconv.Atoi(match[2])
			if err != nil {
				return nil, err
			}
			// Добавляем все числа из диапазона
			for i := start; i <= end; i++ {
				if _, found := seen[i]; !found {
					result = append(result, i)
					seen[i] = struct{}{}
				}
			}
		} else {
			// Если это одиночное число
			if _, found := seen[start]; !found {
				result = append(result, start)
				seen[start] = struct{}{}
			}
		}
	}
	return result, nil
}
