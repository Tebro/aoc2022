package day7

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed input.txt
var data string

func Test_runpart1(t *testing.T) {
	lines := strings.Split(data, "\n")
	res, err := part1(lines)
	t.Log(res, err)
	t.Fail()

}

func Test_runpart2(t *testing.T) {
	lines := strings.Split(data, "\n")
	res, err := part2(lines)
	t.Log(res, err)
	t.Fail()

}

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
