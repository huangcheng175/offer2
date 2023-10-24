package offer2

import "testing"

func TestSqrt(t *testing.T) {
	for i := 0; i < 100; i++ {

		println(i, "........", sqrt(i))
	}
}

func sqrt(x int) int {
	return sqrt1(x, 0, x)

}

func sqrt1(x int, start, end int) int {
	if x < 0 {
		return -1
	}
	if x == 0 || x == 1 {
		return x
	}

	cur := (start + end) / 2

	if cur*cur == x {
		return cur
	}

	if cur*cur < x && (cur+1)*(cur+1) > x {
		return cur
	}
	if cur*cur > x && (cur-1)*(cur-1) < x {
		return cur - 1
	}

	if cur*cur > x {
		return sqrt1(x, start, cur-1)
	}
	if cur*cur < x {
		return sqrt1(x, cur+1, end)
	}

	return -1
}
