package strings

import (
	"testing"
)

//#region NonRepeatingChar

func TestNonRepeatingChar_ExistingValue(t *testing.T) {
	methodName := "NonRepeatingChar"

	// Initial values
	stringVal := "abbccdddeeeecccddddbbbffffff"

	// Expected and result values
	var expected byte = 'a'
	result := NonRepeatingChar(stringVal)

	// Comparison
	if result != expected {
		t.Errorf("%s method failed, expected %c, but got %c", methodName, expected, result)
		return
	}

	// Success!
	t.Logf("%s method success, expected %c and got %c", methodName, expected, result)
}

func TestNonRepeatingChar_NonExistingValue(t *testing.T) {
	methodName := "NonRepeatingChar"

	// Initial values
	stringVal := "aaabbccdddeeeecccddddbbbffffff"

	// Expected and result values
	var expected byte
	result := NonRepeatingChar(stringVal)

	// Comparison
	if result != expected {
		t.Errorf("%s method failed, expected %c, but got %c", methodName, expected, result)
		return
	}

	// Success!
	t.Logf("%s method success, expected %c and got %c", methodName, expected, result)
}

//#endregion
