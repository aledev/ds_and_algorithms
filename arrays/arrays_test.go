package arrays

import (
	"testing"
)

func TestMostFrequentItem(t *testing.T) {
	methodName := "MostFrequentItem"

	// Initial values
	items := []int{1, 2, 4, 4, 5, 10, 4, 1, 2, 4, 10}

	// Expected and result values
	expected := 4
	result := MostFrequentItem(items)

	// Comparison
	if result != expected {
		t.Errorf("%s method failed, expected %d, but got %d", methodName, expected, result)
		return
	}

	// Success!
	t.Logf("%s method success, expected %d and got %d", methodName, expected, result)
}

func TestCommonElements(t *testing.T) {
	methodName := "CommonElements"

	// Initial values
	a := []int{1, 3, 4, 6, 7, 9}
	b := []int{1, 2, 4, 5, 9, 10}

	// Expected and result values
	expected := []int{1, 4, 9}
	result := CommonElements(a, b)

	// Comparison
	if !IntArrayEquals(result, expected) {
		t.Errorf("%s method failed, expected %d, but got %d", methodName, expected, result)
		return
	}

	// Success!
	t.Logf("%s method success, expected %d and got %d", methodName, expected, result)
}
