package _17_merge_two_binary_trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var merge func(node1, node2 *TreeNode) *TreeNode
	merge = func(node1, node2 *TreeNode) *TreeNode {
		if node1 == nil && node2 == nil {
			return nil
		}
		node1Value, node2Value := 0, 0
		var node1Left, node2Left, node1Right, node2Right *TreeNode
		mergedNode := &TreeNode{}
		if node1 != nil {
			node1Value = node1.Val
			node1Left = node1.Left
			node1Right = node1.Right
		}
		if node2 != nil {
			node2Value = node2.Val
			node2Left = node2.Left
			node2Right = node2.Right
		}
		mergedNode.Val = node1Value + node2Value
		mergedNode.Left = merge(node1Left, node2Left)
		mergedNode.Right = merge(node1Right, node2Right)
		return mergedNode
	}
	return merge(root1, root2)
}
