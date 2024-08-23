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
