package pointersanderrors

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// using a pointer here as values are copied by default
// a copy would be expensive and would not update
// the struct this method is used on
func (w *Wallet) Deposit(deposit Bitcoin) {
	w.balance += deposit
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficentFunds = errors.New("cannot withdraw insufficient funds")

func (w *Wallet) Withdraw(withdrawal Bitcoin) error {
	if withdrawal > w.balance {
		return ErrInsufficentFunds
	}

	w.balance -= withdrawal
	return nil
}
