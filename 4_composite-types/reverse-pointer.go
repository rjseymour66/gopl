package main

import (
	"fmt"
)

func main() {

	x := []int{1, 2, 3, 4, 5}
	reverse(x)
	fmt.Println(x)

	y := [...]int{1, 2, 3, 4, 5}
	revPtr(&y)
	fmt.Println(y)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func revPtr(s *[5]int) {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
