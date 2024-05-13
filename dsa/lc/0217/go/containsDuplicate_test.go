package leetcode

import "testing"

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{[]int{}, false},
		{[]int{1}, false},
	}

	for _, tc := range testCases {
		result := containsDuplicate(tc.nums)
		if result != tc.expected {
			t.Errorf("containsDuplicate: expected %v, but got %v for input %v", tc.expected, result, tc.nums)
		}

		result = containsDuplicateV2(tc.nums)
		if result != tc.expected {
			t.Errorf("containsDuplicateV2: expected %v, but got %v for input %v", tc.expected, result, tc.nums)
		}
	}
}
