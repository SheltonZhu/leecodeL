package _06_reverse_linked_list

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	output := &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}
	ret := reverseList(head)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	head := &ListNode{1, &ListNode{2, nil}}
	output := &ListNode{2, &ListNode{1, nil}}
	ret := reverseList(head)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test3(t *testing.T) {
	head := &ListNode{}
	output := &ListNode{}
	ret := reverseList(head)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
