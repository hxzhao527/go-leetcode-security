package btree

// Name: 543. 二叉树的直径
// URL: https://leetcode-cn.com/problems/diameter-of-binary-tree/
func diameterOfBinaryTree(root *TreeNode) int {
	d := _diameterOfBinaryTree{}
	d.maxDepth(root)
	return d.max - 1
}

type _diameterOfBinaryTree struct {
	max int
}

func (m *_diameterOfBinaryTree) maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left := m.maxDepth(root.Left)
	right := m.maxDepth(root.Right)

	di := left + right + 1
	if di > m.max {
		m.max = di
	}
	if left > right {
		return left + 1
	}
	return right + 1
}
