package btree

// Name: 114. 二叉树展开为链表
// URL: https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	left, right := root.Left, root.Right

	root.Left, root.Right = nil, left

	curr := root
	for curr.Right != nil {
		curr = curr.Right
	}
	curr.Right = right
}
