package _02_binary_tree_level_order_traversal

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	output := [][]int{{3}, {9, 20}, {15, 7}}
	ret := levelOrder(root)
	if !reflect.DeepEqual(output, ret) {
		t.Errorf("failed. output: %v, return: %v", output, ret)
	}
}
