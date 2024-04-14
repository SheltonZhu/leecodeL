package palindromenumber

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	s := 121
	expected := true
	ret := isPalindrome(s)
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test2(t *testing.T) {
	s := -121
	expected := false
	ret := isPalindrome(s)
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test3(t *testing.T) {
	s := 10
	expected := false
	ret := isPalindrome(s)
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}
