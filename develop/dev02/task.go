package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

//Обоработать если число будет 0 и обработать escape последовательность

func main() {
	arr := "qwe\\45"
	// fmt.Println(Revers(arr))
	fmt.Println(Unpack(arr))
}

func Unpack(arr string) string {
	arr = Revers(arr)
	num := 1
	check, escaped := false, false
	sb := strings.Builder{}

	for _, val := range arr {
		if val >= '0' && val <= '9' && !escaped {
			num, _ = strconv.Atoi(string(val))
			if check {
				log.Fatal("dublicate num")
			}
			check = true
		} else {
			if val == '\\' && !escaped {
				escaped = true
			} else {
				sb.WriteString(strings.Repeat(string(val), num))
				num = 1
				check, escaped = false, false

			}
		}
	}
	return Revers(sb.String())
}

func Revers(str string) string {
	strRune := []rune(str)
	r := make([]rune, len(strRune))
	for i, val := range strRune {
		r[len(r)-i-1] = val
	}
	return string(r)
}
