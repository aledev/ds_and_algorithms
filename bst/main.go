package main

import (
	"fmt"
)

func main() {
	tree := &Node{Key: 100}

	fmt.Println(tree)

	nodes := []int{52, 203, 19, 76, 150, 310, 7, 24, 88, 276}
	tree.AddNodes(nodes)

	// Should return true
	fmt.Printf("Does %d exists: %t\n", 88, tree.Search(88))
	// Should return false
	fmt.Printf("Does %d exists: %t\n", 765, tree.Search(765))

	// Print the best in order
	tree.InOrder()
}
