package main

import (
	"fmt"

	"github.com/jieunjeon/Web-Scrapper-With-GoLang_Echo/myDict"
)

func main() {
	dictionary := myDict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
