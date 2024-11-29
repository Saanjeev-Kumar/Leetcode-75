package main

import "fmt"

func isSubsequence(s string, t string) bool {
	sIndex, tIndex := 0, 0

	for sIndex < len(s) && tIndex < len(t) {
		if s[sIndex] == t[tIndex] {
			sIndex++
		}
		tIndex++
	}

	return sIndex == len(s)
}

func main() {
	s := "abc"
	t := "ahbgdc"
	result := isSubsequence(s, t)
	fmt.Println(result) // Output: true
}
