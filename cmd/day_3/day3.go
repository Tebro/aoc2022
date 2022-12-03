package day3

func characterPrio(c rune) int {
	if c >= 'a' {
		return int(c-'a') + 1
	}
	if c <= 'Z' {
		return int(c-'A') + 27
	}
	return 0
}

func setFromLine(l []rune) map[rune]bool {
	s := map[rune]bool{}
	for _, c := range l {
		s[c] = true
	}
	return s
}

func setContains(m map[rune]bool, c rune) bool {
	_, ok := m[c]
	return ok
}

func part1(lines []string) int {
	total := 0
	for _, l := range lines {
		cs := []rune(l)
		halfWay := len(cs) / 2
		first := cs[:halfWay]
		second := cs[halfWay:]

		set := setFromLine(first)

		for _, c := range second {
			if setContains(set, c) {
				total += characterPrio(c)
				break
			}
		}

	}

	return total
}

func part2(lines []string) int {
	total := 0
	for i := 0; i < len(lines); i += 3 {
		line1 := lines[i]
		set2 := setFromLine([]rune(lines[i+1]))
		set3 := setFromLine([]rune(lines[i+2]))

		for _, c := range line1 {
			if setContains(set2, c) && setContains(set3, c) {
				total += characterPrio(c)
				break
			}
		}
	}
	return total
}
