package _22_coin_change

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	output := 3
	ret := coinChange(coins, amount)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test2(t *testing.T) {
	coins := []int{2}
	amount := 3
	output := -1
	ret := coinChange(coins, amount)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test3(t *testing.T) {
	coins := []int{1}
	amount := 0
	output := 0
	ret := coinChange(coins, amount)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test4(t *testing.T) {
	coins := []int{1}
	amount := 1
	output := 1
	ret := coinChange(coins, amount)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test5(t *testing.T) {
	coins := []int{1}
	amount := 2
	output := 2
	ret := coinChange(coins, amount)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
