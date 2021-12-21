package _2_generate_parentheses

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	n := 3
	output := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	ret := generateParenthesis(n)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	n := 1
	output := []string{"()"}
	ret := generateParenthesis(n)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
