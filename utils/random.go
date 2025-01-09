package utils

import (
	"math/rand"
	"strings"
	"time"
)

var rng *rand.Rand

const letters = "abcdefghijklmnopqrstuvwxyz"

var currencies = []string{"USD", "EUR", "GEL"}

// without seed, rand will generate same values on every run
func init() {
	source := rand.NewSource(time.Now().UnixNano())

	rng = rand.New(source)
}

func RandomInt(min, max int64) int64 {
	return rng.Int63n(max-min+1) + min
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(letters)
	for i := 0; i < n; i++ {
		sb.WriteByte(letters[rng.Intn(k)])
	}
	return sb.String()
}

func RandomCurrency() string {
	return currencies[rng.Intn(len(currencies))]
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}
