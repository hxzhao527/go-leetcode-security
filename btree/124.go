package btree

// Name: 124. 二叉树中的最大路径和
// URL: https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
func maxPathSum(root *TreeNode) int {
	m := _maxPathSum{max: root.Val}
	m.maxPathSum(root)

	return m.max
}

type _maxPathSum struct {
	max int
}

func (m *_maxPathSum) maxPathSum(root *TreeNode) int {
	if root == nil {
		return -1001
	}

	left := m.maxPathSum(root.Left)
	right := m.maxPathSum(root.Right)

	maxPath := m.maxNum([]int{left + root.Val, right + root.Val, root.Val})
	m.max = m.maxNum([]int{maxPath, left, right, m.max, left + right + root.Val})

	return maxPath
}

func (m *_maxPathSum) maxNum(nums []int) int {
	_max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > _max {
			_max = nums[i]
		}
	}
	return _max
}
