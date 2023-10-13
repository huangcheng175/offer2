package offer2

import (
	"fmt"
	"math"
	"testing"
)

func TestLeastLenArr(t *testing.T) {
	nums := []int{5, 1, 4, 3}
	target := 7

	fmt.Println(leastLenArr(nums, target))
}

func leastLenArr(nums []int, target int) int {
	least := math.MaxInt

	tmpSum := 0
	p1 := 0
	for p2 := 0; p2 < len(nums); p2++ {
		tmpSum += nums[p2]
		if tmpSum < target {
			continue
		}

		if p2-p1+1 < least {
			least = p2 - p1 + 1
		}

		// 左节点移动
		for p1 <= p2 {
			tmpSum -= nums[p1]
			p1++

			if tmpSum < target {
				break
			}

			if p2-p1+1 < least {
				least = p2 - p1 + 1
			}
		}
	}

	return least
}
