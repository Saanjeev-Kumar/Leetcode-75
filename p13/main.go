package main

import (
        "fmt"
)

// func maxOperations(nums []int, k int) int {
//         count := make(map[int]int)
//         operations := 0

//         for _, num := range nums {
//                 complement := k - num
//                 if count[complement] > 0 {
//                         operations++
//                         count[complement]--
//                         count[num]--
//                 } else {
//                         count[num]++
//                 }
//         }

//         return operations
// }

func main() {
        nums := []int{1, 2, 3, 4}
        k := 5
        result := maxOperations(nums, k)
        fmt.Println(result) // Output: 2
}


func maxOperations(nums []int, k int) int {
        count := make(map[int]int)
        operations := 0

        for _, num := range nums {
                complement := k - num
                if count[complement] > 0 {
                        operations++
                        count[complement]--
                        count[num]--
                } else {
                        count[num]++
                }
        }

        return operations
}

