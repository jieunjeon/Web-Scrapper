package main

import (
	"fmt"
	"github.com/jieunjeon/Web-Scrapper-With-GoLang_Echo/accounts"
)

func main() {
	account := accounts.NewAccount("jieun")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
}