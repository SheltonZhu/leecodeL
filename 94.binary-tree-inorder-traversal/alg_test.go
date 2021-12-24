package _4_binary_tree_inorder_traversal

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	output := []int{1, 3, 2}
	ret := inorderTraversal(root)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test2(t *testing.T) {
	root := &TreeNode{}
	var output []int
	ret := inorderTraversal(root)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test3(t *testing.T) {
	root := &TreeNode{1, nil, nil}
	output := []int{1}
	ret := inorderTraversal(root)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
func Test4(t *testing.T) {
	root := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	output := []int{2, 1}
	ret := inorderTraversal(root)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}

func Test5(t *testing.T) {
	root := &TreeNode{1, nil, &TreeNode{2, nil, nil}}
	output := []int{1, 2}
	ret := inorderTraversal(root)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
