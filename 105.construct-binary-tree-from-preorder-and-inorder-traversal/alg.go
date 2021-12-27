package _05_construct_binary_tree_from_preorder_and_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{}
	root.Val = preorder[0]

	position := 0
	for ; position < len(inorder); position++ {
		if inorder[position] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:position])+1], inorder[:position])
	root.Right = buildTree(preorder[len(inorder[:position])+1:], inorder[position+1:])
	return root
}
