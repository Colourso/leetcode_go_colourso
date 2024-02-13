package l0104_maximum_depth_of_binary_tree

/*
给定一个二叉树 root ，返回其最大深度。

二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数

https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 左、右、跟，聚合左右的节点，计算根节点的树
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	depth := 1
	if leftDepth >= rightDepth {
		depth += leftDepth
	} else {
		depth += rightDepth
	}
	return depth
}
