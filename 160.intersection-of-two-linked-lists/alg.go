package _60_intersection_of_two_linked_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	nodeMap := map[*ListNode]bool{}
//	for tmp := headA; tmp != nil; tmp = tmp.Next {
//		nodeMap[tmp] = true
//	}
//	for tmp := headB; tmp != nil; tmp = tmp.Next {
//		if nodeMap[tmp] {
//			return tmp
//		}
//	}
//	return nil
//}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
