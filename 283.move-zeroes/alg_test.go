package _83_move_zeroes

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	input := []int{0, 1, 0, 3, 12}
	output := []int{1, 3, 12, 0, 0}
	moveZeroes(input)
	if !reflect.DeepEqual(output, input) {
		t.Errorf("failed. output: %v, return: %v", output, input)
	}
}
