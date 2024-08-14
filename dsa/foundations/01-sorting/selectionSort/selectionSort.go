package main

import "fmt"

func main() {
	nums := []int{100, 82, 83, 24, 26, 32, 15, 16, 88, 0, -1, -32, 28}
	fmt.Println(nums)
	selectionsSort(nums)
	fmt.Println(nums)
}

func selectionsSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	var min int
	for i := 0; i < len(nums)-1; i++ {
		min = i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
}
