package linkedlist

// Name: 24. 两两交换链表中的节点
// URL: https://leetcode-cn.com/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	return reverseKGroup(head, 2)
}
