package _39_daily_temperatures

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	output := []int{1, 1, 4, 2, 1, 1, 0, 0}
	ret := dailyTemperatures(temperatures)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	temperatures := []int{30, 40, 50, 60}
	output := []int{1, 1, 1, 0}
	ret := dailyTemperatures(temperatures)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test3(t *testing.T) {
	temperatures := []int{30, 60, 90}
	output := []int{1, 1, 0}
	ret := dailyTemperatures(temperatures)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
