package day6

type Set[T comparable] struct {
	data map[T]bool
}

func CreateSet[T comparable](initial ...T) *Set[T] {
	data := make(map[T]bool)
	s := Set[T]{data: data}
	s.Add(initial...)
	return &s
}

func (s *Set[T]) Add(vs ...T) {
	for _, v := range vs {
		s.data[v] = true
	}
}

func (s *Set[T]) Len() int {
	return len(s.data)
}

func (s *Set[T]) Remove(v T) {
	delete(s.data, v)
}

// This was stupid... trying to be clever, only for part2 to come around and slap me
func part1(lines []string) int {

	l := lines[0]

	goingStrong := 2

	for i := range l {
		if i < 3 {
			continue
		}
		c := string(l[i])
		p1 := string(l[i-1])
		p2 := string(l[i-2])
		p3 := string(l[i-3])

		switch c {
		case p1:
			goingStrong = 1
		case p2:
			goingStrong = 2
		case p3:
			if goingStrong == 3 { // catches edgecase 'lssl'
				goingStrong = 3
			} else {
				goingStrong++
			}
		default:
			goingStrong++
		}
		if goingStrong == 4 {
			return i + 1
		}
	}

	return 0
}

func alg(numDistinct int, line string) int {

	for i := range line {
		if i < numDistinct-1 {
			continue
		}
		s := CreateSet[string]()

		for j := i; j > i-numDistinct; j-- {
			s.Add(string(line[j]))
		}

		if s.Len() == numDistinct {
			return i + 1
		}
	}
	return 0
}

func part2(lines []string) int {
	return alg(14, lines[0])
}
