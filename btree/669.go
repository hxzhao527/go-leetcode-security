package btree

// Name: 669. 修剪二叉搜索树
// URL: https://leetcode-cn.com/problems/trim-a-binary-search-tree/
func trimBST(root *TreeNode, low int, high int) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Val < low || root.Val > high {
		var tmp *TreeNode
		if root.Left != nil {
			tmp = trimBST(root.Left, low, high)
		}
		if tmp == nil && root.Right != nil {
			tmp = trimBST(root.Right, low, high)
		}

		return tmp
	} else {
		root.Left = trimBST(root.Left, low, high)
		root.Right = trimBST(root.Right, low, high)
	}

	return root
}
