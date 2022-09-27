package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_ShouldInitialiseWithEmpty(t *testing.T) {
	q := NewQueue()
	assert.Equal(t, true, q.IsEmpty())
}

func TestQueue_ShouldErrorIfDequeuePerformedOnEmpty(t *testing.T) {
	q := NewQueue()
	_, err := q.Dequeue()
	assert.NotNil(t, err)
	assert.Equal(t, "nothing to dequeue", err.Error())
}

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	dequeued1, err1 := q.Dequeue()
	assert.Nil(t, err1)
	assert.Equal(t, 1, dequeued1)

	dequeued2, err2 := q.Dequeue()
	assert.Nil(t, err2)
	assert.Equal(t, 2, dequeued2)

	dequeued3, err3 := q.Dequeue()
	assert.Nil(t, err3)
	assert.Equal(t, 3, dequeued3)

	dequeued4, err4 := q.Dequeue()
	assert.Nil(t, err4)
	assert.Equal(t, 4, dequeued4)
}
