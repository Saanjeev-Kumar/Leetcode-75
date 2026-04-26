func maximumSubarraySum(nums []int, k int) int64 {
    n := len(nums)
    if n < k {
        return 0
    }

    max := int64(0)

    // Slide window
    for i := 0; i <= n-k; i++ {
        win := Sum(nums[i : i+k])
        
        // Check if all elements in the window are distinct
        if isDistinct(nums[i : i+k]) {
            if int64(win) > max {
                max = int64(win)
            }
        }
    }

    return max
}

func Sum(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return sum
}

func isDistinct(window []int) bool {
    for i := 0; i < len(window); i++ {
        for j := i + 1; j < len(window); j++ {
            if window[i] == window[j] {
                return false
            }
        }
    }
    return true
}   