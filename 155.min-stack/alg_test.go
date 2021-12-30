package _55_min_stack

import "testing"

func Test(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	min := minStack.GetMin() // return -3

	if min != -3 {
		t.Errorf("min[%d] != %d", -3, min)
	}
	minStack.Pop()
	top := minStack.Top() // return 0
	if top != 0 {
		t.Errorf("top[%d] != %d", 0, top)
	}
	min = minStack.GetMin() // return -2
	if min != -2 {
		t.Errorf("min[%d] != %d", -3, min)
	}

}

func Test2(t *testing.T) {
	minStack := Constructor()
	minStack.Push(1)
	minStack.Push(2)
	minStack.Top()
	minStack.GetMin()
	minStack.Pop()
	minStack.GetMin()
	minStack.Top()
}
