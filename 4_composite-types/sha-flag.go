package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)


func main() {

	var hash string

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stdout, "No argument provided: \"prog <arg> -flags\"\n")
		os.Exit(1)
	}

	flag.StringVar(&hash, "hash", "256", "provide hash digest value: 256, 384, 512")
	flag.Parse()

	value := os.Args[1]


	switch hash {
	case "256":
		fmt.Printf("%x", sha256.Sum256([]byte(value)))
	case "384":
		fmt.Printf("%x", sha512.Sum384([]byte(value)))
	case "512":
		fmt.Printf("%x", sha512.Sum512([]byte(value)))
	}
	fmt.Printf("\nhash flag value:\t%s\n", hash)
}
