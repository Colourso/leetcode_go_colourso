package l0203_remove_linked_list_elements

// 给你一个链表的头节点 head 和一个整数 val ，
// 请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
// https://leetcode.cn/problems/remove-linked-list-elements/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	Head := &ListNode{Next: head}
	var pre, point = Head, head
	for point != nil {
		if point.Val == val {
			pre.Next = point.Next
		} else {
			pre = point
		}
		point = point.Next
	}
	return Head.Next
}

// removeElements1 转为for循环
func removeElements1(head *ListNode, val int) *ListNode {
	Head := &ListNode{Next: head}
	for point := Head; point.Next != nil; {
		if point.Next.Val == val {
			point.Next = point.Next.Next
		} else {
			point = point.Next
		}
	}
	return Head.Next
}
