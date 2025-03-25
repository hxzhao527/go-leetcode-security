package queue

import "container/list"

type Queue struct {
	inner *list.List
}

func (queue *Queue) Len() int {
	return queue.inner.Len()
}

func (queue *Queue) Push(val interface{}) {
	queue.inner.PushBack(val)
}
func (queue *Queue) Pop() interface{} {
	head := queue.inner.Front()
	queue.inner.Remove(head)
	return head.Value
}

func (queue *Queue) Size() int {
	return queue.inner.Len()
}
