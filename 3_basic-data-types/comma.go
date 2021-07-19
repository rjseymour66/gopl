package main

import(
	"bytes"
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("comma:\t", comma("12345"))
	fmt.Println("nrComma:", nrComma("12345"))

	x := 12345
	y := fmt.Sprintf("%d", x)
	fmt.Println(y)

	s, err := strconv.Atoi("654")
	if err != nil {
		fmt.Println("This didn't work")
	}
	fmt.Printf("%d\n", s)
}

// comma inserts commas in a non-negative decimal integer string
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func nrComma(s string) string {
	var buf bytes.Buffer

	for _, v := range s {
		buf.WriteRune(v)
	}

	return buf.String()
}
