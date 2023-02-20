package queue

import (
	"container/list"
)

type Queue struct {
	container *list.List
}

func (q *Queue) Push(value any) {
	q.container.PushFront(value)
}

func (q *Queue) Pop() any {
	if q.container.Len() > 0 {
		ele := q.container.Back()
		q.container.Remove(ele)
		return ele.Value
	}
	return nil
}

func (c *Queue) Front() any {
	if c.container.Len() > 0 {
		return c.container.Back().Value
	}
	return nil
}

func (c *Queue) Size() int {
	return c.container.Len()
}

func (c *Queue) Empty() bool {
	return c.container.Len() == 0
}

func New() *Queue {
	q := &Queue{
		container: list.New(),
	}

	return q
}
