package main

import "fmt"

func main() {
	nums := []int{2, 4, 3, 1, 6, 8, 5}
	fmt.Println(nums)
	fmt.Println(selectionSort(nums))
}

func selectionSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	var min, temp int
	for i := 0; i < len(nums)-1; i++ {
		min = i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		temp = nums[i]
		nums[i] = nums[min]
		nums[min] = temp
	}
	return nums
}
