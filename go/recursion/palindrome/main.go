package main

func IsPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}

	if s[0] != s[len(s)-1] {
		return false
	}

	return IsPalindrome(s[1 : len(s)-1])
}

func main() {
}
