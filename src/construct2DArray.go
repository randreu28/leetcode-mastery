package main

/* https://leetcode.com/problems/convert-1d-array-into-2d-array/d */
func Construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	result := make([][]int, m)

	for i := range result {
		result[i] = make([]int, n)
	}

	for index, value := range original {
		row := index / n
		col := index % n
		result[row][col] = value
	}

	return result
}
