package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// Initialize the answer array with 1s
	for i := 0; i < n; i++ {
		answer[i] = 1
	}

	// Calculate the prefix product
	prefixProduct := 1
	for i := 0; i < n; i++ {
		answer[i] *= prefixProduct
		prefixProduct *= nums[i]
	}
	//[1, 2, 6, 24]
	// Calculate the postfix product
	postfixProduct := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= postfixProduct
		postfixProduct *= nums[i]
	}

	return answer
}

func main() {
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	fmt.Println(result) // Output: [24, 12, 8, 6]
}
