package main

/* https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array */
func FindDisappearedNumbers(nums []int) []int {
	length := len(nums)
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	disappeared := make([]int, 0)
	for i := 1; i <= length; i++ {
		if !numSet[i] {
			disappeared = append(disappeared, i)
		}
	}
	return disappeared
}