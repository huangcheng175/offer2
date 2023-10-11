package offer2

import (
	"fmt"
	"testing"
)

// ca
// dgcaf

func checkInclusion(s1, s2 string) bool {
	cnt := [26]int{}
	for i := 0; i < len(s1); i++ {
		cnt[s1[i]-'a']++
		cnt[s2[i]-'a']--
	}

	if isAllZero(cnt) {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		cnt[s2[i]-'a']--
		cnt[s2[i-len(s1)]-'a']++
		if isAllZero(cnt) {
			return true
		}
	}

	return false

}

func isAllZero(arr [26]int) bool {
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}

	return true
}

func TestCheckInclusion(t *testing.T) {
	fmt.Println(checkInclusion("ca", "dgcaf"))
	fmt.Println(checkInclusion("cad", "dgcaf"))
}
