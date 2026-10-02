package transaction

import (
	"errors"
	"fmt"
	"time"
)

const maxAmountPence int64 = 100_000_000 // £1,000,000.00

var (
	ErrEmptyAccountRef  = errors.New("account reference is empty")
	ErrZeroAmount       = errors.New("amount is zero")
	ErrAmountOutOfRange = errors.New("amount outside allowed range")
)

type Transaction struct {
	AccountRef  string
	AmountPence int64
	CreatedAt   time.Time
}

func NewTransaction(accountRef string, amountPence int64) Transaction {
	return Transaction{AccountRef: accountRef, AmountPence: amountPence}
}

func (t Transaction) Validate() error {
	if t.AccountRef == "" {
		return ErrEmptyAccountRef
	}
	if t.AmountPence == 0 {
		return ErrZeroAmount
	}
	if t.AmountPence > maxAmountPence || t.AmountPence < -maxAmountPence {
		return fmt.Errorf("%d pence: %w", t.AmountPence, ErrAmountOutOfRange)
	}
	return nil
}

func formatPence(input int64) string {
	sign := ""
	if input < 0 {
		sign = "-"
		input = -input
	}
	pounds := input / 100
	pence := input % 100

	return fmt.Sprintf("%s£%d.%02d", sign, pounds, pence)
}
