package main

import "fmt"

func longestOnes(nums []int, k int) int {
	left := 0
	maxLen := 0
	zeroes := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroes++
		}

		// If the number of zeroes exceeds k, we need to shrink the window
		for zeroes > k {
			if nums[left] == 0 {
				zeroes--
			}
			left++
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2
	result := longestOnes(nums, k)
	fmt.Println(result) // Output: 6
}
