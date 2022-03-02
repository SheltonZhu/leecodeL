package _79_perfect_squares

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	n := 12
	output := 3
	ret := numSquares(n)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	n := 13
	output := 2
	ret := numSquares(n)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test3(t *testing.T) {
	n := 4
	output := 1
	ret := numSquares(n)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
