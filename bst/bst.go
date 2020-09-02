package main

import "fmt"

// Node Structure
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// This method will insert a new value depending on the existing
// values that the current node have already stored
func (n *Node) Insert(k int) {
	// If the new value is greater than the current node value
	// it means that should be stored at the right side of the node
	if n.Key < k {
		// If the current right node is null, then we should create
		// a new instance for the right leaf with his respective value
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			// Otherwise we will add the value in the existing node, so as this is
			// a recursive call, it could be added on the right or the left side of the node
			n.Right.Insert(k)
		}
	} else if n.Key > k { // Value is lower than the current node value
		// Check if the left node is null
		if n.Left == nil {
			//
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

// Search a k value inside the BST
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		// move right
		return n.Right.Search(k)
	} else if n.Key > k {
		// move left
		return n.Left.Search(k)
	}

	return true
}

// Print in order the BST
func (n *Node) InOrder() {
	if n != nil {
		n.Left.InOrder()
		fmt.Println(n.Key)
		n.Right.InOrder()
	}
}

// This method will add one or more nodes to our BST
func (n *Node) AddNodes(nodes []int) {
	for i := 0; i < len(nodes); i++ {
		n.Insert(nodes[i])
	}
}
