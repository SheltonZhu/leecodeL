package _1_merge_two_sorted_lists

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	output := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}}
	ret := mergeTwoLists(list1, list2)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	list1 := &ListNode{}
	list2 := &ListNode{}
	output := &ListNode{}
	ret := mergeTwoLists(list1, list2)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test3(t *testing.T) {
	list1 := &ListNode{}
	list2 := &ListNode{0, nil}
	output := &ListNode{0, nil}
	ret := mergeTwoLists(list1, list2)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
