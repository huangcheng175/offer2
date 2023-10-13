package offer2

import (
	"fmt"
	"testing"
)

func TestAddBinary(t *testing.T) {
	fmt.Println(addBinary("100", "11"))
	fmt.Println(addBinary("101", "11"))
	fmt.Println(addBinary("111", "11"))
}

func addBinary(a, b string) string {
	la := len(a) - 1
	lb := len(b) - 1

	var (
		va, vb, vSum, flag uint8
	)
	vs := make([]byte, 0)
	for la >= 0 || lb >= 0 || flag > 0 {

		fmt.Println(la, lb)
		va = 0
		vb = 0
		if la >= 0 {
			va = a[la] - '0'
		}
		if lb >= 0 {
			vb = b[lb] - '0'
		}

		vSum += va + vb + flag
		fmt.Println("aa", la, lb, flag, a, b, vSum, vs, vSum%2)

		vs = append([]byte{vSum%2 + '0'}, vs...)
		if vSum > 1 {
			flag = 1
		} else {
			flag = 0
		}

		la--
		lb--
	}

	return string(vs)
}
