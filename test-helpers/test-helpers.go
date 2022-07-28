package test_helpers

import (
	"math/rand"
	"strconv"
	"time"
)

func RandomString() string {
	return strconv.Itoa(newRand().Intn(100000))
}

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
