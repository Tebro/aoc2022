package day6

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed input.txt
var data string

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
