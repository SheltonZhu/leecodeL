package _21_best_time_to_buy_and_sell_stock

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	input := []int{7, 1, 5, 3, 6, 4}
	output := 5
	ret := maxProfit(input)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	input := []int{7, 6, 4, 3, 1}
	output := 0
	ret := maxProfit(input)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
