package calculator

import (
	"log"
	"math"
)

type Calculator struct {
}

func New() *Calculator {
	return &Calculator{}
}

func (c Calculator) Add(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		if sum == math.MaxInt || num == math.MaxInt {
			log.Print("The number is too big, not able to calculate the sum")
			return 0
		}
		sum += num
	}
	return sum
}
