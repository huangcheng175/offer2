package offer2

import (
	"fmt"
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	fmt.Println(divide(15, 2))
	fmt.Println(divide(-15, 2))
	fmt.Println(divide(15, -2))
	fmt.Println(divide(-15, -2))
}

func divide(a, b int) int {
	// 先考虑溢出的问题，只有当a为math.MinInt32，并且b = -1
	if a == math.MinInt32 && b == -1 {
		return math.MaxInt32
	}

	negativeCnt := 2
	if a > 0 {
		negativeCnt--
		a = -a
	}
	if b > 0 {
		negativeCnt--
		b = -b
	}

	result := divideCore(a, b)
	if negativeCnt == 1 {
		return -result
	}

	return result

}

func divideCore(a, b int) int {
	var result int

	for a <= b {
		lb := b
		quotient := 1
		for lb >= math.MinInt32 && a <= lb+lb {
			quotient += quotient
			lb += lb
		}
		
		result += quotient
		a -= lb
	}

	return result
}
