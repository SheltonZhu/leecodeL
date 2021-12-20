package _7_letter_combinations_of_a_phone_number

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	digits := "23"
	output := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	ret := letterCombinations(digits)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	digits := ""
	var output []string
	ret := letterCombinations(digits)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test3(t *testing.T) {
	digits := "2"
	output := []string{"a", "b", "c"}
	ret := letterCombinations(digits)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
