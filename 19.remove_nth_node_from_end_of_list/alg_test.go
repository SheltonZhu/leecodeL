package _9_remove_nth_node_from_end_of_list

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	n := 2
	output := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, nil}}}}
	ret := removeNthFromEnd(head, n)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test2(t *testing.T) {
	head := &ListNode{1, nil}
	n := 1
	output := &ListNode{}
	ret := removeNthFromEnd(head, n)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test3(t *testing.T) {
	head := &ListNode{1, &ListNode{2, nil}}
	n := 1
	output := &ListNode{1, nil}
	ret := removeNthFromEnd(head, n)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
