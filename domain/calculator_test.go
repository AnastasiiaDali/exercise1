package domain

import "testing"

func TestAdd(t *testing.T) {

	t.Run("Should take any number of integers and print out the sum.", func(t *testing.T) {
		numbers := []int{5, 4, 2, -10, 32, 14}
		got := Add(numbers)
		want := 47

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
