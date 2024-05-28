package main

/* https://leetcode.com/problems/contains-duplicate/ */
func ContainsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _,num := range nums {
		_, exists := m[num]
		if exists {
			return true
		}
		m[num] = true;
	}
	return false
}
