package main

import (
	"flag"
	"fmt"
	"github.com/c9s/po2json/po"
	"io/ioutil"
	"os"
	"path"
)

var domainOpt = flag.String("domain", "", "domain")
var localeOpt = flag.String("locale", "locale", "locale directory")
var jsonDirOpt = flag.String("json-dir", "public/locale", "json output directory")

func FileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		return true
	}
	return false
}

func main() {
	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println("Usage: ")
		fmt.Println("   po2json [file] > app.json")
		fmt.Println("   po2json --domain app --locale path/to/locale_dir --json-dir public/locale")
		os.Exit(0)
	}

	// if domain is specified, we should get the po files from locale directory.
	if *domainOpt != "" {
		var domain = *domainOpt
		var localeDir = *localeOpt
		var langs = []string{}
		fileInfos, err := ioutil.ReadDir(*localeOpt)

		// get languages
		for _, fi := range fileInfos {
			if fi.IsDir() {
				fmt.Println(fi.Name())
				langs = append(langs, fi.Name())
			}
		}

		for _, lang := range langs {
			var poFile = path.Join(localeDir, lang, "LC_MESSAGES", domain) + ".po"
			_ = poFile
		}

		_ = err
	} else {
		filename := os.Args[1]
		if _, err := os.Stat(filename); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dict, err := po.ParseFile(filename)
		if err != nil {
			fmt.Println("PO File Parsing Error", err)
			os.Exit(1)
		}
		fmt.Println(dict)
	}
}
