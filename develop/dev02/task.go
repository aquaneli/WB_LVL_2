package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

//Обоработать если число будет 0 и обработать escape последовательность

func main() {
	arr := "䙐0䨸丠刈嗰4"
	num := -1
	sb := strings.Builder{}
	for _, val := range arr {
		if val >= '1' && val <= '9' {
			cnt, _ := strconv.Atoi(string(val))
			r := []rune(sb.String())
			if num == -1 && len(r) > 0 {
				num = cnt
				sb.WriteString(strings.Repeat(string(r[len(r)-1]), cnt-1))
			} else {
				log.Fatalf("incorrect data")
			}
		} else {
			num = -1
			sb.WriteRune(val)
		}
	}
	fmt.Println(sb.String())
}

func Revers(str string) string {
	r := make([]rune, len(str))
	for i, val := range str {
		r[len(str)-1-i] = val
	}
	return string(r)
}
