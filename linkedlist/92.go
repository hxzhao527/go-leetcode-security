package linkedlist

// Name: 92. 反转链表 II
// URL: https://leetcode-cn.com/problems/reverse-linked-list-ii/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		// 从头开始反转, 变化head
		return reverseN(head, right)
	}
	// 不是从头, head保持不变
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

type _reverseN struct {
	successor *ListNode
}

func (s *_reverseN) reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		s.successor = head.Next
		return head
	}
	new_head := s.reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = s.successor
	return new_head
}

func reverseN(head *ListNode, n int) *ListNode {
	s := _reverseN{}
	return s.reverseN(head, n)
}

func reverseNIter(head *ListNode, n int) *ListNode {
	var new_head *ListNode
	var last *ListNode
	curr := head

	for ; n >= 1; n-- {
		node := curr
		curr = curr.Next
		node.Next = new_head
		new_head = node
		if last == nil {
			last = new_head
		}
	}

	last.Next = curr
	return new_head
}

func reverseBetweenIter(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseNIter(head, right)
	}
	curr := head
	for i := 1; i < left-1; i++ {
		curr = curr.Next
	}
	curr.Next = reverseNIter(curr.Next, right-left+1)
	return head
}
