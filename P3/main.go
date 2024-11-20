package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	for _, candy := range candies {
		maxCandies = max(maxCandies, candy)
	}

	result := make([]bool, len(candies))
	for i, candy := range candies {
		result[i] = candy+extraCandies >= maxCandies
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	result := kidsWithCandies(candies, extraCandies)
	fmt.Println(result) // Output: [true true true false true]
}
