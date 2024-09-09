package longestsubstringwithoutrepeatingcharacters_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/sliding-window/longest-substring-without-repeating-characters"
)

func Test_LengthOfLongestSubstring(t *testing.T) {
	tests := map[string]struct {
		input string
		want  int
	}{
		"all characters are the same": {
			input: "wwwwwwwwwwwwwwwww",
			want:  1,
		},
		"empty string": {
			input: "",
			want:  0,
		},
		"single character": {
			input: "s",
			want:  1,
		},
		"with duplicate substrings": {
			input: "abcabcbb",
			want:  3,
		},
		"with duplicate characters": {
			input: "pwwkew",
			want:  3,
		},
		"with digits": {
			input: "qqw39tyy",
			want:  6,
		},
		"with symbols": {
			input: "21!1t$4qq",
			want:  6,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.LengthOfLongestSubstring(test.input)
			if test.want != got {
				t.Fatalf("LengthOfLongestSubstring(%s): expected = %d, got = %d", test.input, test.want, got)
			}
		})
	}
}
