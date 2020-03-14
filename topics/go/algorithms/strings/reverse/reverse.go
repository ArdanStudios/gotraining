package strings

func ReverseString(str string) string {

	// Create a slice of runes.
	runes := []rune(str)

	// Create an index that will traverse the collection of runes from the beginning to the end.
	var beg int

	// Create an index that will traverse the collection of runes from the end to the beginning.
	end := len(runes) - 1

	// Keep swapping runes until the two indexes meet in the middle.
	for beg < end {

		// Swap the code point.
		r := runes[beg]
		runes[beg] = runes[end]
		runes[end] = r

		// Move the indexes closer to each other working towards the middle of the collection.
		beg = beg + 1
		end = end - 1

	}

	return string(runes)
}
