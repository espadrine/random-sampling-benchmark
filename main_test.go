package main

import "testing"

const (
	base = 10
	size = 11
)

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sample(base, size)
	}
}

func BenchmarkRawSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rawSample(base, size)
	}
}
