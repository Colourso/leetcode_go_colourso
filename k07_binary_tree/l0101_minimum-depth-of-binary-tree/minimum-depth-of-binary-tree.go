package l0101_minimum_depth_of_binary_tree

/*
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明：叶子节点是指没有子节点的节点。

https://leetcode.cn/problems/minimum-depth-of-binary-tree/description/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)

	if root.Left != nil && root.Right == nil {
		return 1 + leftDepth
	}
	if root.Left == nil && root.Right != nil {
		return 1 + rightDepth
	}

	depth := 1
	if leftDepth <= rightDepth {
		depth += leftDepth
	} else {
		depth += rightDepth
	}
	return depth
}
