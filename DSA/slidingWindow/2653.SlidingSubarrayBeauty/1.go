// First logic --> Wrong Answer
//625 / 717 testcases passed --> nums := []int{-46, -34, -46} k := 3 x := 3 //Logic 2
// Now Time Limit Exceeded 714 / 717 testcases passed

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, 世界")
	nums := []int{1, -1, -3, -2, 3}
	k := 3
	x := 2
	res1 := getSubarrayBeauty(nums, k, x)
	fmt.Println("========>", res1)

}

func getSubarrayBeauty(nums []int, k int, x int) []int {
	arr := []int{}
	res := make(map[int]int)
	for i := 0; i < k; i++ {
		res[nums[i]]++
	}
	fmt.Println(res)
	
	// count, win := 0, 0 //Worked with 1st testcase only
	// for k, _ := range res {
	// 	count++
	// 	if count == x {
	// 		win = k
	// 	}
	// }

	win := 0
	keys := make([]int, 0, len(res))
	for k := range res {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	// if len(keys) >= x { //logic 1
	// 	win = keys[x-1] 
	// }
	for _, key := range keys { //logic 2
		count += res[key] // Add the frequency of the current key
		if count >= x {   // Once we hit or pass 'x', we found our number!
			win = key
			break
		}
	}
	if win > 0 {
        win = 0
    }

	fmt.Println(win)
	arr = append(arr, win)
	fmt.Println("-->", arr)

	for i := k; i < len(nums); i++ {
		res[nums[i]]++
		fmt.Println(res)
		old := nums[i-k]
		res[old]--
		fmt.Println(res)
		if res[old] == 0 {
			delete(res, old)
		}

		fmt.Println("---->", res)
		keys := make([]int, 0, len(res))
		for k := range res {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		if len(keys) >= x {
			win = keys[x-1] // xth index
		}
		if win > 0 {
        win = 0
        }
		arr = append(arr, win)

	}
	return arr
}
