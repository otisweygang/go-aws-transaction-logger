package transaction

import "testing"

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
