package main

import (
	"testing"
)

func TestUnpack1(t *testing.T) {
	input := "a4bc2d5e"
	expected := "aaaabccddddde"
	result, err := Unpack(input)
	if err != nil {
		t.Errorf("Incorrect error handling")
	}
	if result != expected {
		t.Errorf("Expected %s, result %s", expected, result)
	}
}

func TestUnpack2(t *testing.T) {
	input := "abcd"
	expected := "abcd"
	result, err := Unpack(input)
	if err != nil {
		t.Errorf("Incorrect error handling")
	}
	if result != expected {
		t.Errorf("Expected %s, result %s", expected, result)
	}
}

func TestUnpack3(t *testing.T) {
	input := "45"
	expected := ""
	result, err := Unpack(input)
	if err == nil {
		t.Errorf("Incorrect error handling")
	}
	if result != expected {
		t.Errorf("Expected %s, result %s", expected, result)
	}
}

func TestUnpack4(t *testing.T) {
	input := ""
	expected := ""
	result, err := Unpack(input)
	if err != nil {
		t.Errorf("Incorrect error handling")
	}
	if result != expected {
		t.Errorf("Expected %s, result %s", expected, result)
	}
}

func TestUnpack5(t *testing.T) {
	input := "qwe\\4\\5"
	expected := "qwe45"
	result, err := Unpack(input)
	if err != nil {
		t.Errorf("Incorrect error handling")
	}
	if result != expected {
		t.Errorf("Expected %s, result %s", expected, result)
	}
}

func TestUnpack6(t *testing.T) {
	input := "qwe\\45"
	expected := "qwe44444"
	result, err := Unpack(input)
	if err != nil {
		t.Errorf("Incorrect error handling")
	}
	if result != expected {
		t.Errorf("Expected %s, result %s", expected, result)
	}
}

func TestUnpack7(t *testing.T) {
	input := "qwe\\\\5"
	expected := "qwe\\\\\\\\\\"
	result, err := Unpack(input)
	if err != nil {
		t.Errorf("Incorrect error handling")
	}
	if result != expected {
		t.Errorf("Expected %s, result %s", expected, result)
	}
}

func TestUnpack8(t *testing.T) {
	input := "3\\\\0a4bc\\\\\\\\\\2d5e\\35"
	expected := "aaaabc\\\\2ddddde33333"
	result, err := Unpack(input)
	if err != nil {
		t.Errorf("Incorrect error handling")
	}
	if result != expected {
		t.Errorf("Expected %s, result %s", expected, result)
	}
}

func TestUnpack9(t *testing.T) {
	input := "\\\\03"
	expected := ""
	result, err := Unpack(input)
	if err == nil {
		t.Errorf("Incorrect error handling")
	}
	if result != expected {
		t.Errorf("Expected %s, result %s", expected, result)
	}
}

func TestUnpack10(t *testing.T) {
	input := "\\\\0\\3q0q0q5\\\\\\\\\\\\\\8\\"
	expected := "3qqqqq\\\\\\8"
	result, err := Unpack(input)
	if err != nil {
		t.Errorf("Incorrect error handling")
	}
	if result != expected {
		t.Errorf("Expected %s, result %s", expected, result)
	}
}
