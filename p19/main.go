package main

import "fmt"

func pivotIndex(nums []int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	leftSum := 0
	for i := 0; i < len(nums); i++ {
		rightSum := totalSum - leftSum - nums[i]
		if leftSum == rightSum {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}

func main() {
	nums1 := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums1)) // Output: 3

	nums2 := []int{1, 2, 3}
	fmt.Println(pivotIndex(nums2)) // Output: -1

	nums3 := []int{2, 1, -1}
	fmt.Println(pivotIndex(nums3)) // Output: 0
}
