package l0144_binary_tree_preorder_traversal

/*
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

https://leetcode.cn/problems/binary-tree-preorder-traversal/description/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	preOrder(root, &result)
	return result
}

func preOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)
	preOrder(root.Left, result)
	preOrder(root.Right, result)
}

// preorderTraversalV2 二叉树前序遍历迭代法
func preorderTraversalV2(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	// 使用栈实现
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result
}
