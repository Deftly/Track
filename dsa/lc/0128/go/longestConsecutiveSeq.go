package leetcode

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool, len(nums))
	for _, n := range nums {
		numSet[n] = true
	}

	longest := 0
	for n := range numSet {
		if !numSet[n-1] {
			currentNum := n
			currentStreak := 1

			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if currentStreak > longest {
				longest = currentStreak
			}
		}
	}
	return longest
}
