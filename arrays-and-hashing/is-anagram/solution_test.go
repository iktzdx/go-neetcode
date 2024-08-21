package validanagram_test

import (
	"testing"

	validanagram "github.com/iktzdx/go-neetcode/arrays-and-hashing/is-anagram"
)

func Test_IsAnagram(t *testing.T) {
	tests := map[string]struct {
		input       string
		compareWith string
		want        bool
	}{
		"is a valid anagram": {
			input:       "anagram",
			compareWith: "nagaram",
			want:        true,
		},
		"different set of characters": {
			input:       "rat",
			compareWith: "car",
			want:        false,
		},
		"different length of strings": {
			input:       "hello",
			compareWith: "ollehh",
			want:        false,
		},
		"the same word": {
			input:       "neetcode",
			compareWith: "neetcode",
			want:        true,
		},
		"contains unicode": {
			input:       "вижу зверей",
			compareWith: "живу резвей",
			want:        true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := validanagram.IsAnagram(test.input, test.compareWith)
			if test.want != got {
				t.Fatalf("IsAnagram(%s, %s), expect = %v, got = %v\n", test.input, test.compareWith, test.want, got)
			}
		})
	}
}
