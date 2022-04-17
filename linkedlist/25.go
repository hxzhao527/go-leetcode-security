package linkedlist

// Name: 25. K 个一组翻转链表
// URL: https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	curr := dummy

	for remainCount(curr.Next) >= k {
		curr.Next = reverseNIter(curr.Next, k)
		// seek
		for i := 1; i <= k; i++ {
			curr = curr.Next
		}
	}
	return dummy.Next
}

func remainCount(head *ListNode) (cnt int) {
	curr := head
	for curr != nil {
		cnt += 1
		curr = curr.Next
	}
	return
}
