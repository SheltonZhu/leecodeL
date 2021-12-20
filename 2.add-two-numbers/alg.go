package __add_two_numbers

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := ListNode{}
	carry := 0
	sumNode := &ret
	for l1 != nil || l2 != nil || carry > 0 {
		x := 0
		y := 0
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}
		sum := x + y + carry
		carry = sum / 10
		sumNode.Next = &ListNode{Val: sum % 10}
		sumNode = sumNode.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return ret.Next
}
