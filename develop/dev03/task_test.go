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
	res := sortStrings(data, args)

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

	if !reflect.DeepEqual(res, dataRes) {
		t.Errorf("Incorrect sort")
	}
}

func TestSort2(t *testing.T) {
	args := func() flags {
		kf := 0
		nf := true
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
	res := sortStrings(data, args)

	dataRes := []string{
		"mouse 200 qwe qwe qwe qwe",
		"February 1 computer 150",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"March 1 computer 150",
		"January 1 computer 150",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"camputer January 450 b ",
		"laptop 30 a",
		"a.1.1qwe",
		" mouse 200 qwe qwe qwe qwe",
		"0.1,1Gaas 1 q 23 -1",
		"0.123G bdebian 100",
		"1M laptop January 90 zcx",
		"3.001Glaptop 90 zcx",
		"12345M computer 300 c",
		"10000000000000000000000000000000000000000000000000000000000000000000M computer 300 c",
	}

	if !reflect.DeepEqual(res, dataRes) {
		t.Errorf("Incorrect sort")
	}
}

func TestSort3(t *testing.T) {
	args := func() flags {
		kf := 0
		nf := false
		rf := true
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
	res := sortStrings(data, args)

	dataRes := []string{
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"laptop 30 a",
		"camputer January 450 b ",
		"a.1.1qwe",
		"March 1 computer 150",
		"January 1 computer 150",
		"February 1 computer 150",
		"3.001Glaptop 90 zcx",
		"1M laptop January 90 zcx",
		"12345M computer 300 c",
		"10000000000000000000000000000000000000000000000000000000000000000000M computer 300 c",
		"0.123G bdebian 100",
		"0.1,1Gaas 1 q 23 -1",
		" mouse 200 qwe qwe qwe qwe",
	}

	if !reflect.DeepEqual(res, dataRes) {
		t.Errorf("Incorrect sort")
	}
}

func TestSort4(t *testing.T) {
	args := func() flags {
		kf := 0
		nf := false
		rf := false
		uf := true
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
	res := sortStrings(data, args)

	dataRes := []string{
		"0.1,1Gaas 1 q 23 -1",
		"0.123G bdebian 100",
		"February 1 computer 150",
		"a.1.1qwe",
		"March 1 computer 150",
		"January 1 computer 150",
		"mouse 200 qwe qwe qwe qwe",
		" mouse 200 qwe qwe qwe qwe",
		"1M laptop January 90 zcx",
		"3.001Glaptop 90 zcx",
		"laptop 30 a",
		"camputer January 450 b ",
		"12345M computer 300 c",
		"10000000000000000000000000000000000000000000000000000000000000000000M computer 300 c",
	}

	if !reflect.DeepEqual(res, dataRes) {
		t.Errorf("Incorrect sort")
	}
}

func TestSort5(t *testing.T) {
	args := func() flags {
		kf := 0
		nf := false
		rf := false
		uf := false
		Mf := true
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
	res := sortStrings(data, args)

	dataRes := []string{
		"1M laptop January 90 zcx",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"0.123G bdebian 100",
		"mouse 200 qwe qwe qwe qwe",
		"camputer January 450 b ",
		"10000000000000000000000000000000000000000000000000000000000000000000M computer 300 c",
		"12345M computer 300 c",
		"a.1.1qwe",
		" mouse 200 qwe qwe qwe qwe",
		"0.1,1Gaas 1 q 23 -1",
		"3.001Glaptop 90 zcx",
		"laptop 30 a",
		"January 1 computer 150",
		"February 1 computer 150",
		"March 1 computer 150",
	}

	if !reflect.DeepEqual(res, dataRes) {
		t.Errorf("Incorrect sort")
	}
}

func TestSort6(t *testing.T) {
	args := func() flags {
		kf := 0
		nf := false
		rf := false
		uf := false
		Mf := false
		bf := true
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
	res := sortStrings(data, args)

	dataRes := []string{
		"0.1,1Gaas 1 q 23 -1",
		"0.123G bdebian 100",
		"10000000000000000000000000000000000000000000000000000000000000000000M computer 300 c",
		"12345M computer 300 c",
		"1M laptop January 90 zcx",
		"3.001Glaptop 90 zcx",
		"February 1 computer 150",
		"January 1 computer 150",
		"March 1 computer 150",
		"a.1.1qwe",
		"camputer January 450 b ",
		"laptop 30 a",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		" mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
	}

	if !reflect.DeepEqual(res, dataRes) {
		t.Errorf("Incorrect sort")
	}
}

func TestSort7(t *testing.T) {
	args := func() flags {
		kf := 0
		nf := false
		rf := false
		uf := false
		Mf := false
		bf := false
		cf := true
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
	res := sortStrings(data, args)
	if res != nil {
		t.Errorf("The file is not sorted")
	}
}

func TestSort8(t *testing.T) {
	args := func() flags {
		kf := 0
		nf := false
		rf := false
		uf := false
		Mf := false
		bf := false
		cf := false
		hf := true
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
	res := sortStrings(data, args)

	dataRes := []string{
		"mouse 200 qwe qwe qwe qwe",
		"February 1 computer 150",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"March 1 computer 150",
		"January 1 computer 150",
		"mouse 200 qwe qwe qwe qwe",
		"mouse 200 qwe qwe qwe qwe",
		"camputer January 450 b ",
		"laptop 30 a",
		"a.1.1qwe",
		" mouse 200 qwe qwe qwe qwe",
		"0.1,1Gaas 1 q 23 -1",
		"1M laptop January 90 zcx",
		"12345M computer 300 c",
		"10000000000000000000000000000000000000000000000000000000000000000000M computer 300 c",
		"0.123G bdebian 100",
		"3.001Glaptop 90 zcx",
	}

	if !reflect.DeepEqual(res, dataRes) {
		t.Errorf("Incorrect sort")
	}

}
