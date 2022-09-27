package main

type BST struct {
	Node *Node
}

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

func NewBST(v int, left, right *Node) *BST {
	return &BST{
		Node: &Node{
			Left:  left,
			Right: right,
			Value: v,
		},
	}
}

func NewNode(v int) *Node {
	return &Node{
		Left:  nil,
		Right: nil,
		Value: v,
	}
}

func (n *Node) PreOrderTraversal() []int {
	numbers := make([]int, 0)

	if n != nil {
		numbers = append(numbers, n.Value)
		lefts := n.Left.PreOrderTraversal()
		rights := n.Right.PreOrderTraversal()

		numbers = append(numbers, lefts...)
		numbers = append(numbers, rights...)
	}
	return numbers
}

func (n *Node) InOrderTraversal() []int {
	numbers := make([]int, 0)

	if n != nil {
		lefts := n.Left.InOrderTraversal()
		numbers = append(numbers, lefts...)

		numbers = append(numbers, n.Value)

		rights := n.Right.InOrderTraversal()
		numbers = append(numbers, rights...)
	}
	return numbers
}

func (n *Node) PostOrderTraversal() []int {
	numbers := make([]int, 0)

	if n != nil {
		lefts := n.Left.PostOrderTraversal()
		numbers = append(numbers, lefts...)

		rights := n.Right.PostOrderTraversal()
		numbers = append(numbers, rights...)

		numbers = append(numbers, n.Value)
	}
	return numbers
}
