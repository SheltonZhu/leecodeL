package _39_word_break

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	output := true
	ret := wordBreak(s, wordDict)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	s := "applepenapple"
	wordDict := []string{"apple", "pen"}
	output := true
	ret := wordBreak(s, wordDict)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test3(t *testing.T) {
	s := "catsandog"
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	output := false
	ret := wordBreak(s, wordDict)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
