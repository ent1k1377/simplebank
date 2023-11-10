package util

import (
	randon "math/rand"
	"strings"
	"time"
)

const (
	alphabet = "qwertyuiopasdfghjklzxcvbnm"
)

var rand *randon.Rand

func init() {
	source := randon.NewSource(time.Now().UnixNano())
	rand = randon.New(source)
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	la := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(la)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	lc := len(currencies)
	return currencies[rand.Intn(lc)]
}
