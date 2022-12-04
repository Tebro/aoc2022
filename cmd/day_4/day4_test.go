package day4

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

func Test_runpart1(t *testing.T) {
	lines := strings.Split(data, "\n")
	res := part1(lines)
	t.Log(res)
	t.Fail()

}

func Test_runpart2(t *testing.T) {
	lines := strings.Split(data, "\n")
	res := part2(lines)
	t.Log(res)
	t.Fail()

}

func Test_part1(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want int
	}{
		// TODO: Add test cases.
		{
			name: "One matching line",
			args: []string{
				"2-5,3-4",
			},
			want: 1,
		},
		{
			name: "One matching line, switched",
			args: []string{
				"2-5,2-6",
			},
			want: 1,
		},
		{
			name: "Solo on the edge",
			args: []string{
				"6-6,2-6",
			},
			want: 1,
		},
		{
			name: "Mixed",
			args: []string{
				"2-3,4-5",
				"3-8,4-5",
				"4-8,7-10",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
