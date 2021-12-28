package _41_linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	p, q := head, head.Next
	for q != p {
		if q == nil || q.Next == nil {
			return false
		}
		p = p.Next
		q = q.Next.Next
	}
	return true
}
