package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			count++
			if count >= n {
				return true
			}
		}
	}
	return false
}

func main() {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1
	result := canPlaceFlowers(flowerbed, n)
	fmt.Println(result) // Output: true
}
