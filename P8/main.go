package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	firstNum, secondNum := math.MaxInt32, math.MaxInt32

	for _, num := range nums {
		if num <= firstNum {
			firstNum = num
		} else if num <= secondNum {
			secondNum = num
		} else {
			return true
		}
	}

	return false
}

func main() {
	nums := []int{2, 1, 5, 0, 4, 6}
	result := increasingTriplet(nums)
	fmt.Println(result) // Output: true
}
