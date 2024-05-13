package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letters := [26]int{}
	for i := 0; i < len(s); i++ {
		letters[s[i]-'a']++
		letters[t[i]-'a']--
	}

	for _, v := range letters {
		if v != 0 {
			return false
		}
	}
	return true
}

func isAnagramRune(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letters := map[rune]int{}
	for _, r := range s {
		letters[r]++
	}

	for _, r := range t {
		letters[r]--
		if letters[r] < 0 {
			return false
		}
	}

	for _, value := range letters {
		if value != 0 {
			return false
		}
	}

	return true
}

func isAnagramRuneV2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	characters := make(map[rune]int)
	for _, r := range s {
		characters[r]++
	}

	for _, r := range t {
		count, ok := characters[r]
		if !ok {
			return false
		}
		if count == 1 {
			delete(characters, r)
		} else {
			characters[r]--
		}
	}

	return len(characters) == 0
}
