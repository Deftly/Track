package leetcode

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	frequency := make([][]int, len(nums)+1)
	for k, v := range m {
		frequency[v] = append(frequency[v], k)
	}

	result := make([]int, 0, k)
	for i := len(frequency) - 1; i >= 0; i-- {
		for _, v := range frequency[i] {
			if k > 0 {
				result = append(result, v)
				k--
			}
		}
	}
	return result
}
