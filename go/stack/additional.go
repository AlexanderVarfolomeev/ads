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
		switch val {
		case "=":
			return calculateEqual(&s2)
		case "+":
			err := calculatePlus(&s2)
			if err != nil {
				return 0, errors.New("incorrect expression")
			}
		case "*":
			err := calculateMultiply(&s2)
			if err != nil {
				return 0, errors.New("incorrect expression")
			}
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

func stringToStack(expression string) Stack[string] {
	s1 := Stack[string]{}
	str := strings.Split(expression, " ")

	for i := len(str) - 1; i >= 0; i-- {
		s1.Push(str[i])
	}

	return s1
}

func calculateEqual(s2 *Stack[int]) (int, error) {
	res, err := s2.Pop()
	if err != nil {
		return 0, errors.New("incorrect expression")
	}

	if s2.Size() != 0 {
		return 0, errors.New("incorrect expression")
	}

	return res, nil
}

func calculatePlus(s2 *Stack[int]) error {
	val1, err := s2.Pop()
	if err != nil {
		return errors.New("incorrect expression")
	}
	val2, err := s2.Pop()
	if err != nil {
		return errors.New("incorrect expression")
	}

	s2.Push(val1 + val2)
	return nil
}

func calculateMultiply(s2 *Stack[int]) error {
	val1, err := s2.Pop()
	if err != nil {
		return errors.New("incorrect expression")
	}
	val2, err := s2.Pop()
	if err != nil {
		return errors.New("incorrect expression")
	}

	s2.Push(val1 * val2)
	return nil
}
