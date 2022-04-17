package btree

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}

	rootIdx := _find_where(inorder, root.Val)

	leftInOrder := inorder[:rootIdx]
	leftPreOrder := preorder[1 : 1+len(leftInOrder)]
	root.Left = buildTree(leftPreOrder, leftInOrder)

	rightInOrder := inorder[rootIdx+1:]
	rightPreOrder := preorder[1+len(leftInOrder):]
	root.Right = buildTree(rightPreOrder, rightInOrder)

	return root
}

func _find_where(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}
