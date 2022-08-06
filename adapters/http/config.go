package temphttp

type Calculator interface {
	Add(numbers []int) int
}

type RouterConfig struct {
	Calculator Calculator
}
