package offer2

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	return max(rob1(nums[1:]), rob1(append(nums[:len(nums)-1])))
}

func rob1(nums []int) int {
	dp := make([]int, len(nums)+2)
	for i := 0; i < len(nums); i++ {
		dp[i+2] = max(dp[i]+nums[i], dp[i+1])
	}

	return dp[len(nums)+1]

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
