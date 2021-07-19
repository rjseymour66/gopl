package main

import (
	"fmt"
)

var sample string = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

func main() {
/*
	fmt.Println(sample)

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}

	fmt.Println("\n+q formatting:\n")

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%+q ", sample[i])
	}

	fmt.Println("\n\n")

	fmt.Println()
	fmt.Printf("%x\n", sample)
	fmt.Printf("% x\n", sample)
	fmt.Printf("%q\n", sample)
	fmt.Printf("%+q\n", sample)

	fmt.Println("* * * * * * * * * ")
	bSample := []byte(sample)

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}

	fmt.Println()
	fmt.Printf("%x\n", bSample)
	fmt.Printf("% x\n", bSample)
	fmt.Printf("%q\n", bSample)
	fmt.Printf("%+q\n", bSample)
*/

	/* UTF-8 and string literals */

	const placeOfInterest = `⌘"`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")

	const helloWorld = "Hello, World!"
	for i, runeValue := range helloWorld {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
	}

	fmt.Println()

	const tester = "Hello, \u3003"
	for i, runeValue := range tester {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
	}
}













































