package _21_task_scheduler

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2
	output := 8
	ret := leastInterval(tasks, n)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test2(t *testing.T) {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 0
	output := 6
	ret := leastInterval(tasks, n)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test3(t *testing.T) {
	tasks := []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	n := 2
	output := 16
	ret := leastInterval(tasks, n)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
