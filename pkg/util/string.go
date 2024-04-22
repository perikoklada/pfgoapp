package util

import (
	"fmt"
	"math/rand"
)

var randx = rand.NewSource(42)

// RandString returns a random string of length n.
func RandString(n int) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("cannot generate string of invalid length %d", n)
	}

	const letterBytes = "0123456789ABCDEabcde"
	const (
		letterIdxBits = 5                    // 5 bits to represent a character index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	b := make([]byte, n+2)
	b[0] = '0'
	b[1] = 'x'
	effectiveStartIdx := 2

	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := len(b)-1, randx.Int63(), letterIdxMax; i >= effectiveStartIdx; {
		if remain == 0 {
			cache, remain = randx.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b), nil
}
