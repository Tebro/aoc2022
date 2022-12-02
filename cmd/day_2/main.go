package main

import (
	"strings"
)

const loss, draw, win = 0, 3, 6
const rock, paper, scissor = 1, 2, 3

var results1 = map[string]map[string]int{
	"A": {
		"X": rock + draw,
		"Y": paper + win,
		"Z": scissor + loss,
	},
	"B": {
		"X": rock + loss,
		"Y": paper + draw,
		"Z": scissor + win,
	},
	"C": {
		"X": rock + win,
		"Y": paper + loss,
		"Z": scissor + draw,
	},
}

var results2 = map[string]map[string]int{
	"A": {
		"Y": rock + draw,
		"Z": paper + win,
		"X": scissor + loss,
	},
	"B": {
		"X": rock + loss,
		"Y": paper + draw,
		"Z": scissor + win,
	},
	"C": {
		"Z": rock + win,
		"X": paper + loss,
		"Y": scissor + draw,
	},
}

func alg(lines []string, rules map[string]map[string]int) int {
	total := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		total += rules[parts[0]][parts[1]]
	}
	return total
}

func part1(lines []string) int {
	return alg(lines, results1)
}

func part2(lines []string) int {
	return alg(lines, results2)
}
