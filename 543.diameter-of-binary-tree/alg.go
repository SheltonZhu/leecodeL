package _43_diameter_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var ans = 1
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		L := depth(node.Left)
		R := depth(node.Right)
		ans = max(ans, L+R+1)
		return max(L, R) + 1
	}
	depth(root)
	return ans - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
