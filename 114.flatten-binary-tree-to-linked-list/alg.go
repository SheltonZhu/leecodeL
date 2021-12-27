package _05_construct_binary_tree_from_preorder_and_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var queue []*TreeNode
	var VLR func(root *TreeNode)
	VLR = func(root *TreeNode) {
		if root != nil {
			queue = append(queue, root)
			VLR(root.Left)
			VLR(root.Right)
		}
	}
	VLR(root)
	queue = append(queue, nil)
	for i := 0; i < len(queue)-1; i++ {
		queue[i].Left = nil
		queue[i].Right = queue[i+1]
	}
}
