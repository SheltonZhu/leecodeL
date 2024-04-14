package stringtointeger

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	s := "42"
	expected := 42
	ret := myAtoi(s)
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test2(t *testing.T) {
	s := "   -42"
	expected := -42
	ret := myAtoi(s)
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test3(t *testing.T) {
	s := "4193 with words"
	expected := 4193
	ret := myAtoi(s)
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}
