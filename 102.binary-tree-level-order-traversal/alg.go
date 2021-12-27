package _02_binary_tree_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return ret
	}
	queue := []*TreeNode{root}
	var scanLevel func(q []*TreeNode)
	scanLevel = func(q []*TreeNode) {
		if len(q) <= 0 {
			return
		}
		var level []int
		var newLevelQ []*TreeNode
		for i := 0; i < len(q); i++ {
			level = append(level, q[i].Val)
			if q[i].Left != nil {
				newLevelQ = append(newLevelQ, q[i].Left)
			}
			if q[i].Right != nil {
				newLevelQ = append(newLevelQ, q[i].Right)
			}
		}
		ret = append(ret, level)
		scanLevel(newLevelQ)
	}
	scanLevel(queue)
	return ret
}
