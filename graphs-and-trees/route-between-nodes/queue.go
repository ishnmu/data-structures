package main

import "errors"

type Queue struct {
	data []int
}

func (q *Queue) Enqueue(d int) {
	q.data = append(q.data, d)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.data) > 0 {
		dequeued := q.data[0]
		q.data = q.data[1:]
		return dequeued, nil
	}
	return 0, errors.New("nothing to dequeue")
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]int, 0),
	}
}
