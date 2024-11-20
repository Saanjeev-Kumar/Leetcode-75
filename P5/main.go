package main

import (
	"fmt"
)

func reverseVowels(s string) string {
	bytes := []byte(s)
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	left, right := 0, len(bytes)-1

	for left < right {
		if !vowels[bytes[left]] {
			left++
			continue
		}
		if !vowels[bytes[right]] {
			right--
			continue
		}
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left++
		right--
	}

	return string(bytes)
}

func main() {
	s := "IceCreAm"
	result := reverseVowels(s)
	fmt.Println(result) // Output: AceCreIm

	s = "leetcode"
	result = reverseVowels(s)
	fmt.Println(result) // Output: leotcede
}
