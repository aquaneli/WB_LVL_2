package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	arr := []string{"тко", "тяпка", "кот", "листок", "Пятак", "слиток", "Cлиток", "СЛИТОК", "котилс", "столик", "окт", "окт", "пЯтка", "слово"}
	anagrams := searchAnagram(arr)

	for key, group := range anagrams {
		fmt.Printf("Ключ: %s Анаграммы: %v\n", key, group)
	}
}

// Поиск анаграмм
func searchAnagram(arr []string) map[string][]string {
	setAnagrams, uniqKeys := sortAnagram(arr)
	result := buildUniqAnagram(setAnagrams, uniqKeys)
	return result
}

// Для каждого уникального первого попавшегося ключа в множестве делаем анаграммы ,
// избавляемся от дублирования и добавляем в мапу если больше 1 анаграммы
func buildUniqAnagram(setAnagrams map[string]string, uniqKeys map[string]int) map[string][]string {
	result := make(map[string][]string, 1)
	for uKey, uVal := range uniqKeys {
		uniqStr := make([]string, 0, uVal)
		for key, val := range setAnagrams {
			if uKey == val {
				uniqStr = append(uniqStr, key)
			}
		}
		if len(uniqStr) > 1 {
			sort.Strings(uniqStr)
			result[uniqStr[0]] = uniqStr
		}
	}
	return result
}

// Создаем уникальные ключи без дублирования , так же в значение каждого ключа указываем отсортированную строку т.е. множество
// которому принадлежит каждый уникальный ключ.
// Так же отлавливаем первый попавшийся ключ для множества.
func sortAnagram(arr []string) (map[string]string, map[string]int) {
	setAnagrams := make(map[string]string, 1)
	uniqKeys := make(map[string]int, 1)

	for _, val := range arr {
		lowWord := strings.ToLower(val)
		sortedWord := sortChars(lowWord)
		setAnagrams[lowWord] = sortedWord
		if _, ok := uniqKeys[sortedWord]; !ok {
			uniqKeys[sortedWord] = 1
		} else {
			uniqKeys[sortedWord]++
		}
	}
	return setAnagrams, uniqKeys
}

// Сортировка символов в слове
func sortChars(lowWord string) string {
	sliceChar := strings.Split(lowWord, "")
	sort.Strings(sliceChar)
	return strings.Join(sliceChar, "")
}
