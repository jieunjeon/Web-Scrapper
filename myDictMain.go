package main

import (
	"fmt"

	"github.com/jieunjeon/Web-Scrapper-With-GoLang_Echo/myDict"
)

func main() {
	dictionary := myDict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println("found", word, "definition:", hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}
