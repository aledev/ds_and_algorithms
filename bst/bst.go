package main

import "fmt"

// Node Structure
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// This method will insert a new value depending on the existing
// values that the current node have already stored
func (n *Node) Insert(v int) {
	// If the new value is greater than the current node value
	// it means that should be stored at the right side of the node
	if n.Value < v {
		// If the current right node is null, then we should create
		// a new instance for the right leaf with his respective value
		if n.Right == nil {
			n.Right = &Node{Value: v}
		} else {
			// Otherwise we will add the value in the existing node, so as this is
			// a recursive call, it could be added on the right or the left side of the node
			n.Right.Insert(v)
		}
	} else if n.Value > v { // Value is lower than the current node value
		// Check if the left node is null
		if n.Left == nil {
			//
			n.Left = &Node{Value: v}
		} else {
			n.Left.Insert(v)
		}
	}
}

// This method will search a "v" value inside the BST
func (n *Node) Search(v int) bool {
	// If the current node is null it means that the value
	// doesn't exists in the BST
	if n == nil {
		return false
	}
	// Check if the current value is greater than the current node value
	if n.Value < v {
		// In this case we will check if the value exists inside
		// the right node using recursion
		return n.Right.Search(v)
	} else if n.Value > v { // Value lower than the current node value
		// So in this case we will check if the value exists
		// in the left node, so pretty same as the previous case... we use recursion
		return n.Left.Search(v)
	}
	// So if we reach this point.. it seems that the current
	// node value is equals to the "v" value
	return true
}

// This method will print the BST in order
func (n *Node) InOrder() {
	// Check if the current node has instance
	if n != nil {
		// Do the recursion for the left side of the node
		n.Left.InOrder()
		// Print the current node value
		fmt.Println(n.Value)
		// Do the recursion for the right side of the node
		n.Right.InOrder()
	}
}

// This method will add k nodes to our BST
func (n *Node) AddNodes(nodes []int) {
	for i := 0; i < len(nodes); i++ {
		n.Insert(nodes[i])
	}
}

// This method will return the Lowest Common Ancestor
// for two given numbers: a, b.
func (n *Node) LowestCommonAncestor(a int, b int) (ancestor *Node) {
	// There's no future here boy... :(
	if n == nil {
		return nil
	}
	// If a and b are smaller than the current node,
	// then we should continue searching on the left child
	if n.Value > a && n.Value > b {
		return n.Left.LowestCommonAncestor(a, b)
	}
	// Otherwise if a and b are greater than the current node,
	// then definitely we should continue searching on the right child
	if n.Value < a && n.Value < b {
		return n.Right.LowestCommonAncestor(a, b)
	}
	// If we reach at this point... it means that here's our ancestor!
	return n
}
