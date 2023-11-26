package main

func IsPalindrome(s string) bool {
	return isPalindrome(s, 0, len(s)-1)
}

func isPalindrome(s string, start, end int) bool {
	if start >= end {
		return true
	}

	if s[start] != s[end] {
		return false
	}

	return isPalindrome(s, start+1, end-1)
}

func main() {
}
