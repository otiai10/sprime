package primes

import "testing"

func BenchmarkUntil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Until(123456)
	}
}