package main

/* https://leetcode.com/problems/missing-number/ */
func MissingNumber(nums []int) int {
	n := len(nums)
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	for i := 0; i <= n; i++ {
		if !numSet[i] {
			return i
		}
	}
	return -1
}
