package main

import (
	"errors"
	"strconv"
	"strings"
)

func CalculatePostfix(expression string) (float64, error) {
	s1 := stringToStack(expression)
	s2 := Stack[float64]{}

	for s1.Size() != 0 {
		val, _ := s1.Pop()

		isBinaryOp := isBinaryOperation(val)
		var val1, val2 float64

		if isBinaryOp && s2.Size() < 2 {
			return 0, errors.New("incorrect expression")
		}

		if isBinaryOp {
			val1, _ = s2.Pop()
			val2, _ = s2.Pop()
		}

		switch val {
		case "=":
			res, err := s2.Pop()
			if err != nil || s2.Size() != 0 {
				return 0, errors.New("incorrect expression")
			}
			return res, nil
		case "+":
			s2.Push(val1 + val2)
		case "*":
			s2.Push(val1 * val2)
		case "/":
			s2.Push(val2 / val1)
		case "-":
			s2.Push(val2 - val1)
		default:
			num, err := strconv.ParseFloat(val, 32)
			if err != nil {
				return num, errors.New("incorrect expression")
			}
			s2.Push(num)
		}
	}

	return 0, errors.New("incorrect expression")
}

func isBinaryOperation(op string) bool {
	return op == "+" || op == "-" || op == "*" || op == "/"
}

func stringToStack(expression string) Stack[string] {
	s1 := Stack[string]{}
	str := strings.Split(expression, " ")

	for i := len(str) - 1; i >= 0; i-- {
		s1.Push(str[i])
	}

	return s1
}
