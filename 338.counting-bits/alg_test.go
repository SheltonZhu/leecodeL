package _38_counting_bits

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	n := 2
	output := []int{0, 1, 1}
	ret := countBits(n)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	n := 5
	output := []int{0, 1, 1, 2, 1, 2}
	ret := countBits(n)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
