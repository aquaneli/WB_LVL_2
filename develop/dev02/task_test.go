package main

import (
	"testing"
)

func TestUnpack1(t *testing.T) {
	input := "a4bc2d5e"
	expected := "aaaabccddddde"
	result, err := Unpack(input)
	if err != nil {
		t.Errorf("incorrect string")
	}
	if result != expected {
		t.Errorf("Expected %s, result %s", expected, result)
	}
}
