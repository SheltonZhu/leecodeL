package _61_hamming_distance

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	x := 1
	y := 4
	output := 2
	ret := hammingDistance(x, y)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
	//fmt.Printf("%b ^ %b = %b\n", 1, 4, 1^4)
}

func Test2(t *testing.T) {
	x := 3
	y := 1
	output := 1
	ret := hammingDistance(x, y)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
