package btree

// Name: 116. 填充每个节点的下一个右侧节点指针
// URL: https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/
func connect(root *Node) *Node {
	// 可以用层序遍历的方式串
	// 需要额外的辅助空间存储节点
	// 也可以递归
	if root == nil {
		return nil
	}
	var traverse func(node1 *Node, node2 *Node)
	traverse = func(node1 *Node, node2 *Node) {
		if node1 == nil || node2 == nil {
			return
		}
		node1.Next = node2

		traverse(node1.Left, node1.Right)
		traverse(node2.Left, node2.Right)
		//
		traverse(node1.Right, node2.Left)
	}

	traverse(root.Left, root.Right)
	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
