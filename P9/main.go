package main

import "fmt"

func compress(chars []byte) int {
	slow, fast := 0, 0
	for fast < len(chars) {
		chars[slow] = chars[fast]
		count := 1
		for fast+1 < len(chars) && chars[fast] == chars[fast+1] {
			fast++
			count++
		}

		if count > 1 {
			for count > 0 {
				digit := byte(count%10 + '0')
				chars[slow+1] = digit
				slow++
				count /= 10
			}
		}
		slow++
		fast++
	}

	return slow
}

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	length := compress(chars)
	fmt.Println(chars[:length]) // Output: ["a","2","b","2","c","3"]

	chars_1 := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	length_1 := compress_v1(chars_1)
	fmt.Println("Compressed length:", length_1)
	fmt.Println("Compressed array:", string(chars[:length_1]))
}
