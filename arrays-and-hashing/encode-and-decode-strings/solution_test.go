package encodeanddecodestrings_test

import (
	"reflect"
	"testing"

	solution "github.com/iktzdx/go-neetcode/arrays-and-hashing/encode-and-decode-strings"
)

func Test_EncodeString(t *testing.T) {
	tests := map[string]struct {
		input []string
		want  string
	}{
		"words only": {
			input: []string{"hello", "world"},
			want:  "5#hello5#world",
		},
		"with punctuation": {
			input: []string{"we", "say", ":", "brrr"},
			want:  "2#we3#say1#:4#brrr",
		},
		"contains delimiter": {
			input: []string{"inp#t"},
			want:  "5#inp#t",
		},
		"contains delimiter and length": {
			input: []string{"5#input"},
			want:  "7#5#input",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.Encode(test.input)
			if test.want != got {
				t.Fatalf("Encode(%v): expected = %s, got = %s", test.input, test.want, got)
			}
		})
	}
}

func Test_DecodeString(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []string
	}{
		"words only": {
			input: "5#hello5#world",
			want:  []string{"hello", "world"},
		},
		"with punctuation": {
			input: "2#we3#say1#:4#brrr",
			want:  []string{"we", "say", ":", "brrr"},
		},
		"contains delimiter": {
			input: "5#inp#t",
			want:  []string{"inp#t"},
		},
		"contains delimiter and length": {
			input: "7#5#input",
			want:  []string{"5#input"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.Decode(test.input)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("Decode(%s): expected = %v, got = %v", test.input, test.want, got)
			}
		})
	}
}

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
		"contains delimiter and length": {
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
