package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMHT(t *testing.T) {
	n := MinHeightTree([]int{1, 2, 3, 4, 5})
	assert.Equal(t, []int{3, 1, 2, 4, 5}, n.PreOrderTraversal())
	assert.Equal(t, []int{1, 2, 3, 4, 5}, n.InOrderTraversal())
	assert.Equal(t, []int{2, 1, 5, 4, 3}, n.PostOrderTraversal())
}
