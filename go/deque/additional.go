package main

import "strings"

func isPalindrome(s string) bool {
	d := Deque[string]{}
	for _, c := range strings.ToLower(s) {
		if string(c) != " " {
			d.AddTail(string(c))
		}
	}

	for d.Size() > 1 {

		c1, _ := d.RemoveFront()
		c2, _ := d.RemoveTail()

		if c1 != c2 {
			return false
		}
	}

	return true
}
