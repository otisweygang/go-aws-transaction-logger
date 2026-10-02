package transaction

import "fmt"

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
