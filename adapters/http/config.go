package temphttp

type Calculator interface {
	Add(numbers []int) int
}

type Fibonacci interface {
	FibonacciNumber(number int) (string, error)
}

type RouterConfig struct {
	Calculator Calculator
	Fibonacci  Fibonacci
}
