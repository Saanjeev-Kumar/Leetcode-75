package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	// Calculate the sum of the first k elements
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	// Initialize maxAverage with the average of the first k elements
	maxAverage := float64(sum) / float64(k)

	// Slide the window
	for i := k; i < len(nums); i++ {
		// Remove the first element of the previous window
		sum -= nums[i-k]
		// Add the new element to the window
		sum += nums[i]
		// Calculate the average of the current window
		currentAverage := float64(sum) / float64(k)
		// Update maxAverage if necessary
		maxAverage = max(maxAverage, currentAverage)
	}

	return maxAverage
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	result := findMaxAverage(nums, k)
	fmt.Println(result) // Output: 12.75
}
