func maximumSubarraySum(nums []int, k int) int64 {
    n := len(nums)
    if n < k {
        return 0
    }
    
    // Frequency map for current window
    freq := make(map[int]int)
    win := 0
    
    // Initialize first window
    for i := 0; i < k; i++ {
        freq[nums[i]]++
        win += nums[i]
    }
    
    max := int64(0)
    if len(freq) == k { // All elements are distinct
        max = int64(win)
    }
    
    // Slide window
    for i := k; i < n; i++ {
        // Add new element
        freq[nums[i]]++
        win += nums[i]
        
        // Remove old element
        old := nums[i-k]
        win -= old
        freq[old]--
        if freq[old] == 0 {
            delete(freq, old)
        }
        
        // Update max if current window has all distinct elements
        if len(freq) == k {
            if int64(win) > max {
                max = int64(win)
            }
        }
    }
    
    return max
}   