package day5

import (
	"strconv"
	"strings"
)

type item[T any] struct {
	value T
	next  *item[T]
}

type Stack[T any] struct {
	top  *item[T]
	size int
}

func (s *Stack[T]) Len() int {
	return s.size
}

func (s *Stack[T]) Push(value T) {
	s.top = &item[T]{
		value: value,
		next:  s.top,
	}
	s.size++
}

func (s *Stack[T]) Pop() T {
	if s.Len() > 0 {
		value := s.top.value
		s.top = s.top.next
		s.size--
		return value
	}
	return *new(T)
}

// This empties the original stack!
func (s *Stack[T]) Invert() *Stack[T] {
	ns := Stack[T]{}
	for s.size > 0 {
		ns.Push(s.Pop())
	}
	return &ns
}

func part1(lines []string) string {

	initHasBeenParsed := false

	stacks := make([]*Stack[string], 9)
	for i := range stacks {
		stacks[i] = &Stack[string]{}
	}

	for _, l := range lines {
		if !initHasBeenParsed {
			if string(l[1]) == "1" || l == "" {
				if !initHasBeenParsed {
					for i := range stacks {
						stacks[i] = stacks[i].Invert()
					}
				}
				initHasBeenParsed = true
				continue
			}

			for i := 0; i < len(l); i += 4 {
				if v := string(l[i+1]); v != "" && v != " " {
					stack := stacks[i/4]
					stack.Push(v)
				}
			}

		} else if strings.HasPrefix(l, "move") {
			parts := strings.Split(l, " ")
			numToMove, _ := strconv.Atoi(parts[1])
			src, _ := strconv.Atoi(parts[3])
			dst, _ := strconv.Atoi(parts[5])

			for i := 0; i < numToMove; i++ {
				stacks[dst-1].Push(stacks[src-1].Pop())
			}

		} else {
			// Well shit
		}
	}

	res := make([]string, len(stacks))
	for _, s := range stacks {
		res = append(res, s.Pop())
	}
	return strings.Join(res, "")
}

func part2(lines []string) string {

	initHasBeenParsed := false

	stacks := make([]*Stack[string], 9)
	for i := range stacks {
		stacks[i] = &Stack[string]{}
	}

	for _, l := range lines {
		if !initHasBeenParsed {
			if string(l[1]) == "1" || l == "" {
				if !initHasBeenParsed {
					for i := range stacks {
						stacks[i] = stacks[i].Invert()
					}
				}
				initHasBeenParsed = true
				continue
			}

			for i := 0; i < len(l); i += 4 {
				if v := string(l[i+1]); v != "" && v != " " {
					stack := stacks[i/4]
					stack.Push(v)
				}
			}

		} else if strings.HasPrefix(l, "move") {
			parts := strings.Split(l, " ")
			numToMove, _ := strconv.Atoi(parts[1])
			src, _ := strconv.Atoi(parts[3])
			dst, _ := strconv.Atoi(parts[5])

			tmpStack := Stack[string]{}

			for i := 0; i < numToMove; i++ {
				tmpStack.Push(stacks[src-1].Pop())
			}

			for tmpStack.Len() > 0 {
				stacks[dst-1].Push(tmpStack.Pop())
			}

		} else {
			// Well shit
		}
	}

	res := make([]string, len(stacks))
	for _, s := range stacks {
		res = append(res, s.Pop())
	}
	return strings.Join(res, "")
}
