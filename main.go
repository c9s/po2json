package main

import "os"
import "fmt"

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: po2json [file]")
	}

	filename := os.Args[1]

	if _, err := os.Stat(filename); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err := ParsePOFile(filename)
	_ = err
}
