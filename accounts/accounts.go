package accounts

import "errors"

type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't widthdraw")

// function: NewAccount cretes Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// method, reciever a. Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	// pointer reciever
	// Go copies and pass the copied value. So if we want to edit the original account reciever ( not to copied account receiver), need to put *.
	a.balance += amount
}

// Balance of account
func (a Account) Balance() int {
	return a.balance
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
