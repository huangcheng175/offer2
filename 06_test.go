package offer2

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	v := []int{1, 3, 5, 6, 7, 10}
	target := 10
	fmt.Println(twoSum(v, target))
}

func twoSum(nums []int, target int) []int {
	var p1, p2 int = 0, len(nums) - 1

	var tmpSum int = 0

	for p1 < p2 {
		tmpSum = nums[p1] + nums[p2]
		if tmpSum == target {
			return []int{p1, p2}
		} else if tmpSum < target {
			p1++
		} else {
			p2--
		}
	}

	return []int{-1, -1}
}
