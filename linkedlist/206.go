package linkedlist

// Name: 206. 反转链表
// URL: https://leetcode-cn.com/problems/reverse-linked-list
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	new_head := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return new_head
}
