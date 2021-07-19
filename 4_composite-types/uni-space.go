package main

import (
	"fmt"
	"unicode"
)

func main() {
	// transform adjacent Unicode spaces in UTF-8 encoded byte
	// slice into single ASCII space

	tester := []byte("This is a byte slice")

	answer := adjSpace(tester)

	fmt.Printf("start:\t%s\nend:\t%s\n", tester, answer)


}

func adjSpace(value []byte) []byte {
	for i, j := 0, 1; j <= len(value); i, j = i+1, j+1 {
		if unicode.IsSpace(i) && unicode.IsSpace(j) {
			copy(value[i:], value[j:])
			value = value[:len(value)-1]
		}
	}

	return value
}
