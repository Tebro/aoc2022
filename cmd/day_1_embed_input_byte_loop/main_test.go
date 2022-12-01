package main

import "testing"

func Benchmark_main(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
