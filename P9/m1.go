package main

import "fmt"

func compress_v1(chars []byte) int {
	n := len(chars)
	if n == 0 {
		return 0
	}

	writeIndex := 0
	readIndex := 0

	for readIndex < n {
		currentChar := chars[readIndex]
		count := 0
		for readIndex < n && chars[readIndex] == currentChar {
			readIndex++
			count++
		}

		chars[writeIndex] = currentChar
		writeIndex++

		if count > 1 {
			countStr := []byte(fmt.Sprintf("%d", count))
			for _, c := range countStr {
				chars[writeIndex] = c
				writeIndex++
			}
		}
	}

	return writeIndex
}

// chars := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
// length := compress(chars)
// fmt.Println("Compressed length:", length)
// fmt.Println("Compressed array:", string(chars[:length]))
