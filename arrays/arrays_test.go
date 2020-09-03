package arrays

import (
	"testing"
)

//#region MostFrequentItem

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

//#endregion

//#region CommonElements

func TestCommonElements_ReturnArrayWithValues(t *testing.T) {
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

func TestCommonElements_ReturnEmptyArray(t *testing.T) {
	methodName := "CommonElements"

	// Initial values
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{7, 8, 9, 10, 11, 12}

	// Expected and result values
	expected := []int{}
	result := CommonElements(a, b)

	// Comparison
	if !IntArrayEquals(result, expected) {
		t.Errorf("%s method failed, expected %d, but got %d", methodName, expected, result)
		return
	}

	// Success!
	t.Logf("%s method success, expected %d and got %d", methodName, expected, result)
}

//#endregion

//#region ArrayIsRotation

func TestArrayIsRotation_TrueCase(t *testing.T) {
	methodName := "ArrayIsRotation"

	// Initial values
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{4, 5, 6, 7, 1, 2, 3}

	// Expected and result values
	expected := true
	result := ArrayIsRotation(a, b)

	// Comparison
	if result != expected {
		t.Errorf("%s method failed, expected %t, but got %t", methodName, expected, result)
		return
	}

	// Success!
	t.Logf("%s method success, expected %t and got %t", methodName, expected, result)
}

func TestArrayIsRotation_FalseCase(t *testing.T) {
	methodName := "ArrayIsRotation"

	// Initial values
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{5, 7, 6, 4, 1, 2, 3}

	// Expected and result values
	expected := false
	result := ArrayIsRotation(a, b)

	// Comparison
	if result != expected {
		t.Errorf("%s method failed, expected %t, but got %t", methodName, expected, result)
		return
	}

	// Success!
	t.Logf("%s method success, expected %t and got %t", methodName, expected, result)
}

//#endregion
