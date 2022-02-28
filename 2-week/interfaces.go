package main

import (
	"errors"
	"fmt"
)

var ErrNoSufficientBalance = errors.New("Yeterli bakiyeniz bulunmamaktadır")

func main() {

	b := BankAccount{"12345", 100}

	err := b.Withdraw(150)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	PrintDepositableBankAccountsBalance(b)
}

func PrintDepositableBankAccountsBalance(a Depositable) {

}

type Depositable interface {
	Deposit(a int)
	Print()
}

type BankAccount struct {
	AccountNumber string
	Balance       int
}

func (b *BankAccount) Withdraw(a int) error {
	if b.Balance < a {
		return ErrNoSufficientBalance
	}

	b.Balance -= a
	return nil
}

func (b BankAccount) Deposit(a int) {
	b.Balance += a
}
func (b BankAccount) Print() {
	fmt.Printf("Hesabınızdan para çekim işlemi başarıyla tamamlandı. Kalan bakiye : %d\n", b.Balance)
}
