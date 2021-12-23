package _5_jump_game

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	output := true
	ret := canJump(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	nums := []int{3, 2, 1, 0, 4}
	output := false
	ret := canJump(nums)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
