package main

import "testing"

func TestSearch_ValueExists(t *testing.T) {
	methodName := "Node.Search"

	// Initial values
	tree := &Node{Value: 4}
	nodes := []int{52, 203, 19, 76, 150, 310, 7, 24, 88, 276}
	tree.AddNodes(nodes)

	// Expected and result values
	expected := true
	result := tree.Search(150)

	// Comparison
	if result != expected {
		t.Errorf("%s method failed, expected %t, but got %t", methodName, expected, result)
		return
	}

	// Success!
	t.Logf("%s method success, expected %t and got %t", methodName, expected, result)
}

func TestSearch_ValueNotExists(t *testing.T) {
	methodName := "Node.Search"

	// Initial values
	tree := &Node{Value: 4}
	nodes := []int{52, 203, 19, 76, 150, 310, 7, 24, 88, 276}
	tree.AddNodes(nodes)

	// Expected and result values
	expected := false
	result := tree.Search(790)

	// Comparison
	if result != expected {
		t.Errorf("%s method failed, expected %t, but got %t", methodName, expected, result)
		return
	}

	// Success!
	t.Logf("%s method success, expected %t and got %t", methodName, expected, result)
}
