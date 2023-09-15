package main

import (
	"errors"
	"strconv"
	"strings"
)

func CalculatePostfix(expression string) (int, error) {
	s1 := stringToStack(expression)
	s2 := Stack[int]{}

	for s1.Size() != 0 {
		val, _ := s1.Pop()

		if isBinaryOperation(val) && s2.Size() < 2 {
			return 0, errors.New("incorrect expression")
		}

		switch val {
		case "=":
			res, err := s2.Pop()
			if err != nil || s2.Size() != 0 {
				return 0, errors.New("incorrect expression")
			}

			return res, nil
		case "+":
			val1, _ := s2.Pop()
			val2, _ := s2.Pop()

			s2.Push(val1 + val2)
		case "*":
			val1, _ := s2.Pop()
			val2, _ := s2.Pop()

			s2.Push(val1 * val2)
		default:
			num, err := strconv.Atoi(val)
			if err != nil {
				return num, errors.New("incorrect expression")
			}
			s2.Push(num)
		}
	}

	return 0, errors.New("incorrect expression")
}

func isBinaryOperation(op string) bool {
	return op == "+" || op == "*"
}

func stringToStack(expression string) Stack[string] {
	s1 := Stack[string]{}
	str := strings.Split(expression, " ")

	for i := len(str) - 1; i >= 0; i-- {
		s1.Push(str[i])
	}

	return s1
}
