func getSubarrayBeauty(nums []int, k int, x int) []int {
	sorted := [101]int{}
	l := 0

	beauty := make([]int, 0)
	for r := 0; r < len(nums); r++ {
		sorted[50+nums[r]]++

		if r-l+1 == k {
            answ := xMin(sorted, x)
            if answ > 0 {
                answ = 0
            }
			beauty = append(beauty, answ)
			sorted[50+nums[l]]--
			l++
		}
	}
	return beauty
}

func xMin(nums [101]int, x int) int {
	skip := 0
	for i := -50; i <= 50; i++ {
		if nums[i+50] != 0 {
			skip += nums[i+50]
		}
		if skip >= x {
			return i
		}
	}
	return 0
}