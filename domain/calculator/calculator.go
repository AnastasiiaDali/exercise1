package calculator

import (
	"log"
	"math"
)

func Add(numbers []int) int {
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
