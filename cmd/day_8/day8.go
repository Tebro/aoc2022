package day8

import "sort"

func toInt(r rune) int {
	return int(r - '0')
}

func part1(lines []string) int {

	isOnEdge := func(x, y int) bool {
		return x == 0 || y == 0 || y == len(lines)-1 || x == len(lines[0])-1
	}

	sumVisible := len(lines) * len(lines[0]) // start with all, remove as we discover hidden

	for y, line := range lines {
	colLoop:
		for x, v := range line {
			if isOnEdge(x, y) {
				continue
			}
			treeSize := toInt(v)
			if v == 0 {
				continue colLoop
			}

			xFoundBefore := false
			xFoundAfter := false
			for x2, v2 := range line {
				treeSize2 := toInt(v2)
				if x2 < x {
					if treeSize2 >= treeSize {
						xFoundBefore = true
					}
				} else if x2 > x {
					if !xFoundBefore {
						continue colLoop
					}
					if treeSize2 >= treeSize {
						xFoundAfter = true
					}
				}
			}

			// If found before we never look after, if found after then we should check Y
			if !xFoundAfter {
				continue colLoop
			}

			yFoundBefore := false
			for y2, line2 := range lines {
				v2 := toInt(rune(line2[x]))
				if y2 < y {
					if v2 >= treeSize {
						yFoundBefore = true
					}
				} else if y2 > y {
					if !yFoundBefore {
						continue colLoop
					}
					if v2 >= treeSize {
						sumVisible--
						continue colLoop
					}
				}
			}

		}
	}
	return sumVisible
}

func part2(lines []string) int {

	isOnEdge := func(x, y int) bool {
		return x == 0 || y == 0 || y == len(lines)-1 || x == len(lines[0])-1
	}
	numTrees := len(lines) * len(lines[0]) // start with all, remove as we discover hidden
	results := make([]int, 0, numTrees)

	for y, line := range lines {
		for x, v := range line {
			if isOnEdge(x, y) {
				continue
			}

			treeSize := toInt(v)
			var up, right, down, left int = 1, 1, 1, 1

			// Check up
			for y2 := y - 1; y2 > 0; y2-- {
				if treeSize > toInt(rune(lines[y2][x])) {
					up++
				} else {
					break
				}
			}

			// check right
			for x2 := x + 1; x2 < len(line)-1; x2++ {
				if treeSize > toInt(rune(line[x2])) {
					right++
				} else {
					break
				}
			}

			// check down
			for y2 := y + 1; y2 < len(lines)-1; y2++ {
				if treeSize > toInt(rune(lines[y2][x])) {
					down++
				} else {
					break
				}
			}

			// check left
			for x2 := x - 1; x2 > 0; x2-- {
				if treeSize > toInt(rune(line[x2])) {
					left++
				} else {
					break
				}
			}
			results = append(results, up*right*down*left)
		}
	}

	sort.Ints(results)
	return results[len(results)-1]
}
