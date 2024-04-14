package regularexpressionmatching

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	s := "aa"
	p := "a"
	expected := false
	ret := isMatch(s, p)
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test2(t *testing.T) {
	s := "aa"
	p := "a*"
	expected := true
	ret := isMatch(s, p)
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test3(t *testing.T) {
	s := "aB"
	p := ".*"
	expected := true
	ret := isMatch(s, p)
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}
