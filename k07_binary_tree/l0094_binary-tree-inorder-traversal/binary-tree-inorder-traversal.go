package l0094_binary_tree_inorder_traversal

/*
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

https://leetcode.cn/problems/binary-tree-inorder-traversal/description/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	return result
}

// inorderTraversalV2 二叉树的中序遍历-迭代法
func inorderTraversalV2(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}

	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, cur.Val)
			cur = cur.Right
		}
	}
	return result
}
