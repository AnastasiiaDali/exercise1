package fibonacci

import (
	"errors"
	"math/big"
)

type Fibonacci struct {
}

func NewFibonacci() *Fibonacci {
	return &Fibonacci{}
}

func (c Fibonacci) FibonacciNumber(number int) (string, error) {
	if number == 0 {
		return "0", nil
	}
	if number == 1 {
		return "1", nil
	}
	if number < 0 {
		return "0", errors.New("Input cannot be a negative number")
	}
	if number >= 93 {
		return "0", errors.New("integer overflow - please choose number smaller than 93")
	}

	a := big.NewInt(0)
	b := big.NewInt(1)

	for i := 2; i <= number; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	return b.String(), nil
}
