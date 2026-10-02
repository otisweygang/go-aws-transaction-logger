package transaction

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name        string
		accountRef  string
		amountPence int64
		wantErr     error
	}{
		{name: "validPositive", accountRef: "AAA111", amountPence: 500, wantErr: nil},
		{name: "validNegative", accountRef: "AAA111", amountPence: -500, wantErr: nil},
		{name: "atMax", accountRef: "AAA111", amountPence: maxAmountPence, wantErr: nil},
		{name: "atMin", accountRef: "AAA111", amountPence: -maxAmountPence, wantErr: nil},
		{name: "emptyAccountRef", accountRef: "", amountPence: 500, wantErr: ErrEmptyAccountRef},
		{name: "zeroAmount", accountRef: "AAA111", amountPence: 0, wantErr: ErrZeroAmount},
		{name: "overMax", accountRef: "AAA111", amountPence: maxAmountPence + 1, wantErr: ErrAmountOutOfRange},
		{name: "underMin", accountRef: "AAA111", amountPence: -maxAmountPence - 1, wantErr: ErrAmountOutOfRange},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tx := NewTransaction(tc.accountRef, tc.amountPence)
			err := tx.Validate()
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("Validate(%q, %d) err = %v, want %v", tc.accountRef, tc.amountPence, err, tc.wantErr)
			}
		})
	}
}

func TestFormatPence(t *testing.T) {
	tests := []struct {
		name string
		in   int64
		want string
	}{
		{name: "zero", in: 0, want: "£0.00"},
		{name: "5pence", in: 5, want: "£0.05"},
		{name: "100pence", in: 100, want: "£1.00"},
		{name: "3456pence", in: 3456, want: "£34.56"},
		{name: "-130pence", in: -130, want: "-£1.30"},
		{name: "-5pence", in: -5, want: "-£0.05"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := formatPence(tc.in)
			if got != tc.want {
				t.Errorf("formatPence(%d) got %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
