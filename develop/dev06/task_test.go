package main

import (
	"log"
	"os"
	"os/exec"
	"reflect"
	"testing"
)

func TestCut1(t *testing.T) {
	args := &flags{
		f: []int{0, 1},
		d: "\t",
		s: false,
	}

	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	originalStdin := os.Stdin
	defer func() {
		os.Stdin = originalStdin
	}()
	os.Stdin = r

	//r значит считываем данные и эти считанные данные мы закидываем в os.Stdin

	example := []string{"qwe\tqwe\n", "zxc\tzxc"}

	go func() {
		defer w.Close()
		for _, v := range example {
			w.WriteString(v)
		}
	}()

	res := getResult(*args)

	expected := []string{"qwe\tqwe", "zxc\tzxc"}

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %q but got %q", expected, res)
	}

}

func TestCut2(t *testing.T) {
	args := &flags{
		f: []int{0, 1},
		d: "0",
		s: false,
	}

	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}

	originalStdin := os.Stdin
	defer func() {
		os.Stdin = originalStdin
	}()
	os.Stdin = r

	go func() {
		defer w.Close()
		cmd := exec.Command("cat", "test1.txt")
		cmd.Stdout = w
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	results := getResult(*args)

	expected := []string{
		"mouse 20",
		"mouse 20",
		"3.0",
		"Q",
		"0.1",
		"0.1",
		"0.1",
		"mouse 20",
		"mouse 20",
		"mouse 20",
		"0.1",
		"mouse 20",
		" mouse 20",
		"1M laptop January 90 zcx",
		"3.0",
		"laptop 30 a",
		"camputer January 450 b ",
		"12345M computer 30",
		"10",
		"0.1",
	}

	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expected %q but got %q", expected, results)
	}
}

func TestCut3(t *testing.T) {
	args := &flags{
		f: []int{0},
		d: " ",
		s: true,
	}

	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}

	originalStdin := os.Stdin
	defer func() {
		os.Stdin = originalStdin
	}()
	os.Stdin = r

	go func() {
		defer w.Close()
		cmd := exec.Command("cat", "test1.txt")
		cmd.Stdout = w
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	results := getResult(*args)

	expected := []string{
		"mouse",
		"mouse",
		"3.00.1Glaptop",
		"mouse",
		"mouse",
		"mouse",
		"mouse",
		"",
		"1M",
		"3.00.1Glaptop",
		"laptop",
		"camputer",
		"12345M",
		"10000000000000000000000000000000000000000000000000000000000000000000M",
	}

	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expected %q but got %q", expected, results)
	}
}
