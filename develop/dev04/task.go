package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	arr := []string{"тко", "тяпка", "кот", "листок", "Пятак", "слиток", "столик", "окт", "окт", "пЯтка", "слово"}
	anagrams := searchAnagram(arr)

	for key, group := range anagrams {
		fmt.Printf("Ключ: %s Анаграммы: %v\n", key, group)
	}
}

/* Поиск анаграмм */
func searchAnagram(arr []string) map[string][]string {
	setAnagrams, uniqKeys := sortAnagram(arr)
	result := buildUniqAnagram(setAnagrams, uniqKeys)
	return result
}

/*
Для каждого уникального первого попавшегося ключа в множестве делаем анаграммы , избавляемся от дублирования и удаляем
множества если там меньше 2 значений
*/

func buildUniqAnagram(setAnagrams map[string]string, uniqKeys map[string]string) map[string][]string {
	result := make(map[string][]string, 1)
	for key, val := range uniqKeys {
		for k, v := range setAnagrams {
			if key == v {
				result[val] = append(result[val], k)
			}
		}
		if len(result[val]) <= 1 {
			delete(result, val)
		}
		sort.Strings(result[val])
	}
	return result
}

/*
Создаем уникальные ключи без дублирования , так же в значение каждого ключа указываем отсортированную строку т.е. множество
которому принадлежит каждый уникальный ключ.
Так же отлавливаем первый попавшийся ключ для множества.
*/
func sortAnagram(arr []string) (map[string]string, map[string]string) {
	setAnagrams := make(map[string]string, 1)
	uniqKeys := make(map[string]string, 1)

	for _, val := range arr {
		lowWord := strings.ToLower(val)
		sortedWord := sortChars(lowWord)
		setAnagrams[lowWord] = sortedWord
		uniqKeys[sortedWord] = lowWord
	}
	return setAnagrams, uniqKeys
}

/* Сортировка символов в слове */
func sortChars(lowWord string) string {
	sliceChar := strings.Split(lowWord, "")
	sort.Strings(sliceChar)
	return strings.Join(sliceChar, "")
}
