package calc

import "errors"

func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Denominator cannot be zero.")
	}
	return (a / b), nil
}
