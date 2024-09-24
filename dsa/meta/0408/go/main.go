package main

func validWordAbbreviation(word string, abbr string) bool {
	m, n := len(word), len(abbr)
	if m == 0 {
		return n == 0
	}

	i := 0
	j := 0
	for i < m && j < n {
		if abbr[j] >= '0' && abbr[j] <= '9' {
			if abbr[j] == '0' {
				return false
			}
			skip := 0
			for j < n && (abbr[j] >= '0' && abbr[j] <= '9') {
				skip *= 10
				skip += int(abbr[j] - '0')
				j++
			}
			i += skip
		} else {
			if word[i] != abbr[j] {
				return false
			}
			i++
			j++
		}
	}

	return i == m && j == n
}

func validWordAbbreviationV2(word string, abbr string) bool {
	m, n := len(word), len(abbr)
	i, j := 0, 0

	for i < m && j < n {
		if abbr[j] >= '0' && abbr[j] <= '9' { // Check if current char is a digit
			// Invalid if the number starts with '0'
			if abbr[j] == '0' {
				return false
			}

			// Parse the number and move j accordingly
			skip := 0
			for j < n && abbr[j] >= '0' && abbr[j] <= '9' {
				skip = skip*10 + int(abbr[j]-'0')
				j++
			}

			// Move the `i` pointer by `skip` characters in the word
			i += skip

		} else { // If it's not a digit, check for character match
			if word[i] != abbr[j] {
				return false
			}
			i++
			j++
		}
	}

	// Ensure both pointers reached the end of their respective strings
	return i == m && j == n
}
