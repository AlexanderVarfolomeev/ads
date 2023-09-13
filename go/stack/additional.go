package main

const (
	openBracket  = "("
	closeBracket = ")"
)

func main() {
	println(CorrectnessOfBrackets(")(()())("))
}
func CorrectnessOfBrackets(str string) bool {
	s := Stack[string]{}
	for _, sym := range str {
		if val, err := s.Peek(); err == nil && val == openBracket && string(sym) == closeBracket {
			_, _ = s.Pop()
		} else {
			s.Push(string(sym))
		}
	}

	return s.Size() == 0
}
