package encodeanddecodestrings

import (
	"fmt"
	"strconv"
)

const delimiter rune = '#'

func Encode(strs []string) string {
	var res string

	for _, str := range strs {
		res += fmt.Sprintf("%d%s%s", len(str), string(delimiter), str)
	}

	return res
}

func Decode(str string) []string {
	res := make([]string, 0)

	start := 0
	rs := []rune(str)

	for start < len(rs) {
		end := start
		for rs[end] != delimiter {
			end++
		}

		length, err := strconv.Atoi(string(rs[start:end]))
		if err != nil {
			return res
		}

		res = append(res, string(rs[end+1:end+1+length]))

		start = end + 1 + length
	}

	return res
}
