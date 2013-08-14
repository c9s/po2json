package main

import (
	"fmt"
	"os"
	"po2json/pofile"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: po2json [file]")
		os.Exit(0)
	}

	filename := os.Args[1]

	if _, err := os.Stat(filename); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dict, err := pofile.ParseFile(filename)
	if err != nil {
		fmt.Println("PO File Parsing Error", err)
		os.Exit(1)
	}

	fmt.Println(dict)

	_ = err
	_ = dict
}
