package evaluatereversepolishnotation

import "strconv"

func EvalRPN(tokens []string) int {
	if len(tokens) == 1 {
		val, _ := strconv.Atoi(tokens[0])

		return val
	}

	stack := []int{}

	for i := 0; i < len(tokens); i++ {
		for tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/" {
			val, _ := strconv.Atoi(tokens[i])
			stack = append(stack, val)

			i++
		}

		a := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		b := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		stack = append(stack, calc(b, a, tokens[i]))
	}

	return stack[len(stack)-1]
}

func calc(a, b int, op string) int {
	if op == "+" {
		return a + b
	}
	if op == "-" {
		return a - b
	}
	if op == "*" {
		return a * b
	}
	if op == "/" {
		return a / b
	}

	return -1 // for compiler
}
