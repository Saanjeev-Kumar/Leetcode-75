//Logic: First

func maximumSubarraySum(nums []int, k int) int64{
    // n := len(nums)
    win := Sum(nums[:k])
    max := win
    // l := n-k
    for i:=0; i< len(nums)-k;i++{
        // win = win - nums[i] + nums[i+k]
        win1 := nums[i : i+k]
        if hasUniqueElements(win1) {
            fmt.Println(win1)
            max = Max(max,Sum(win1))
        }
    }
    return int64(max)
}

func hasUniqueElements(arr []int) bool {
    seen := make(map[int]bool)
    for _, num := range arr {
        if seen[num] {
            return false
        }
        seen[num] = true
    }
    return true
}

func Sum(nums []int)int{
    sum :=0
    for i,_:= range nums{
        sum= sum+nums[i]
    }
    return sum
}

func Max(a,b int) int{
    if a > b{
        return a
    }
    return b
}