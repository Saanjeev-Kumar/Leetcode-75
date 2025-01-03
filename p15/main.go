package main

import (
	"fmt"
	"strings"
)

func maxVowels(s string, k int) int {
	vowels := "aeiouAEIOU"
	maxVowels := 0
	currentVowels := 0

	// Count vowels in the first k characters
	for i := 0; i < k; i++ {
		if strings.Contains(vowels, string(s[i])) {
			currentVowels++
		}
	}

	maxVowels = currentVowels

	// Slide the window
	for i := k; i < len(s); i++ {
		// Remove the first character of the previous window
		if strings.Contains(vowels, string(s[i-k])) {
			currentVowels--
		}
		// Add the new character to the window
		if strings.Contains(vowels, string(s[i])) {
			currentVowels++
		}
		// Update maxVowels if necessary
		maxVowels = max(maxVowels, currentVowels)
	}

	return maxVowels
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "aeiou"
	k := 2
	result := maxVowels(s, k)
	fmt.Println(result) // Output: 3
}
