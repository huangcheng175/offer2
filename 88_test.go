package offer2

import (
	"fmt"
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1] + cost[0]
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	if len(cost) > 2 {
		dp[len(cost)-1] = min(dp[len(cost)-2], dp[len(cost)-1])
	}

	dp1 := make([]int, len(cost))
	dp1[0] = cost[1]
	dp1[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp1[i] = min(dp1[i-1], dp1[i-2]) + cost[i]
	}

	if len(cost) > 2 {
		dp1[len(cost)-1] = min(dp1[len(cost)-2], dp1[len(cost)-1])
	}

	return min(dp[len(dp)-1], dp1[len(dp)-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
