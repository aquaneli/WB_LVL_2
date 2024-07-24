package main

import (
	"reflect"
	"testing"
)

func TestAnagram1(t *testing.T) {
	arr := []string{"тко", "тяпка", "кот", "листок", "Пятак", "слиток", "Cлиток", "СЛИТОК", "котилс", "столик", "окт", "окт", "пЯтка", "слово"}
	anagrams := searchAnagram(arr)
	expected := map[string][]string{
		"кот":    {"кот", "окт", "тко"},
		"пятак":  {"пятак", "пятка", "тяпка"},
		"котилс": {"котилс", "листок", "слиток", "столик"},
	}

	if !reflect.DeepEqual(anagrams, expected) {
		t.Errorf("The anagram is incorrect")
	}
}

func TestAnagram2(t *testing.T) {
	arr := []string{"кот", "kot", "cat"}
	anagrams := searchAnagram(arr)
	expected := map[string][]string{}

	if !reflect.DeepEqual(anagrams, expected) {
		t.Errorf("The anagram is incorrect")
	}
}

func TestAnagram3(t *testing.T) {
	arr := []string{"пес", "сеп", "псе"}
	anagrams := searchAnagram(arr)
	expected := "пес"
	for k := range anagrams {
		if !reflect.DeepEqual(k, expected) {
			t.Errorf("The anagram is incorrect")
		}
	}
}
