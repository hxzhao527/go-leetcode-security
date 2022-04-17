package btree

import (
	"container/list"
	"encoding/json"
)

func buildTreeByLevel(rawVals string) *TreeNode {
	if len(rawVals) == 0 {
		return nil
	}

	var vals []*int
	if err := json.Unmarshal([]byte(rawVals), &vals); err != nil {
		return nil
	}

	deque := list.New()

	root := &TreeNode{
		Val: *vals[0],
	}
	deque.PushBack(root)

	index := 1

	for deque.Len() > 0 {
		ele := deque.Front()
		deque.Remove(ele)
		node := ele.Value.(*TreeNode)
		if len(vals) <= index {
			break
		}
		if left := vals[index]; left != nil {
			node.Left = &TreeNode{Val: *left}
			deque.PushBack(node.Left)
		}
		index++
		if len(vals) <= index {
			break
		}
		if right := vals[index]; right != nil {
			node.Right = &TreeNode{Val: *right}
			deque.PushBack(node.Right)
		}
		index++
		if len(vals) <= index {
			break
		}
	}
	return root
}

func (tree *TreeNode) Eq(other *TreeNode) bool {
	if tree == nil && other == nil {
		return true
	} else if tree == nil && other != nil {
		return false
	} else if tree != nil && other == nil {
		return false
	}

	if tree.Val != other.Val {
		return false
	}

	return tree.Left.Eq(other.Left) && tree.Right.Eq(other.Right)
}
