package offer2

import (
	"fmt"
	"testing"
)

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		ch1 := s[i]
		ch2 := s[j]

		if ch2 != ch1 {
			return false
		}
		i++
		j--
	}

	return true
}

func TestIsPalindrome(t *testing.T) {
	fmt.Println(isPalindrome("aba"))
	fmt.Println(isPalindrome("abba"))
	fmt.Println(isPalindrome("a"))
	fmt.Println(isPalindrome(""))
	fmt.Println(isPalindrome("abc"))

}
