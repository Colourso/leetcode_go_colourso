package l0141_linked_list_cycle

// 给你一个链表的头节点 head ，判断链表中是否有环。
//
//如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/linked-list-cycle
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
// 怎么证明快慢指针在环里一定能遇到

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	low, fast := head, head.Next.Next
	for fast != nil && fast.Next != nil {
		if low == fast {
			return true
		}

		low = low.Next
		fast = fast.Next.Next
	}

	return false
}
