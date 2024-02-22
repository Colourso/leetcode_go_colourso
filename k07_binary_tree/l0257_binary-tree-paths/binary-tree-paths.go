package l0257_binary_tree_paths

import "strconv"

/*
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。

叶子节点 是指没有子节点的节点。

https://leetcode.cn/problems/binary-tree-paths/description/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	result := []string{}
	if root == nil {
		return result
	}

	var backtracking func(root *TreeNode, path string)
	backtracking = func(root *TreeNode, path string) {
		// 前序遍历，根左右
		path += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			// 当左右子节点均为空时，表示到达叶子节点
			result = append(result, path)
			return
		}

		if root.Left != nil {
			path += "->"
			backtracking(root.Left, path)
			// 需要立即回溯, 参看回溯的流程， for循环递归结束后使用下一个
			path = path[:len(path)-2] // 这里-2是因为左闭右开，去除 -> 字符串
		}

		if root.Right != nil {
			path += "->"
			backtracking(root.Right, path)
			path = path[:len(path)-2]
		}
	}

	backtracking(root, "")
	return result
}
