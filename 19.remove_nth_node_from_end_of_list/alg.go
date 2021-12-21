package _9_remove_nth_node_from_end_of_list

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
func getLength(head *ListNode) int {
	l := 0
	tmp := head
	for tmp != nil {
		l++
		tmp = tmp.Next
	}
	return l
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodeLen := getLength(head)
	deleteP := nodeLen - n
	curr := 0
	tmp := head
	LastNodeP := deleteP - 1
	for curr < LastNodeP {
		curr++
		tmp = tmp.Next
	}
	if deleteP == 0 {
		return head.Next
	}
	tmp.Next = tmp.Next.Next
	return head
}
