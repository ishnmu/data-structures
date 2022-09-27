package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// source: https://www.geeksforgeeks.org/find-if-there-is-a-path-between-two-vertices-in-a-given-graph/
func TestIsRouteExists_1(t *testing.T) {
	graph := map[int][]int{
		0: {1, 2},
		1: {2},
		2: {0, 3},
		3: {3},
	}
	assert.Equal(t, true, IsRouteExists(graph, 1, 3))
}

func TestIsRouteExists_2(t *testing.T) {
	graph := map[int][]int{
		0: {1, 2},
		1: {2},
		2: {0, 3},
		3: {3},
	}
	assert.Equal(t, false, IsRouteExists(graph, 3, 2))
}
