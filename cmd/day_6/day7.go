package day6

func (s *Set[T]) Has(v T) bool {
	_, ok := s.data[v]
	return ok
}
