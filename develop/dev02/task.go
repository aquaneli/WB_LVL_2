package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Обоработать если число будет 0 и обработать escape последовательность

func main() {
	arr := "\\\\\\5"
	fmt.Println(Unpack(arr))
}

func Unpack(arr string) string {
	r := make([]rune, 0, 2)
	sb := strings.Builder{}

	for _, val := range arr {
		push(&r, val)
		if val == '\\' {
			continue
		} else if len(r) > 1 && r[len(r)-2] == '\\' {
			for i := 0; i < len(r); i++ {

			}
		}

		if len(r) == 2 {
			sb.WriteString(strings.Repeat(pop(&r), number(&r)))
		}
	}

	if len(r) == 1 {
		sb.WriteString(pop(&r))
	}

	return sb.String()
}

func number(r *[]rune) int {
	num := 1
	if len(*r) > 0 && (*r)[0] >= '0' && (*r)[0] <= '9' {
		num, _ = strconv.Atoi(pop(r))
	}
	return num
}

func push(queue *[]rune, r rune) {
	*queue = append(*queue, r)
}

func pop(queue *[]rune) string {
	r := (*queue)[0]
	if len(*queue) > 1 {
		copy((*queue)[:len(*queue)-1], (*queue)[1:])
	}
	*queue = (*queue)[:len(*queue)-1]
	return string(r)
}
