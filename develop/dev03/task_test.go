package main

import (
	"reflect"
	"testing"
)

func TestSort1(t *testing.T) {
	args := func() flags {
		kf := 3
		nf := false
		rf := false
		uf := false
		Mf := false
		bf := false
		cf := false
		hf := false
		return flags{
			k:    &kf,
			n:    &nf,
			r:    &rf,
			u:    &uf,
			M:    &Mf,
			b:    &bf,
			c:    &cf,
			h:    &hf,
			path: []string{"test.txt"},
		}
	}()

	data, err := parseFiles(&args)
	if err != nil {
		t.Errorf("Incorrect parsing")
	}
	sortStrings(&data, args)

	dataRes := []string{
		"a.1.1qwe",
		"0.123G bdebian 100",
		"10000000000000000000000000000000000000000000000000000000000000000000M computer 300 c",
		"12345M computer 300 c",
		"camputer January 450 b ",
		"1M laptop January 90 zcx",
		"laptop 30 a",
		"March 1 computer 150",
		"January 1 computer 150",
		"February 1 computer 150",
		"0.1,1Gaas 1 q 23 -1",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		" mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"3.001Glaptop 90 zcx",
	}

	if !reflect.DeepEqual(data, dataRes) {
		t.Errorf("Incorrect sort")
	}
}
