package main

import "fmt"

func longestSubarray(nums []int) int {
	maxOnes := 0
	zeroCount := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		maxOnes = max(maxOnes, right-left)
	}

	return maxOnes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 1, 0, 1}
	result := longestSubarray(nums)
	fmt.Println(result) // Output: 3
}
