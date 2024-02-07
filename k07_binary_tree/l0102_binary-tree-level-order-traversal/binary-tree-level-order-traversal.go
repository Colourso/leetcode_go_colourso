package l0102_binary_tree_level_order_traversal

/*
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。
（即逐层地，从左到右访问所有节点）。

https://leetcode.cn/problems/binary-tree-level-order-traversal/description/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		path := make([]int, n)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			path[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, path)
	}

	return result
}
