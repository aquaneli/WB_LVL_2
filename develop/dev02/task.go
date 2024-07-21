package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//Обоработать если число будет 0 и обработать escape последовательность

func main() {
	arr := "a4bc2d5e"
	fmt.Println(Unpack(arr))
}

func Unpack(arr string) (string, error) {
	r := make([]rune, 0, 2)
	sb := strings.Builder{}
	check := false
	escape := 0

	for i, val := range arr {

		if val >= '0' && val <= '9' {
			if i == 0 {
				continue
			} else if check {
				return "", errors.New("incorrect string")
			}
		} else {
			check = false
		}

		if val == '\\' {
			escape++
			if len(r) > 0 {
				sb.WriteString(pop(&r))
			}
		} else {
			escape = 0
		}

		if val != '\\' || escape%2 == 0 {
			push(&r, val)
		}

		if len(r) == 2 {
			sb.WriteString(strings.Repeat(pop(&r), number(&r, &check)))
		}
	}

	if len(r) == 1 {
		sb.WriteString(pop(&r))
	}

	return sb.String(), nil
}

// func checkIncorrectString(val *rune, check *bool) bool {

// }

func number(r *[]rune, check *bool) int {
	num := 1
	if len(*r) > 0 && (*r)[0] >= '0' && (*r)[0] <= '9' {
		num, _ = strconv.Atoi(pop(r))
		*check = true
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
