package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	s := "the sky is blue"
	result := reverseWords(s)
	fmt.Println(result) // Output: blue is sky the

	s = "  hello world  "
	result = reverseWords(s)
	fmt.Println(result) // Output: world hello

	s = "a good   example"
	result = reverseWords(s)
	fmt.Println(result) // Output: example good a
}
