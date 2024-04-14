package zigzagconversion

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	s := "PAYPALISHIRING"
	numRows := 3
	ret := convert(s, numRows)
	expected := "PAHNAPLSIIGYIR"
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test2(t *testing.T) {
	s := "PAYPALISHIRING"
	numRows := 4
	ret := convert(s, numRows)
	expected := "PINALSIGYAHRPI"
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test3(t *testing.T) {
	s := "A"
	numRows := 1
	ret := convert(s, numRows)
	expected := "A"
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}
