package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputFileName = "input/day_2"

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

func main() {

	f, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	total1 := 0
	total2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		score1 := results1[parts[0]][parts[1]]
		score2 := results2[parts[0]][parts[1]]

		total1 += score1
		total2 += score2
	}

	fmt.Println("Part 1 score is:", total1)
	fmt.Println("Part 2 score is:", total2)
}
