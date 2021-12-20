package __add_two_numbers

import (
	"fmt"
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	ret := addTwoNumbers(l1, l2)
	expected := &ListNode{7, &ListNode{0, &ListNode{8, nil}}}
	fmt.Println(ret)
	fmt.Println(expected)
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test2(t *testing.T) {
	l1 := &ListNode{0, nil}
	l2 := &ListNode{0, nil}
	ret := addTwoNumbers(l1, l2)
	expected := &ListNode{0, nil}
	fmt.Println(ret)
	fmt.Println(expected)
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test3(t *testing.T) {
	l1 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	l2 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	ret := addTwoNumbers(l1, l2)
	expected := &ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}}

	fmt.Println(ret)
	fmt.Println(expected)
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}

func Test4(t *testing.T) {
	l1 := &ListNode{9, &ListNode{9, &ListNode{1, nil}}}
	l2 := &ListNode{1, nil}
	ret := addTwoNumbers(l1, l2)
	expected := &ListNode{0, &ListNode{0, &ListNode{2, nil}}}
	fmt.Println(ret)
	fmt.Println(expected)
	if !reflect.DeepEqual(expected, ret) {
		t.Error("test failed.", expected, ret)
	}
}
