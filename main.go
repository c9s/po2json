package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: po2json [file]")
	}

	filename := os.Args[1]

	if _, err := os.Stat(filename); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dict, err := ParsePOFile(filename)
	if err != nil {
		fmt.Println("PO File Parsing Error", err)
		os.Exit(1)
	}

	jsonBytes, err := json.MarshalIndent(dict, "", "  ")
	if err != nil {
		fmt.Println("JSON Encode Error: ", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonBytes))

	_ = err
	_ = dict
}
