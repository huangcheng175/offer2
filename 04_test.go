package offer2

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	fmt.Println(singleNumber([]int{3, 3, 3, 121, 4, 4, 4}))
	fmt.Println(singleNumber1([]int{3, 3, 3, 3, 121, 4, 4, 4, 4}, 4))
}

func singleNumber(nums []int) int {
	x := [32]int{}
	for _, num := range nums {
		i := 0
		for num != 0 {
			x[i] += num & 1
			num = num >> 1
			i++
		}
	}

	cnt := 0
	for i, v := range x {
		if v%3 == 1 {
			cnt += 1 << i
		}

	}

	return cnt

}

func singleNumber1(nums []int, m int) int {
	x := [32]int{}
	for _, num := range nums {
		i := 0
		for num != 0 {
			x[i] += num & 1
			num = num >> 1
			i++
		}
	}

	cnt := 0
	for i, v := range x {
		if v%m == 1 {
			cnt += 1 << i
		}

	}

	return cnt

}
