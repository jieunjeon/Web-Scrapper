package main

import (
	"fmt"

	"github.com/jieunjeon/Web-Scrapper-With-GoLang_Echo/accounts"
	"github.com/jieunjeon/Web-Scrapper-With-GoLang_Echo/myDict"
)

func main2() {
	account := accounts.NewAccount("jieun")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())

	runDict()
}

func runDict() {
	dictionary := myDict.Dictionary{}
	dictionary["hello"] = "hello"
	fmt.Println(dictionary)
}
