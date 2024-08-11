package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"testing"
)

func TestCdPwd(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	rInScan, wInScan, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer rInScan.Close()

	go func() {
		defer wInScan.Close()
		cmd := exec.Command("go", "run", "task.go")
		cmd.Stdin = r
		cmd.Stdout = wInScan
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer w.Close()
		w.WriteString("cd ../dev07/\n")
		w.WriteString("pwd\n")
	}()

	expected := "/Users/aquaneli/WB_LVL_2/develop/dev07"
	scanner := bufio.NewScanner(rInScan)
	scanner.Scan()
	fileds := strings.Fields(scanner.Text())
	res := strings.Join(fileds[len(fileds)-1:], " ")

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %q but got %q", expected, res)
	}

}

func TestEcho(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	rInScan, wInScan, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer rInScan.Close()

	go func() {
		defer wInScan.Close()
		cmd := exec.Command("go", "run", "task.go")
		cmd.Stdin = r
		cmd.Stdout = wInScan
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		defer w.Close()
		w.WriteString("echo TestEcho work\n")
	}()

	expected := "TestEcho work"
	scanner := bufio.NewScanner(rInScan)
	scanner.Scan()
	fileds := strings.Fields(scanner.Text())
	res := strings.Join(fileds[len(fileds)-2:], " ")

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %q but got %q", expected, res)
	}

}
