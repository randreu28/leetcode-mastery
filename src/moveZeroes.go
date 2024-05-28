package main

/* https://leetcode.com/problems/move-zeroes/ */
func MoveZeroes(nums []int) {
	res := []int{}
	zeros := []int{}

	for _, value := range nums {
		if value == 0 {
			zeros = append(zeros, value)
		} else {
			res = append(res, value)
		}
	}

	res = append(res, zeros...)

	copy(nums, res)
}

func MoveZeroes2(nums []int) {
	lastNonZeroFoundAt := 0 // Pointer for the position of the last non-zero element found

	// Move all non-zero elements to the beginning of the array
	for i := range nums {
		if nums[i] != 0 {
			nums[lastNonZeroFoundAt] = nums[i]
			lastNonZeroFoundAt++
		}
	}

	// Fill the remaining array with zeros
	for i := lastNonZeroFoundAt; i < len(nums); i++ {
		nums[i] = 0
	}
}
