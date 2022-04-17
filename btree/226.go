package btree

// Name: 226. 翻转二叉树
// URL: https://leetcode-cn.com/problems/invert-binary-tree
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
