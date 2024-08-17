package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	arr := "\\\\4\\5"
	fmt.Println(Unpack(arr))
}

// Unpack распаковывает символы
func Unpack(arr string) (string, error) {
	r := make([]rune, 0, 2)
	sb := strings.Builder{}
	check := false
	escape := 0
	// Проходим в строке по рунам
	for i, val := range arr {
		// Находим ошибку если идут 2 числа подряд, где ни один из них не является escape символом
		if val >= '0' && val <= '9' {
			if i == 0 {
				check = true
				continue
			} else if check {
				return "", errors.New("incorrect string")
			}
		} else {
			check = false
		}
		// Проверка escape последовательности
		checkEscape(&r, val, &escape, &sb)

		// Пушим в результирующий String Builder если в стеке есть 2 руны */
		if len(r) == 2 {
			sb.WriteString(strings.Repeat(pop(&r), number(&r, &check)))
		}
	}
	// Если остался последний символ то пушим в String Builder
	if len(r) == 1 {
		sb.WriteString(pop(&r))
	}
	return sb.String(), nil
}

func checkEscape(r *[]rune, val rune, escape *int, sb *strings.Builder) {
	if val == '\\' {
		*escape++
		if len(*r) > 0 {
			(*sb).WriteString(pop(r))
		}
	} else {
		*escape = 0
	}
	if val != '\\' || (*escape)%2 == 0 {
		push(r, val)
	}
}

// Функция для поиска количества дублирования символа, если для
// дублирования стоит число 0, то символ вообще не будет выводиться
func number(r *[]rune, check *bool) int {
	num := 1
	if len(*r) > 0 && (*r)[0] >= '0' && (*r)[0] <= '9' {
		num, _ = strconv.Atoi(pop(r))
		*check = true
	}
	return num
}

// Запушить в очередь
func push(queue *[]rune, r rune) {
	*queue = append(*queue, r)
}

// Удалить из очереди
func pop(queue *[]rune) string {
	r := (*queue)[0]
	if len(*queue) > 1 {
		copy((*queue)[:len(*queue)-1], (*queue)[1:])
	}
	*queue = (*queue)[:len(*queue)-1]
	return string(r)
}
