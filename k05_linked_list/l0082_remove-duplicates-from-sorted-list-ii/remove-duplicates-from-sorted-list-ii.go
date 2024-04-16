package l0082_remove_duplicates_from_sorted_list_ii

/*
给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。

https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/description/?envType=study-plan-v2&envId=top-interview-150
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	Head := &ListNode{Next: head}
	p, pre := head, Head

	for p != nil {
		point := p
		for point != nil && p.Val == point.Val {
			point = point.Next
		}
		if p.Next == point {
			pre = p
			p = point
		} else {
			pre.Next = point
			p = point
		}
	}
	return Head.Next
}
