package _34_palindrome_linked_list

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	head := &ListNode{1,
		&ListNode{2,
			&ListNode{2,
				&ListNode{1, nil},
			}}}
	output := true
	ret := isPalindrome(head)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	head := &ListNode{1,
		&ListNode{2, nil}}
	output := false
	ret := isPalindrome(head)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
