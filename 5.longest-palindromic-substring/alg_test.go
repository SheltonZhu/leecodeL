package __longest_palindromic_substring

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	s := "babad"
	ret := longestPalindrome(s)
	expected := "bab"
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test2(t *testing.T) {
	s := "cbbd"
	ret := longestPalindrome(s)
	expected := "bb"
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test3(t *testing.T) {
	s := "a"
	ret := longestPalindrome(s)
	expected := "a"
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test4(t *testing.T) {
	s := "ac"
	ret := longestPalindrome(s)
	expected := "a"
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test5(t *testing.T) {
	s := "aacabdkacaa"
	ret := longestPalindrome(s)
	expected := "aca"
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test6(t *testing.T) {
	s := "eabcb"
	ret := longestPalindrome(s)
	expected := "bcb"
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test7(t *testing.T) {
	s := ""
	ret := longestPalindrome(s)
	expected := ""
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}
