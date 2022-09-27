package main

func MinHeightTree(numbers []int) *Node {
	return minTree(numbers, 0, len(numbers)-1)
}

func minTree(n []int, start, end int) *Node {
	if end < start {
		return nil
	}

	mid := (start + end) / 2
	node := NewNode(mid)
	node.Left = minTree(n, start, mid-1)
	node.Right = minTree(n, mid+1, end)

	return node
}
