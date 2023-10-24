package offer2

import (
	"fmt"
	"math"
	"testing"
)

type A struct {
	a int
}

func aa(as []A) {
	for i := range as {
		as[i].a = 2
	}
}

func TestMinEatingSpeed(t *testing.T) {
	as := []A{
		{1},
		{1},
		{1},
	}
	aa(as)

	fmt.Println(as)

	//piles := []int{30, 11, 23, 4, 20}
	//h := 6
	//t.Log(minEatingSpeed(piles, h))
}

func minEatingSpeed(piles []int, h int) int {

	if len(piles) > h {
		return -1
	}

	start := 1
	var end int = math.MinInt
	for _, v := range piles {
		if v > end {
			end = v
		}
	}

	var n int

	for start <= end {

		n = (start + end) / 2
		//	fmt.Println("........", start, end, calEatingSpendTime(piles, n), calEatingSpendTime(piles, n+1), calEatingSpendTime(piles, n-1), n)
		//if calEatingSpendTime(piles, n) >= h && calEatingSpendTime(piles, n+1) == h {
		//	return n
		//}

		if calEatingSpendTime(piles, n) >= h && calEatingSpendTime(piles, n+1) < h {
			//	fmt.Println("a.......")
			return n + 1
		}
		//if calEatingSpendTime(piles, n) <= h && calEatingSpendTime(piles, n-1) >= h {
		//	fmt.Println("b.......")
		//	return n
		//}

		if calEatingSpendTime(piles, n) > h {
			start = n + 1
		} else {
			end = n - 1
		}
		//fmt.Println("........", start, end)

	}

	return -1
}

func calEatingSpendTime(piles []int, n int) int {
	sum := 0
	for _, pile := range piles {
		sum += pile / n
		if pile%n != 0 {
			sum++
		}
	}

	return sum
}
