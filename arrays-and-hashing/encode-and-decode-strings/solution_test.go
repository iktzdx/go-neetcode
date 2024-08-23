package encodeanddecodestrings_test

import (
	"reflect"
	"testing"

	solution "github.com/iktzdx/go-neetcode/arrays-and-hashing/encode-and-decode-strings"
)

func Test_EncodeDecodeString(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  []string
	}{
		"words only": {
			input: []string{"hello", "darkness", "my", "old", "friend"},
			want:  []string{"hello", "darkness", "my", "old", "friend"},
		},
		"with punctuation": {
			input: []string{"we", "say", ":", "hello", ",", "world", "!"},
			want:  []string{"we", "say", ":", "hello", ",", "world", "!"},
		},
		"empty string": {
			input: []string{""},
			want:  []string{""},
		},
		"contains delimiter": {
			input: []string{"hello", "d#l#m#t#r"},
			want:  []string{"hello", "d#l#m#t#r"},
		},
		"contains word's length and delimiter": {
			input: []string{"5hel#o", "w4#rld"},
			want:  []string{"5hel#o", "w4#rld"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.Decode(solution.Encode(test.input))
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("Decode(Encode(%v)): expected = %v, got = %v", test.input, test.want, got)
			}
		})
	}
}
