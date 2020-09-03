package arrays

// This method will return the most frequent item that exists inside an array
func MostFrequentItem(a []int) int {
	freqs := make(map[int]int)
	var maxOccurs, maxItem = 0, 0

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
	// Create an array for the common elements that we could found
	var result []int
	// And will use two indexes to loop through the arrays items
	var idxA, idxB = 0, 0
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

// This method will check if the array b is rotation of the array a
func ArrayIsRotation(a []int, b []int) bool {
	// Array sizes are different?
	if len(a) != len(b) {
		// You shall not pass!
		return false
	}
	// Create two variables to check if the first element of a[0]
	// exists in the array b. That would be the entry point to
	// start checking the rotation
	var itemA, idxB = a[0], -1
	// Start looping the array b
	for i := 0; i < len(b); i++ {
		if b[i] == itemA {
			// So if the current item is equals to itemA,
			// then we have our entry point (idxB) and have to break the loop
			idxB = i
			break
		}
	}
	// Check if idxB has changed after the loop
	if idxB == -1 {
		// We didn't found the first element of the array a inside
		// the array b... so  it's impossible that a[] could be rotation of b[]
		return false
	}
	// Start looping the array a
	for i := 0; i < len(a); i++ {
		// Will calculate the current index for b[] to start comparing a[i] with b[j].
		// For this we will use the following formula: (indexB + i) % arrayA.Size
		// (I.E. assuming that idxB = 2 and len(a) = 4)
		//  > 1st iteration => (2 + 0) % 4 = 2
		//  > 2nd iteration => (2 + 1) % 4 = 3
		//  > 3th iteration => (2 + 2) % 4 = 0
		//  > 4th iteration => (2 + 3) % 4 = 1
		j := (idxB + i) % len(a)
		// Check if the elements in the current iteration are different
		if a[i] != b[j] {
			// Nope... they're no rotation from each other
			return false
		}
	}
	// And finally! If everything goes well, it means that a[] is rotation of b[]
	return true
}

// This method will compare if two int arrays are equals or not
func IntArrayEquals(a []int, b []int) bool {
	// If the sizes of the arrays are different...
	// we shouldn't bother to continue
	if len(a) != len(b) {
		return false
	}
	// Start loop through the arrays items
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
