package main

const loss, draw, win = 0, 3, 6
const rock, paper, scissor = 1, 2, 3

var results1 = map[string]int{
	"A X": rock + draw,
	"A Y": paper + win,
	"A Z": scissor + loss,
	"B X": rock + loss,
	"B Y": paper + draw,
	"B Z": scissor + win,
	"C X": rock + win,
	"C Y": paper + loss,
	"C Z": scissor + draw,
}

var results2 = map[string]int{
	"A Y": rock + draw,
	"A Z": paper + win,
	"A X": scissor + loss,
	"B X": rock + loss,
	"B Y": paper + draw,
	"B Z": scissor + win,
	"C Z": rock + win,
	"C X": paper + loss,
	"C Y": scissor + draw,
}

func alg(lines []string, rules map[string]int) int {
	total := 0
	for _, line := range lines {
		total += rules[line]
	}
	return total
}

func part1(lines []string) int {
	return alg(lines, results1)
}

func part2(lines []string) int {
	return alg(lines, results2)
}
