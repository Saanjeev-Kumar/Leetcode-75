Optimized Logic:

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

/*
Func maxDistinctWindow (nuts []int, k int) int{
	if k > n{
		return 0
	}
	freq := make(map[int]int)
	for I:= 0; I< k; I++ {
		freq[num[I]]++
		win += number[I]
	}
	max = win
	if len(freq) == k {
		return max
	}
	
	for I= k ; I < n; I++ {
		freq[num[I]]++
		win += num[I]
		old := num[I-k]
		freq[old]—
		win -= old
		if freq[old] == 0{
			delete(freq, old)
		}
		if len(freq) == k{
			if max < win {
				max = win
			}
		}
	}
	return max
}
*/