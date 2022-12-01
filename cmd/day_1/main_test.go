package main

import "testing"

func Benchmark_main(b *testing.B) {
	inputFileName = "../../input/day_1"
	for i := 0; i < b.N; i++ {
		main()
	}
}
