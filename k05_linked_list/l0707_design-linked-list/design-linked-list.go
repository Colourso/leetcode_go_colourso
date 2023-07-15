package l0707_design_linked_list

import "fmt"

// 你可以选择使用单链表或者双链表，设计并实现自己的链表。
//
//单链表中的节点应该具备两个属性：val 和 next 。val 是当前节点的值，next 是指向下一个节点的指针/引用。
//
//如果是双向链表，则还需要属性 prev 以指示链表中的上一个节点。假设链表中的所有节点下标从 0 开始。
//
//实现 MyLinkedList 类：
//
//MyLinkedList() 初始化 MyLinkedList 对象。
//int get(int index) 获取链表中下标为 index 的节点的值。如果下标无效，则返回 -1 。
//void addAtHead(int val) 将一个值为 val 的节点插入到链表中第一个元素之前。在插入完成后，新节点会成为链表的第一个节点。
//void addAtTail(int val) 将一个值为 val 的节点追加到链表中作为链表的最后一个元素。
//void addAtIndex(int index, int val) 将一个值为 val 的节点插入到链表中下标为 index 的节点之前。如果 index 等于链表的长度，那么该节点会被追加到链表的末尾。如果 index 比长度更大，该节点将 不会插入 到链表中。
//void deleteAtIndex(int index) 如果下标有效，则删除链表中下标为 index 的节点。
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/design-linked-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// LinkNode 单链表节点
type LinkNode struct {
	Val  int
	Next *LinkNode
}

// MyLinkedList 单链表结构体 存储头结点
type MyLinkedList struct {
	Len  int
	Head *LinkNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{Head: &LinkNode{}, Len: 0}
}

func (this *MyLinkedList) Get(index int) int {
	if this == nil || this.Len <= index {
		return -1
	}

	point, i := &LinkNode{}, 0
	point = this.Head.Next
	for i < index {
		i++
		point = point.Next
	}
	return point.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	if this == nil {
		return
	}
	node, next := &LinkNode{Val: val}, &LinkNode{}
	this.Len++
	next = this.Head.Next
	node.Next = next
	this.Head.Next = node
	// node.Next,this.Head.Next = this.Head.Next,node ?
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this == nil {
		return
	}
	node, pre := &LinkNode{Val: val}, &LinkNode{}
	this.Len++
	pre = this.Head
	for pre.Next != nil {
		pre = pre.Next
	}
	pre.Next = node
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if this == nil || this.Len < index {
		return
	}
	node, pre, next := &LinkNode{Val: val}, &LinkNode{}, &LinkNode{}
	pre = this.Head
	for i := 0; i < index; i++ {
		pre = pre.Next
	}
	next = pre.Next
	pre.Next = node
	node.Next = next
	this.Len++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this == nil || this.Len <= index {
		return
	}

	pre, next := &LinkNode{}, &LinkNode{}
	pre = this.Head
	for i := 0; i < index; i++ {
		pre = pre.Next
	}
	next = pre.Next
	if next != nil {
		pre.Next = next.Next
	} else {
		pre.Next = nil
	}
	this.Len--
}

func (t *MyLinkedList) PrintList() {
	if t == nil {
		return
	}

	fmt.Print("LinkedList: ")
	next := &LinkNode{}
	next = t.Head.Next
	for next != nil && next.Next != nil {
		fmt.Print(next.Val)
		fmt.Print(" -> ")
		next = next.Next
	}
	if next != nil {
		fmt.Print(next.Val)
		fmt.Print(" .")
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
