package btree

func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: postorder[len(postorder)-1],
	}
	rootIdx := _find_where(inorder, root.Val)

	leftInOrder := inorder[:rootIdx]
	leftPostOrder := postorder[:len(leftInOrder)]
	root.Left = buildTree2(leftInOrder, leftPostOrder)

	rightInOrder := inorder[rootIdx+1:]
	rightPostOrder := postorder[len(leftInOrder) : len(postorder)-1]
	root.Right = buildTree2(rightInOrder, rightPostOrder)

	return root
}
