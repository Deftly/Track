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
