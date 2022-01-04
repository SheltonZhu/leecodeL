package _60_intersection_of_two_linked_lists

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	output := &ListNode{8, nil}
	output.Next = &ListNode{4, &ListNode{5, nil}}
	listA := &ListNode{4, &ListNode{1, output}}
	listB := &ListNode{5, &ListNode{6, &ListNode{1, output}}}
	ret := getIntersectionNode(listA, listB)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	output := &ListNode{2, nil}
	output.Next = &ListNode{4, nil}
	listA := &ListNode{1, &ListNode{9, &ListNode{1, output}}}
	listB := &ListNode{3, output}
	ret := getIntersectionNode(listA, listB)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
