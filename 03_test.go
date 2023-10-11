package offer2

import (
	"fmt"
	"testing"
)

func TestCountBits(t *testing.T) {

	fmt.Println(countBits(5))

	fmt.Println(countBits1(5))

	fmt.Println(countBits2(5))

}

func countBits2(n int) []int {

	result := make([]int, n+1)
	for i := 1; i <= n; i++ {
		result[i] = result[i>>1] + (i & 1)
	}

	return result
}

func countBits1(n int) []int {
	result := make([]int, n+1)
	for i := 1; i <= n; i++ {
		result[i] = result[i&(i-1)] + 1
	}

	return result
}

func countBits(n int) []int {
	v := make([]int, n+1)
	for i := 1; i <= n; i++ {
		v[i] = countXBits(i)
	}

	return v
}

func countXBits(n int) int {
	x := 0
	for n != 0 {
		x++
		n &= n - 1
	}

	return x
}
