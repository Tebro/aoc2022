package day4

import (
	"strconv"
	"strings"
)

func parseInts(ss []string) []int {
	res := make([]int, len(ss))
	for i, s := range ss {
		v, _ := strconv.Atoi(s)
		res[i] = v
	}
	return res
}

func part1(lines []string) int {
	total := 0

	for _, l := range lines {

		gnomes := strings.Split(l, ",")
		gnome1 := parseInts(strings.Split(gnomes[0], "-"))
		gnome2 := parseInts(strings.Split(gnomes[1], "-"))

		if (gnome1[0] >= gnome2[0] && gnome1[1] <= gnome2[1]) ||
			(gnome2[0] >= gnome1[0] && gnome2[1] <= gnome1[1]) {
			total++
		}

	}
	return total
}

func digitInRange(digit int, space []int) bool {
	return digit >= space[0] && digit <= space[1]
}

func part2(lines []string) int {
	total := 0

	for _, l := range lines {

		gnomes := strings.Split(l, ",")
		gnome1 := parseInts(strings.Split(gnomes[0], "-"))
		gnome2 := parseInts(strings.Split(gnomes[1], "-"))

		if digitInRange(gnome1[0], gnome2) ||
			digitInRange(gnome1[1], gnome2) ||
			digitInRange(gnome2[0], gnome1) ||
			digitInRange(gnome2[1], gnome1) {
			total++
		}

	}
	return total

}
