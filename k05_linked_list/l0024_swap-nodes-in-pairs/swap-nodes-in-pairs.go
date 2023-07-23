package l0024_swap_nodes_in_pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

// swapPairs 递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	point, NextPoint, next := head, head.Next, head.Next.Next
	NextPoint.Next = point
	newNode := swapPairs(next)
	point.Next = newNode
	return NextPoint
}

// swapPairs1 模拟
func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	MyHead, prev := &ListNode{}, &ListNode{}
	point, pointNext, next := head, head.Next, head.Next.Next
	pointNext.Next = point
	point.Next = next
	MyHead = pointNext
	prev = point
	point = next

	for point != nil && point.Next != nil {
		pointNext = point.Next
		next = pointNext.Next

		pointNext.Next = point
		point.Next = next
		prev.Next = pointNext

		prev = point
		point = next
	}

	return MyHead
}

// swapPairs2 模拟优化
func swapPairs2(head *ListNode) *ListNode {
	// 抽象一个虚拟头结点，以此结点为point，然后交换身后两个节点的顺序
	MyHead := &ListNode{}
	MyHead.Next = head
	point := MyHead

	for point.Next != nil && point.Next.Next != nil {
		firstNode := point.Next
		next := point.Next.Next.Next

		// cur 指向节点2
		point.Next = point.Next.Next
		point.Next.Next = firstNode
		point.Next.Next.Next = next

		point = point.Next.Next
	}

	return MyHead.Next
}

// swapPairs3 模拟优化错误
//func swapPairs3(head *ListNode) *ListNode {
//	MyHead := &ListNode{}
//	MyHead.Next = head
//	point := MyHead
//
//	for point.Next != nil && point.Next.Next != nil {
//		secondNode := point.Next.Next
//
//      这一步
//		point.Next.Next.Next = point.Next
//
//		point = point.Next.Next
//	}
//
//	return MyHead.Next
//}
