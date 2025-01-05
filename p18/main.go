package main

import "fmt"

func largestAltitude(gain []int) int {
	maxAltitude := 0
	currentAltitude := 0

	for _, altitudeGain := range gain {
		currentAltitude += altitudeGain
		if currentAltitude > maxAltitude {
			maxAltitude = currentAltitude
		}
	}

	return maxAltitude
}

func main() {
	gain := []int{-5, 1, 5, 0, -7}
	result := largestAltitude(gain)
	fmt.Println(result) // Output: 1
}
