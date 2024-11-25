package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// Initialize the answer array with 1s
	for i := 0; i < n; i++ {
		answer[i] = 1
		fmt.Println(answer[i])
	}
	fmt.Println(">>>>>>>", answer)
	// Calculate the prefix product
	prefixProduct := 1
	for i := 0; i < n; i++ {
		answer[i] *= prefixProduct
		fmt.Println(answer[i], "----------", i)
		prefixProduct *= nums[i]
		fmt.Println(prefixProduct)
	}
	fmt.Println(">>>>>>>", answer)
	// Calculate the postfix product
	postfixProduct := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= postfixProduct
		fmt.Println(answer[i], "=========", i)
		postfixProduct *= nums[i]
		fmt.Println(postfixProduct)
	}
	fmt.Println(">>>>>>>", answer)
	return answer
}

func main() {
	nums := []int{1, 2, 3, 4}
	result := productExceptSelf(nums)
	fmt.Println(result) // Output: [24, 12, 8, 6]
}
