package main

import (
	"reflect"
	"testing"
)

func TestGrep1(t *testing.T) {

	args := flags{
		A:        3,
		B:        0,
		C:        0,
		c:        false,
		i:        false,
		v:        false,
		F:        false,
		n:        false,
		pattern:  "q",
		pathFile: []string{"test1.txt"},
	}

	buff, indexBuff, buffStr := parseFile(args)
	resBuff := process(buff, indexBuff, args)

	expected := [][]string{{
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"3.00.1Glaptop 90 zcx",
		"Q",
		"0.1",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"0.1",
		"mouse 200 qwe qwe qwe qwe",
		" mouse 200 qwe qwe qwe qwe",
		"1M laptop January 90 zcx",
		"3.00.1Glaptop 90 zcx",
		"laptop 30 a",
	},
	}

	if len(expected) != len(resBuff) {
		t.Errorf("Processed different number of files")
	}

	for i := range resBuff {
		if len(expected[i]) != len(resBuff[i]) {
			t.Errorf("The number of lines is different")
		}
		for j, val := range resBuff[i] {
			if !reflect.DeepEqual(buffStr[i][val], expected[i][j]) {
				t.Errorf("The lines are different")
			}
		}
	}

}

func TestGrep2(t *testing.T) {

	args := flags{
		A:        0,
		B:        3,
		C:        0,
		c:        false,
		i:        false,
		v:        false,
		F:        false,
		n:        false,
		pattern:  "0.1",
		pathFile: []string{"test1.txt"},
	}

	buff, indexBuff, buffStr := parseFile(args)
	resBuff := process(buff, indexBuff, args)

	expected := [][]string{{
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"3.00.1Glaptop 90 zcx",
		"Q",
		"0.1",
		"0.1",
		"0.1",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"0.1",
		"mouse 200 qwe qwe qwe qwe",
		" mouse 200 qwe qwe qwe qwe",
		"1M laptop January 90 zcx",
		"3.00.1Glaptop 90 zcx",
		"camputer January 450 b ",
		"12345M computer 300 c",
		"10000000000000000000000000000000000000000000000000000000000000000000M computer 300 c",
		"0.1",
	},
	}

	if len(expected) != len(resBuff) {
		t.Errorf("Processed different number of files")
	}

	for i := range resBuff {
		if len(expected[i]) != len(resBuff[i]) {
			t.Errorf("The number of lines is different")
		}
		for j, val := range resBuff[i] {
			if !reflect.DeepEqual(buffStr[i][val], expected[i][j]) {
				t.Errorf("The lines are different")
			}
		}
	}

}

func TestGrep3(t *testing.T) {

	args := flags{
		A:        0,
		B:        0,
		C:        3,
		c:        false,
		i:        false,
		v:        false,
		F:        false,
		n:        false,
		pattern:  "x",
		pathFile: []string{"test1.txt"},
	}

	buff, indexBuff, buffStr := parseFile(args)
	resBuff := process(buff, indexBuff, args)

	expected := [][]string{{
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"3.00.1Glaptop 90 zcx",
		"Q",
		"0.1",
		"0.1",
		"0.1",
		"mouse 200 qwe qwe qwe qwe",
		" mouse 200 qwe qwe qwe qwe",
		"1M laptop January 90 zcx",
		"3.00.1Glaptop 90 zcx",
		"laptop 30 a",
		"camputer January 450 b ",
		"12345M computer 300 c",
	},
	}

	if len(expected) != len(resBuff) {
		t.Errorf("Processed different number of files")
	}

	for i := range resBuff {
		if len(expected[i]) != len(resBuff[i]) {
			t.Errorf("The number of lines is different")
		}
		for j, val := range resBuff[i] {
			if !reflect.DeepEqual(buffStr[i][val], expected[i][j]) {
				t.Errorf("The lines are different")
			}
		}
	}

}

