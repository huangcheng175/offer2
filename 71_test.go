package offer2

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRandByWeight(t *testing.T) {
	x := []int{1, 2, 3, 4}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		t.Log(randByWeight(x))
	}

}

func randByWeight(x []int) int {
	if len(x) == 0 {
		return 0
	}
	vx := generateSumArr(x)
	sum := vx[len(vx)-1]
	fmt.Println(vx)
	rd := rand.Intn(sum)
	//rd = 9
	fmt.Print(rd, ".....")

	return find(rd, vx)
}

func find(rd int, vx []int) int {
	l := 0
	r := len(vx) - 1

	for l <= r {
		mid := (l + r) / 2
		if vx[mid] == rd { // 已经相等了，不需要重新考虑了
			return mid
		}
		if vx[mid] < rd && (mid == len(vx)-1 || vx[mid+1] > rd) {
			return mid + 1
		}

		if vx[mid] > rd && (mid == 0 || vx[mid-1] < rd) {
			return mid
		}

		if vx[mid] > rd {
			r = mid - 1
		}
		if vx[mid] < rd {
			l = mid + 1
		}
		find(rd, vx[l:r])
	}

	return -1
}

func generateSumArr(x []int) []int {
	vx := make([]int, len(x))
	sum := 0
	for i, v := range x {
		sum += v
		vx[i] = sum
	}
	return vx
}
