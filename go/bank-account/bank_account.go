package account

import (
	"sync"
)

// Account struct to save bank transactions
type Account struct {
	mux    sync.Mutex
	Amount int64
	Closed bool
}

// Balance function to return the amount value
func (a *Account) Balance() (int64, bool) {
	if a.Amount >= 0 && !a.Closed {
		return a.Amount, true
	}

	return a.Amount, false
}

// Deposit function to add more money in account
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if a.Closed {
		return a.Amount, false
	}

	if amount < 0 && a.Amount == 0 {
		return 0, false
	}

	a.Amount = a.Amount + amount

	return a.Amount, true
}

// Close function to close account and retrieve the amount value
func (a *Account) Close() (int64, bool) {
	var payout int64

	a.mux.Lock()
	defer a.mux.Unlock()

	if !a.Closed {
		payout = a.Amount
		a.Amount = 0
		a.Closed = true

		return payout, true
	}

	return payout, false
}

// Open function to open a new account
func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	var account Account = Account{
		Amount: amount,
		Closed: false,
	}

	return &account
}
