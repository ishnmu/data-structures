package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBST(t *testing.T) {
	bst := NewBST(10, nil, nil)
	assert.Equal(t, 10, bst.Node.Value)
	assert.Nil(t, bst.Node.Left)
	assert.Nil(t, bst.Node.Right)
}

func TestNode(t *testing.T) {
	node := NewNode(10)
	assert.Equal(t, 10, node.Value)
	assert.Nil(t, node.Left)
	assert.Nil(t, node.Right)
}
