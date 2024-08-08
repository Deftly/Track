package leetcode

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	record := make(map[int]bool, len(nums))
	for _, v := range nums {
		if record[v] {
			return true
		}
		record[v] = true
	}
	return false
}

func containsDuplicateV2(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	record := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, exists := record[v]; exists {
			return true
		}
		record[v] = struct{}{}
	}
	return false
}
