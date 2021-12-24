package _4_binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	var LDR func(node *TreeNode)
	LDR = func(node *TreeNode) {
		if node == nil {
			return
		}
		LDR(node.Left)
		result = append(result, node.Val)
		LDR(node.Right)
	}
	LDR(root)
	return result
}
