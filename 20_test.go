package offer2

import (
	"fmt"
	"testing"
)

func countSubStrings(s string) int {
	if len(s) == 0 {
		return 0
	}

	count := 0
	for i := 0; i < len(s); i++ {
		count += countPalindrome(s, i, i)
		count += countPalindrome(s, i, i+1)
	}

	return count
}

func countPalindrome(s string, start, end int) int {
	var count int = 0

	for start >= 0 && end < len(s) && s[start] == s[end] {
		count++
		start--
		end++
	}

	return count
}

func TestCountSubStrings(t *testing.T) {
	fmt.Println(countSubStrings("aaa"))
}
