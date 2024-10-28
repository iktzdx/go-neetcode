package lettercombinationsofaphonenumber

var letters = map[string]string{ //nolint:gochecknoglobals
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

type backtracking func(pos int, curr string)

func LetterCombinations(digits string) []string {
	result := make([]string, 0)

	var backtrack backtracking

	backtrack = func(pos int, curr string) {
		if len(curr) == len(digits) {
			result = append(result, curr)

			return
		}

		for _, l := range letters[string(digits[pos])] {
			backtrack(pos+1, curr+string(l))
		}
	}

	if digits != "" {
		backtrack(0, "")
	}

	return result
}
