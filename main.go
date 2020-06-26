package main

import (
	"fmt"
	"github.com/jieunjeon/Web-Scrapper-With-GoLang_Echo/accounts"
)

func main() {
	account := accounts.NewAccount("jieun")
	fmt.Println(account)
}