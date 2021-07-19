package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	s, sep := "", ""
	first := time.Now()

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(time.Since(first))

	sec := time.Now()

	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(time.Since(sec))
}
