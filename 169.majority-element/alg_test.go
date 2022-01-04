package _69_majority_element

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	input := []int{3, 2, 3}
	output := 3
	ret := majorityElement(input)
	if !reflect.DeepEqual(ret, output) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	input := []int{2, 2, 1, 1, 1, 2, 2}
	output := 2
	ret := majorityElement(input)
	if !reflect.DeepEqual(ret, output) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
