package arrays

// This method will return the most frequent item that exists inside an array
func MostFrequentItem(a []int) int {
	freqs := make(map[int]int)
	maxOccurs := 0
	maxItem := 0

	// Loop through the array items
	for i := 0; i < len(a); i++ {
		if freqs[a[i]] == 0 {
			// If we reach this point it means that our value doesn't exists yet
			// in the hashmap, so... we have to add it
			freqs[a[i]] = 1
			// And will check if our current item is our best candidate
			CompareAndUpdate(&maxOccurs, &maxItem, 1, a[i])
		} else {
			// We found an existing item, so we increase the frequency
			freqs[a[i]] = freqs[a[i]] + 1
			// Again... will check if our current item is our best candidate
			CompareAndUpdate(&maxOccurs, &maxItem, freqs[a[i]]+1, a[i])
		}
	}

	return maxItem
}

// This method will check if the candidate (canItem) is more frequent than the
// current best candidate (maxItem)
func CompareAndUpdate(maxOccurs *int, maxItem *int, canOccurs int, canItem int) {
	if *maxOccurs < canOccurs {
		*maxOccurs = canOccurs
		*maxItem = canItem
		return
	}
}

// This method will return an array of common elements between two sorted arrays.
// Otherwise, we should add a sorted method before we start with the comparison
func CommonElements(a []int, b []int) []int {
	var result []int
	idxA := 0
	idxB := 0

	// Start the loop through the items on the arrays and will stops
	// just when we reach the last index of one of the arrays
	for idxA < len(a) && idxB < len(b) {
		if a[idxA] == b[idxB] {
			// If element A == B then we have a common element
			result = append(result, a[idxA])
			// Increase the indexes for a and b arrays
			idxA++
			idxB++
		} else if a[idxA] > b[idxB] {
			// If element A > B then we have to increment the b array index
			idxB++
		} else {
			// Otherwise we have to increment the a array index
			idxA++
		}
	}
	// And finally we have our common elements array!
	return result
}

// This method will compare if two int arrays are equals or not
func IntArrayEquals(a []int, b []int) bool {
	// If the sizes of the arrays are different...
	// we shouldn't bother to continue
	if len(a) != len(b) {
		return false
	}
	// Start loop throught the arrays items
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
