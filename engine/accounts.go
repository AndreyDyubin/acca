package engine

import (
	"time"

	"github.com/pkg/errors"
)

//go:generate reform

//reform:acca.currencies
type Currency struct {
	CurrencyID int64   `reform:"curr_id,pk"`
	ClientID   *int64  `reform:"client_id"`
	Key        string  `reform:"key"`
	Meta       *[]byte `reform:"meta"`
}

//reform:acca.accounts
type Account struct {
	AccountID       int64     `reform:"acc_id,pk"`
	ClientID        *int64    `reform:"client_id"`
	CurrencyID      int64     `reform:"curr_id"`
	Key             string    `reform:"key"`
	Balance         int64     `reform:"balance"`
	BalanceAccepted int64     `reform:"balance_accepted"`
	Meta            *[]byte   `reform:"meta"`
	LastTxID        *int64    `reform:"last_tx_id"`
	UpdatedAt       time.Time `reform:"updated_at"`
	CreatedAt       time.Time `reform:"created_at"`
}

func (a *Account) BeforeInsert() error {
	a.UpdatedAt = time.Now()
	a.CreatedAt = time.Now()
	if a.Balance != 0 {
		return errors.New("new account can't be with not zero balance")
	}
	return nil
}

func (a *Account) BeforeUpdate() error {
	a.UpdatedAt = time.Now()
	return nil
}
