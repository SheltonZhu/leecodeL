package _0_valid_parentheses

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	s := "()"
	output := true
	ret := isValid(s)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test2(t *testing.T) {
	s := "()[]{}"
	output := true
	ret := isValid(s)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test3(t *testing.T) {
	s := "(]"
	output := false
	ret := isValid(s)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test4(t *testing.T) {
	s := "([)]"
	output := false
	ret := isValid(s)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test5(t *testing.T) {
	s := "{[]}"
	output := true
	ret := isValid(s)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
