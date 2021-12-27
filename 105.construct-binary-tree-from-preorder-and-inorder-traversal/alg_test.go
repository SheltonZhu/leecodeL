package _05_construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	output := &TreeNode{
		3,
		&TreeNode{9, nil, nil},
		&TreeNode{20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil},
		},
	}
	ret := buildTree(preorder, inorder)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test2(t *testing.T) {
	preorder := []int{-1}
	inorder := []int{-1}
	output := &TreeNode{}
	ret := buildTree(preorder, inorder)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
