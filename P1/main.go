package main

import "fmt"

func mergeStrings(word1, word2 string) string {
	merged := ""
	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		merged += string(word1[i]) + string(word2[j])
		i++
		j++
	}

	// Append remaining characters from the longer string
	merged += word1[i:] + word2[j:]

	return merged
}

func main() {
	word1 := "abc"
	word2 := "pqr"
	result := mergeStrings(word1, word2)
	fmt.Println(result) // Output: apbqcr
}
