package btree

// Name: 654. 最大二叉树
// URL: https://leetcode-cn.com/problems/maximum-binary-tree/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	idx, max := _find_max(nums)
	root := &TreeNode{
		Val: max,
	}
	root.Left = constructMaximumBinaryTree(nums[:idx])
	root.Right = constructMaximumBinaryTree(nums[idx+1:])
	return root
}

func _find_max(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, 0
	}
	idx := 0
	num := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > num {
			num = nums[i]
			idx = i
		}
	}
	return idx, num
}
