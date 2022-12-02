package main

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed input.txt
var data string

func Benchmark_part1(b *testing.B) {
	lines := strings.Split(data, "\n")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(lines)
	}
}

func Benchmark_part2(b *testing.B) {
	lines := strings.Split(data, "\n")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2(lines)
	}
}
