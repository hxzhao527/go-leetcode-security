package btree

import (
	"container/list"
)

func largestValues(root *TreeNode) []int {
	type NodeWithLevel struct {
		node  *TreeNode
		level int
	}
	if root == nil {
		return nil
	}

	deque := list.New()
	deque.PushBack(NodeWithLevel{node: root, level: 1})

	ret := make([]int, 0)
	currLvl := 1
	currMax := root.Val
	for deque.Len() > 0 {
		node := deque.Front()
		deque.Remove(node)
		val := node.Value.(NodeWithLevel)
		if val.level != currLvl {
			ret = append(ret, currMax)
			currMax = val.node.Val
			currLvl = val.level
		} else {
			if val.node.Val > currMax {
				currMax = val.node.Val
			}
		}
		if val.node.Left != nil {
			deque.PushBack(NodeWithLevel{node: val.node.Left, level: val.level + 1})
		}
		if val.node.Right != nil {
			deque.PushBack(NodeWithLevel{node: val.node.Right, level: val.level + 1})
		}
	}

	ret = append(ret, currMax)
	return ret
}
