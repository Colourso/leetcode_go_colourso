package l0101_symmetric_tree

/*
给你一个二叉树的根节点 root ， 检查它是否轴对称。

https://leetcode.cn/problems/symmetric-tree/description/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}

func compare(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left != nil && right == nil {
		return false
	} else if left == nil && right != nil {
		return false
	} else if left.Val != right.Val {
		return false
	}

	// 根节点比较左右孩子的结果，汇总返回
	return compare(left.Left, right.Right) && compare(left.Right, right.Left)
}
