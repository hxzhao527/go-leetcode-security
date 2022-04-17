package btree

type _maxDepth struct {
	max int

	curr int
}

func (m *_maxDepth) maxDepth(root *TreeNode) {
	if root == nil {
		if m.curr > m.max {
			m.max = m.curr
		}
		return
	}
	m.curr++
	m.maxDepth(root.Left)
	m.maxDepth(root.Right)
	m.curr--
}

// Name: 104. 二叉树的最大深度
// URL: https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
func maxDepth(root *TreeNode) int {
	m := _maxDepth{}
	m.maxDepth(root)
	return m.max
}
