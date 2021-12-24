package _0_climbing_stairs

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	input := 2
	output := 2
	ret := climbStairs(input)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	input := 3
	output := 3
	ret := climbStairs(input)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
