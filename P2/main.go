package main

import "fmt"

func gcdOfStrings(str1, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	// Find the greatest common divisor (GCD) of the lengths
	gcd := gcd(len(str1), len(str2))

	return str1[:gcd]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	str1 := "ABCABC"
	str2 := "ABC"
	result := gcdOfStrings(str1, str2)
	fmt.Println(result) // Output: ABC

	str1 = "ABABAB"
	str2 = "ABAB"
	result = gcdOfStrings(str1, str2)
	fmt.Println(result) // Output: AB

	str1 = "LEET"
	str2 = "CODE"
	result = gcdOfStrings(str1, str2)
	fmt.Println(result) // Output:
}
