package _34_palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	half := endOfFirstHalf(head)
	reverseHalf := reverseList(half.Next)

	l := head
	r := reverseHalf
	result := true
	for r != nil && result {
		if l.Val != r.Val {
			result = false
		}
		l = l.Next
		r = r.Next
	}
	half.Next = reverseList(reverseHalf)
	return result
}

func reverseList(head *ListNode) *ListNode {
	var prev, curr *ListNode = nil, head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func endOfFirstHalf(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
