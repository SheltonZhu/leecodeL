package _41_linked_list_cycle

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	posNode := &ListNode{2, nil}
	head := &ListNode{3, posNode}
	midNode := &ListNode{2, &ListNode{0, &ListNode{-4, &ListNode{3, posNode}}}}
	posNode.Next = midNode
	//pos := 1
	output := posNode
	ret := detectCycle(head)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	last := &ListNode{2, nil}
	head := &ListNode{1, last}
	last.Next = head
	//pos := 0
	output := head
	ret := detectCycle(head)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test3(t *testing.T) {
	head := &ListNode{1, nil}
	//pos := -1
	output := &ListNode{}
	ret := detectCycle(head)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
