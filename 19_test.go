package offer2

import (
	"fmt"
	"testing"
)

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			break
		}
		i++
		j--
	}

	return i >= j || isPalindromeExtra(s, i+1, j) || isPalindromeExtra(s, i, j-1)
}

func isPalindromeExtra(s string, start, end int) bool {
	i, j := start, end

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func TestValidPalindrome(t *testing.T) {
	fmt.Println(validPalindrome("abca"))
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("abcd"))
	fmt.Println(validPalindrome("abcca"))
}
