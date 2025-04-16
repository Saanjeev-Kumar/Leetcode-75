package main

import (
	"fmt"
)

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := make([]int, 0, n)
	fmt.Println(stack, "__+++++++++____", result)

	i := 0
	// Single loop iteration over temperatures
	for i < n {
		// If the stack is empty or current temperature is less than or equal to
		// the temperature at the index on top of the stack, push the current index.
		if len(stack) == 0 || temperatures[i] <= temperatures[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++ // Move to the next day only after pushing.
			fmt.Println(stack, "______", result)
		} else {
			// The current temperature is warmer than the temperature at stack's top.
			// Pop the index and calculate the days waited.
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[last] = i - last
			fmt.Println(stack, "++++++", result)
			// Do not increment i here since the same temperature might resolve more indexes.
		}
	}

	return result
}

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	ans := dailyTemperatures(temperatures)
	fmt.Println(ans) // Expected output: [1 1 4 2 1 1 0 0]
}
