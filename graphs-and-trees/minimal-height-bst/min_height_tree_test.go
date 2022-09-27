package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMHT(t *testing.T) {
	n := MinHeightTree([]int{1, 2, 3, 4, 5})
	assert.Equal(t, []int{2, 0, 1, 3, 4}, n.PreOrderTraversal())
	assert.Equal(t, []int{0, 1, 2, 3, 4}, n.InOrderTraversal())
	assert.Equal(t, []int{1, 0, 4, 3, 2}, n.PostOrderTraversal())
}
