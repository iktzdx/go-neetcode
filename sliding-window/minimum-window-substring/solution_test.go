package minimumwindowsubstring_test

import (
	"testing"

	solution "github.com/iktzdx/go-neetcode/sliding-window/minimum-window-substring"
)

func Test_MinWindow(t *testing.T) {
	tests := map[string]struct {
		s    string
		t    string
		want string
	}{
		"upper case letters": {
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		"lower case letter": {
			s:    "ouzodyxazv",
			t:    "xyz",
			want: "yxaz",
		},
		"single character": {
			s:    "a",
			t:    "a",
			want: "a",
		},
		"t is longer than s": {
			s:    "a",
			t:    "aa",
			want: "",
		},
		"s is the same as t": {
			s:    "aa",
			t:    "aa",
			want: "aa",
		},
		"duplicate characters": {
			s:    "bbaa",
			t:    "aba",
			want: "baa",
		},
		"s is empty": {
			s:    "",
			t:    "xyz",
			want: "",
		},
		"t is empty": {
			s:    "zyx",
			t:    "",
			want: "",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := solution.MinWindow(test.s, test.t)
			if test.want != got {
				t.Fatalf("MinWindow(%s, %s): expected = %s, got = %s", test.s, test.t, test.want, got)
			}
		})
	}
}
