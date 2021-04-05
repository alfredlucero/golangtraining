package pointers

import (
	"errors"
	"fmt"
)

// Pointers let us point to some values and then let us change them.

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	// If a symbol starts with a lowercase, then it is private outside the package it's defined in.
	balance Bitcoin
}

// In Go, when you call a function or a method, the arguments are copied.
func (w *Wallet) Deposit(amount Bitcoin) {
	// (*w).balance works too, but struct pointers in general are automatically dereferenced.
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Create a global variable to the package so we don't duplicate error messages everywhere in tests/main code.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
