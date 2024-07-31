package main

import (
	"os/exec"
	"testing"
)

func TestCut1(t *testing.T) {
	args := &flags{
		f: []int{1},
		d: "\t",
		s: false,
	}

	// pReader, pWriter := io.Pipe()
	// pReader.Read()

	parseStrings(*args)
	cmd := exec.Command("echo", "qwe\tqwe")
	cmd.Run()
	// expectedOutput := "hello1qhello2\n"
	// if output.String() != expectedOutput {
	// 	t.Errorf("Expected %q but got %q", expectedOutput, output.String())
	// }

}
