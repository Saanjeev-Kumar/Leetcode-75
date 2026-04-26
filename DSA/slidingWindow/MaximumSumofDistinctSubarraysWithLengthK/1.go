func maximumSubarraySum(nums []int, k int) int64{
    // n := len(nums)
    win := Sum(nums[:k])
    max := win
    // l := n-k
    for i:=0; i< len(nums)-k;i++{
        win = win - nums[i] + nums[i+k]
        max = Max(max,win)

    }
    return int64(max)
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