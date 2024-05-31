package main

/* https://leetcode.com/problems/remove-duplicates-from-sorted-array/ */
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[j] = nums[i]
			j++
		}
	}
	nums[j] = nums[len(nums)-1]
	return j + 1
}
