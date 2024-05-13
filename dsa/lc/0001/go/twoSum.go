package leetcode

func twoSum(nums []int, target int) []int {
	complement := make(map[int]int)
	for i, v := range nums {
		if idx, ok := complement[target-v]; ok {
			return []int{idx, i}
		}
		complement[v] = i
	}
	return nil
}
