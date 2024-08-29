package evaluatereversepolishnotation

import "strconv"

func EvalRPN(tokens []string) int {
	result, _ := strconv.Atoi(tokens[0])
	stack := make([]string, 0)

	for _, token := range tokens {
		if token != "+" && token != "-" && token != "*" && token != "/" {
			stack = append(stack, token)

			continue
		}

		result = calc(stack[len(stack)-2], stack[len(stack)-1], token)
		stack = stack[:len(stack)-1]
		stack[len(stack)-1] = strconv.Itoa(result)
	}

	return result
}

func calc(a, b, op string) int {
	aInt, _ := strconv.Atoi(a)
	bInt, _ := strconv.Atoi(b)

	switch op {
	case "+":
		return aInt + bInt
	case "-":
		return aInt - bInt
	case "*":
		return aInt * bInt
	case "/":
		return aInt / bInt
	}

	return 0 // for compiler
}
