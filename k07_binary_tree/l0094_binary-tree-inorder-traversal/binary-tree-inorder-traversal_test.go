package l0094_binary_tree_inorder_traversal

import (
	"fmt"
	"testing"
)

func TestInOrder(t *testing.T) {
	node1 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	node2 := &TreeNode{
		Val:   2,
		Left:  node1,
		Right: nil,
	}
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: node2,
	}

	result := inorderTraversalV2(root)
	fmt.Println(result)
}
