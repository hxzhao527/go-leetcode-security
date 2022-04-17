package linkedlist

// Name: 234. 回文链表
// URL: https://leetcode-cn.com/problems/palindrome-linked-list/
func isPalindrome(head *ListNode) bool {
	// a,b,c,... [mid] ...c,b,a
	// 先找到中点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}

	left := head
	// 然后反转
	right := reverseListIter(slow)
	for right != nil {
		if right.Val != left.Val {
			return false
		}
		right = right.Next
		left = left.Next
	}
	return true
}

func reverseListIter(head *ListNode) *ListNode {
	var new_head *ListNode
	curr := head
	for curr != nil {
		node := curr
		curr = node.Next
		node.Next = new_head
		new_head = node
	}
	return new_head
}
