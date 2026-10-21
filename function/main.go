package main

import (
	"errors"
	"fmt"
)


type Account struct {
	Owner 	string
	Balance 	float64
}

func Deposit(account *Account, amount float64) (float64, error) {
    if amount <= 0 {
		return 0, errors.New("amount more then 0")
	}
	account.Balance += amount
	return account.Balance, nil
}

func Withdraw(account *Account, amount float64) error {
    if account.Balance < amount  {
		return errors.New("Not enough")
	}

	account.Balance -= amount
	return nil
}


func main() {
	// Initial Balance
	personOne := Account{
		Owner: "Putthakun",
		Balance: 1000,
	}

	balance, err := Deposit(&personOne, 500)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Balance ", balance)

}
