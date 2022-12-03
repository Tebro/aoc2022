package day3

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

func Benchmark_part1String(b *testing.B) {
	lines := strings.Split(data, "\n")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1String(lines)
	}
}

func Benchmark_part2(b *testing.B) {
	lines := strings.Split(data, "\n")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part1(lines)
	}
}

func Benchmark_part2String(b *testing.B) {
	lines := strings.Split(data, "\n")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		part2String(lines)
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

func Test_characterPrio(t *testing.T) {
	tests := []struct {
		name string
		arg  rune
		want int
	}{
		// TODO: Add test cases.
		{
			name: "a should be 1",
			arg:  'a',
			want: 1,
		},
		{
			name: "z should be 26",
			arg:  'z',
			want: 26,
		},
		{
			name: "A should be 27",
			arg:  'A',
			want: 27,
		},
		{
			name: "Z should be 52",
			arg:  'Z',
			want: 52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterPrio(tt.arg); got != tt.want {
				t.Errorf("characterPrio() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_test(t *testing.T) {
	data := "ABCabc"
	arr := []rune(data)

	half := len(arr) / 2

	t.Log(arr[:half], arr[half:])

	t.Fail()
}
