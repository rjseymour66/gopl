package main

import (
	"fmt"
)

func main() {

	test := []string{"one", "two", "two", "three", "three", "four", "four"}
	fmt.Println(test)
	fmt.Println(dups(test))
}

func dups(s []string) []string {

	for i, j := 0, 1; j <= len(s); i, j = i+1, j+1 {
		if s[i] == s[j] {
			copy(s[i:], s[j:])
			s = s[:len(s)-1]
		}
	}
	return s
}
