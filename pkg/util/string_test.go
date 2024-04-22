package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Unit tests
// go test -v ./...
func TestRandString(t *testing.T) {
	res, err := RandString(15)
	assert.NoError(t, err)
	assert.Equal(t, "0x78b9Ea5a46B732e", res, "string generated was not as expected.")
}

func TestRandStringInvalidInputNeg(t *testing.T) {
	_, err := RandString(0)
	assert.Error(t, err)
}

func TestRandStringInvalidInputZero(t *testing.T) {
	_, err := RandString(-1)
	assert.Error(t, err)
}

// Benchmark tests
// go test -v -benchmem -run=XXX -bench=. ./...
// TODO need a strategy for storing baselines and comparing with past resultsgit
func BenchmarkRandString10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RandString(10)
	}
}

func BenchmarkRandString100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RandString(100)
	}
}

func BenchmarkRandString1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RandString(1000)
	}
}
