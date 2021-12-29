package _41_linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	listNodeMap := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := listNodeMap[head]; ok {
			return head
		}
		listNodeMap[head] = struct{}{}
		head = head.Next
	}
	return nil

}
