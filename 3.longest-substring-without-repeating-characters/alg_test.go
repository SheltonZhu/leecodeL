package __longest_substring_without_repeating_characters

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	s := "abcabcbb"
	ret := lengthOfLongestSubstring(s)
	expected := 3
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}
func Test2(t *testing.T) {
	s := "bbbbb"
	ret := lengthOfLongestSubstring(s)
	expected := 1
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}
func Test3(t *testing.T) {
	s := "pwwkew"
	ret := lengthOfLongestSubstring(s)
	expected := 3
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}
