package _9_group_anagrams

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	output := [][]string{
		{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"},
	}
	ret := groupAnagrams(strs)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test1(t *testing.T) {
	strs := []string{""}
	output := [][]string{
		{""},
	}
	ret := groupAnagrams(strs)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	strs := []string{"a"}
	output := [][]string{
		{"a"},
	}
	ret := groupAnagrams(strs)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
