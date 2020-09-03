package strings

func NonRepeatingChar(value string) byte {
	// Create a variable for the result
	// and a hashmap to store the frequency of the chars
	var result byte
	var charsMap = make(map[byte]int)
	// Get the char array from the given string (value)
	chars := []byte(value)
	// Loop through the chars array and add the ammount
	// of times that the chars appears
	for _, char := range chars {
		charsMap[char] = charsMap[char] + 1
	}
	// Loop the chars array one more time but this time
	// we will check for the first item that has 1 appearance
	for _, char := range chars {
		if charsMap[char] == 1 {
			// We have our result, so now it's time to break the cycle..
			result = char
			break
		}
	}
	// This value also could be null by the way...
	return result
}
