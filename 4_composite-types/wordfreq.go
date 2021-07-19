/*
wordfreq reports the frequency of each word in an input text file.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)	// counts of each word

	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		counts[scanner.Text()]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	fmt.Printf("word\tcount\n")
	for w, c := range counts {
		fmt.Printf("%s\t%d\n", w, c)
	}
}
