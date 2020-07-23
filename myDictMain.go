package main

import (
	"fmt"

	"github.com/jieunjeon/Web-Scrapper-With-GoLang_Echo/myDict"
)

func main() {
	dictionary := myDict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
