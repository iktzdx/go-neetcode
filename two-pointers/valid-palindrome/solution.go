package validpalindrome

import (
	"strings"
)

func IsPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isAlphaNum(s[left]) {
			left++
		}

		for right > left && !isAlphaNum(s[right]) {
			right--
		}

		if !strings.EqualFold(string(s[left]), string(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphaNum(ch byte) bool {
	return (ch >= '0' && ch <= '9') || (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z')
}