func TestGrep4(t *testing.T) {

	args := flags{
		A:        0,
		B:        0,
		C:        0,
		c:        true,
		i:        false,
		v:        false,
		F:        false,
		n:        false,
		pattern:  "s",
		pathFile: []string{"test1.txt"},
	}

	buff, indexBuff, _ := parseFile(args)
	resBuff := process(buff, indexBuff, args)

	expected := 7

	for i := range resBuff {
		if expected != len(resBuff[i]) {
			t.Errorf("The number of lines is different")
		}
	}

}

func TestGrep5(t *testing.T) {

	args := flags{
		A:        0,
		B:        0,
		C:        0,
		c:        false,
		i:        true,
		v:        false,
		F:        false,
		n:        false,
		pattern:  "Q",
		pathFile: []string{"test1.txt"},
	}

	buff, indexBuff, buffStr := parseFile(args)
	resBuff := process(buff, indexBuff, args)

	expected := [][]string{{
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"Q",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		" mouse 200 qwe qwe qwe qwe",
	},
	}

	if len(expected) != len(resBuff) {
		t.Errorf("Processed different number of files")
	}

	for i := range resBuff {
		if len(expected[i]) != len(resBuff[i]) {
			t.Errorf("The number of lines is different")
		}
		for j, val := range resBuff[i] {
			if !reflect.DeepEqual(buffStr[i][val], expected[i][j]) {
				t.Errorf("The lines are different")
			}
		}
	}

}

func TestGrep6(t *testing.T) {

	args := flags{
		A:        0,
		B:        0,
		C:        0,
		c:        false,
		i:        false,
		v:        true,
		F:        false,
		n:        false,
		pattern:  "e",
		pathFile: []string{"test1.txt"},
	}

	buff, indexBuff, buffStr := parseFile(args)
	resBuff := process(buff, indexBuff, args)

	expected := [][]string{{
		"3.00.1Glaptop 90 zcx",
		"Q",
		"0.1",
		"0.1",
		"0.1",
		"0.1",
		"1M laptop January 90 zcx",
		"3.00.1Glaptop 90 zcx",
		"laptop 30 a",
		"0.1",
	},
	}

	if len(expected) != len(resBuff) {
		t.Errorf("Processed different number of files")
	}

	for i := range resBuff {
		if len(expected[i]) != len(resBuff[i]) {
			t.Errorf("The number of lines is different")
		}
		for j, val := range resBuff[i] {
			if !reflect.DeepEqual(buffStr[i][val], expected[i][j]) {
				t.Errorf("The lines are different")
			}
		}
	}

}

func TestGrep7(t *testing.T) {

	args := flags{
		A:        0,
		B:        0,
		C:        0,
		c:        false,
		i:        false,
		v:        false,
		F:        true,
		n:        false,
		pattern:  "[0-9]",
		pathFile: []string{"test1.txt"},
	}

	buff, indexBuff, buffStr := parseFile(args)
	resBuff := process(buff, indexBuff, args)

	expected := make([][]string, 1)

	if len(expected) != len(resBuff) {
		t.Errorf("Processed different number of files")
	}

	for i := range resBuff {
		if len(expected[i]) != len(resBuff[i]) {
			t.Errorf("The number of lines is different")
		}
		for j, val := range resBuff[i] {
			if !reflect.DeepEqual(buffStr[i][val], expected[i][j]) {
				t.Errorf("The lines are different")
			}
		}
	}

}

func TestGrep8(t *testing.T) {

	args := flags{
		A:        0,
		B:        0,
		C:        0,
		c:        false,
		i:        false,
		v:        false,
		F:        false,
		n:        true,
		pattern:  "o",
		pathFile: []string{"test1.txt"},
	}

	buff, indexBuff, _ := parseFile(args)
	resBuff := process(buff, indexBuff, args)

	expected := [][]int{
		{0, 1, 2, 7, 8, 9, 11, 12, 13, 14, 15, 17, 18},
	}

	if len(expected) != len(resBuff) {
		t.Errorf("Processed different number of files")
	}

	for i := range resBuff {
		if len(expected[i]) != len(resBuff[i]) {
			t.Errorf("The number of lines is different")
		}
		for j, val := range resBuff[i] {
			if !reflect.DeepEqual(val, expected[i][j]) {
				t.Errorf("The lines are different")
			}
		}
	}

}
