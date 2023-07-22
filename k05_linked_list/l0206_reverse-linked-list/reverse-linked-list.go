package l0206_reverse_linked_list

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
// https://leetcode.cn/problems/reverse-linked-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	// 1.新建链表，采用头插法
	MyHead, point, next := &ListNode{}, &ListNode{}, &ListNode{}
	point = head
	for point != nil {
		next = point.Next

		point.Next = MyHead.Next
		MyHead.Next = point

		point = next
	}
	return MyHead.Next
}

// reverseList1 2.递归
func reverseList1(head *ListNode) *ListNode {
	// 递归的结束条件
	if head == nil || head.Next == nil {
		return head
	}

	node := reverseList1(head.Next)
	// head.Next 是原来的第二个节点，需要指向原来的头结点
	head.Next.Next = head
	// 原来的头结点指向末尾
	head.Next = nil
	return node
}
