package main

func minRemoveToMakeValid(s string) string {
	var stack []int
	chars := []rune(s)

	// First pass: track indices of '(' and mark ')' for removal
	for i, char := range chars {
		switch char {
		case '(':
			stack = append(stack, i)
		case ')':
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				chars[i] = 0 // Mark ')' for removal
			}
		}
	}

	// Mark remaining '(' for removal
	for _, index := range stack {
		chars[index] = 0 // Mark '(' for removal
	}

	// Construct result string
	result := make([]rune, 0, len(chars))
	for _, char := range chars {
		if char != 0 {
			result = append(result, char)
		}
	}

	return string(result)
}

func minRemoveToMakeValidV2(s string) string {
	// Convert the string into a slice of runes (since Go strings are immutable)
	chars := []rune(s)
	var stack []int

	// First pass: track indices of '(' and mark ')' for removal
	for i, char := range chars {
		switch char {
		case '(':
			stack = append(stack, i) // Track index of unmatched '('
		case ')':
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // Pop from stack if matching '(' is found
			} else {
				chars[i] = 0 // Mark unmatched ')' for removal
			}
		}
	}

	// Mark any remaining unmatched '(' for removal
	for _, index := range stack {
		chars[index] = 0 // Mark unmatched '(' for removal
	}

	// In-place reconstruction of the valid string
	// Avoid allocating a new slice by overwriting invalid characters directly
	writeIndex := 0
	for _, char := range chars {
		if char != 0 { // Only copy valid characters
			chars[writeIndex] = char
			writeIndex++
		}
	}

	// Convert back to a string using only the valid portion of the slice
	return string(chars[:writeIndex])
}
