package longestrepeatingcharacterreplacement_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/sliding-window/longest-repeating-character-replacement"
)

func Test_CharacterReplacement(t *testing.T) {
	tests := map[string]struct {
		input string
		ops   int
		want  int
	}{
		"multiple replacements": {
			input: "ABAB",
			ops:   2,
			want:  4, // "AAAA" or "BBBB"
		},
		"only one replacement": {
			input: "CCDCDDC",
			ops:   1,
			want:  4, // "CCCCDDC" or "CCDDDDC"
		},
		"no replacements - has duplicates": {
			input: "BAAABABB",
			ops:   0,
			want:  3, // "AAA" from "BAAABABB"
		},
		"no replacements - no duplicates": {
			input: "AU",
			ops:   0,
			want:  1, // "A" from "AU" or "U" from "AU"
		},
		"single character": {
			input: "X",
			ops:   1,
			want:  1, // "X"
		},
		"k equals input length": {
			input: "ABCDEFGHKJ",
			ops:   10,
			want:  10, // all characters will be the same
		},
		"long string": {
			input: "EOEMQLLQTRQDDCOERARHGAAARRBKCCMFTDAQOLOKARBIJBISTGNKBQGKKTALSQNFSABASNOPBMMGDIOETPTDICRBOMBAAHINTFLH",
			ops:   7,
			want:  11,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.CharacterReplacement(test.input, test.ops)
			if test.want != got {
				t.Fatalf("CharacterReplacement(%s, %d): expected = %d, got = %d", test.input, test.ops, test.want, got)
			}
		})
	}
}
