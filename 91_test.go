package offer2

import (
	"fmt"
	"testing"
)

func TestMinCost(t *testing.T) {

}

func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}

	if len(costs) == 1 {
		return minArray(costs[0])
	}

	dp := make([][]int, len(costs)+1)
	dp[0] = make([]int, 3)
	dp[1] = costs[0]

	for i := 2; i < len(dp); i++ {
		dp[i] = make([]int, 3)
		dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + costs[i-1][0]
		dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + costs[i-1][1]
		dp[i][2] = min(dp[i-1][0], dp[i-1][1]) + costs[i-1][2]
	}

	fmt.Println(dp)

	return minArray(dp[len(dp)-1])

}

func minArray(arr []int) int {
	minV := arr[0]
	for i := 1; i < len(arr); i++ {
		minV = min(minV, arr[i])
	}

	return minV
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
