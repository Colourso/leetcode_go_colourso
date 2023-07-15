package l0707_design_linked_list

import "testing"

func TestLinkedList(t *testing.T) {
	list := Constructor()
	list.AddAtHead(7)
	list.AddAtHead(2)
	list.AddAtHead(1)
	list.AddAtIndex(3, 0)
	list.DeleteAtIndex(2)
	list.AddAtHead(6)
	list.AddAtTail(4)
	list.Get(4)
	list.AddAtHead(4)
	list.AddAtIndex(5, 0)
	list.AddAtHead(6)
	list.PrintList()
}
