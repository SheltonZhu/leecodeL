package _1_merge_two_sorted_lists

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	tmp := l
	str := strings.Builder{}
	for tmp.Next != nil {
		str.WriteString(strconv.Itoa(tmp.Val))
		str.WriteString("->")
		tmp = tmp.Next
	}
	str.WriteString(strconv.Itoa(tmp.Val))
	return str.String()
}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	mergedList := &ListNode{}
	head := mergedList
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}
	if list1 != nil {
		head.Next = list1
	}
	if list2 != nil {
		head.Next = list2
	}
	return mergedList.Next
}
