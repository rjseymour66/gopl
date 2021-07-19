package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	rtype := make(map[string]int)		// count of rune type
	invalid := 0				// count of invalide UTF-8 chars

	// initialize rune types
	rtype["letter"] = 0
	rtype["digit"]  = 0
	rtype["number"] = 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "new-charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			rtype["letter"]++
		}
		if unicode.IsDigit(r) {
			rtype["digit"]++
		}
		if unicode.IsNumber(r) {
			rtype["number"]++
		}
	}
	fmt.Print("type\tcount\n")

	for t, c := range rtype {
		fmt.Printf("%s\t%5d\n", t, c)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
