package validpalindrome

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	var sb strings.Builder
	for _, ch := range s {
		if regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(string(ch)) {
			sb.WriteRune(ch)
		}
	}

	cleaned := strings.ToLower(sb.String())

	sb.Reset()

	for i := len(s) - 1; i >= 0; i-- {
		if regexp.MustCompile(`^[a-zA-Z0-9]*$`).Match([]byte{s[i]}) {
			sb.WriteByte(s[i])
		}
	}

	return cleaned == strings.ToLower(sb.String())
}
