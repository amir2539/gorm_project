package helper

import (
	"math/rand"
	"time"
)

func RandNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
