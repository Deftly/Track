package leetcode

func produceExceptSelf(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	result[0] = 1

	// Calculate prefix product
	for i := 1; i < length; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// Calculate suffix product and multiply with prefix product
	suffixProduct := 1
	for i := length - 2; i >= 0; i-- {
		suffixProduct *= nums[i+1]
		result[i] *= suffixProduct
	}

	return result
}
