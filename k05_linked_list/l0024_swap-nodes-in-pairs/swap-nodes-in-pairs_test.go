package l0024_swap_nodes_in_pairs

import "testing"

func TestSwap(t *testing.T) {

	listNode4 := &ListNode{Val: 4}
	listNode3 := &ListNode{Val: 3, Next: listNode4}
	listNode2 := &ListNode{Val: 2, Next: listNode3}
	listNode1 := &ListNode{Val: 1, Next: listNode2}

	swapPairs2(listNode1)
}
