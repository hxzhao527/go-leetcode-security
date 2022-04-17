package linkedlist

import "strconv"

func buildListFromSlice(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := ListNode{
		Val: vals[0],
	}
	head.Next = buildListFromSlice(vals[1:])
	return &head
}

func (node *ListNode) Eq(other *ListNode) bool {
	if node == nil && other == nil {
		return true
	} else if node != nil && other == nil {
		return false
	} else if node == nil && other != nil {
		return false
	}
	if node.Val != other.Val {
		return false
	}
	return node.Next.Eq(other.Next)
}

func (node *ListNode) String() string {
	if node == nil {
		return "<nil>"
	}
	self := strconv.Itoa(node.Val)
	return self + "->" + node.Next.String()
}
