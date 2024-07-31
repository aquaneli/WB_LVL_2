package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type flags struct {
	f []int
	d string
	s bool
}

func main() {
	args, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}
	getResult(*args)
}

func getResult(args flags) []string {
	resultChan := make(chan string)
	go parseStrings(args, resultChan)
	var results []string
	for result := range resultChan {
		results = append(results, result)
		fmt.Println(result)
	}
	return results
}

func parseStrings(args flags, resultChan chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	sb := strings.Builder{}
	defer close(resultChan)

	for scanner.Scan() {
		line := scanner.Text()

		sepRes := strings.SplitAfter(line, args.d)
		if len(sepRes) == 1 && args.s {
			continue
		}

		lenStr := len(sepRes)

		for _, v := range args.f {
			if v < lenStr {
				sb.WriteString(sepRes[v])
			} else {
				break
			}
		}

		sbLen := sb.Len()
		if sbLen > 0 && sb.String()[sbLen-1:sbLen] == args.d {
			resultChan <- sb.String()[:sbLen-1]
		} else {
			resultChan <- sb.String()[:sbLen]
		}
		sb.Reset()
	}
}

func parseFlags() (*flags, error) {
	fFlag := flag.String("f", "0", "выбрать поля (колонки)")
	dFlag := flag.String("d", "\t", "использовать другой разделитель")
	sFlag := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	if len(flag.Args()) > 0 {
		return nil, errors.New("usage: ./task -f list [-s] [-d delim]")
	}

	if *dFlag == "" || len(*dFlag) > 1 {
		return nil, errors.New("bad delimiter")
	}

	numCols, err := searchReg(*fFlag)
	if err != nil {
		return nil, err
	}

	return &flags{
		f: numCols,
		d: *dFlag,
		s: *sFlag}, nil
}

func searchReg(strReg string) ([]int, error) {
	var result []int
	seen := make(map[int]struct{})

	if matched, _ := regexp.MatchString(`^[\d,-]+$`, strReg); !matched {
		return nil, errors.New("invalid characters in input string")
	}

	re := regexp.MustCompile(`(\d+)(?:-(\d+))?`)

	matches := re.FindAllStringSubmatch(strReg, -1)

	for _, match := range matches {

		start, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, err
		}

		if match[2] != "" {
			end, err := strconv.Atoi(match[2])
			if err != nil {
				return nil, err
			}

			for i := start; i <= end; i++ {
				if _, found := seen[i]; !found {
					result = append(result, i-1)
					seen[i] = struct{}{}
				}
			}
		} else {
			if _, found := seen[start]; !found {
				result = append(result, start-1)
				seen[start] = struct{}{}
			}
		}
	}
	sort.Ints(result)
	if result[0] < 0 {
		return nil, errors.New("column number cannot be less than 1")
	}
	return result, nil
}
