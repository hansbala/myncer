package core

func NewSet[T comparable](v ...T) Set[T] {
	return ToSet(v)
}

func ToSet[T comparable](vs []T) Set[T] {
	r := make(Set[T], len(vs) /*hint*/)
	r.AddAll(vs)
	return r
}

type Set[T comparable] map[T]struct{}

func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s Set[T]) Add(v T) Set[T] {
	s[v] = struct{}{}
	return s
}

func (s Set[T]) AddAll(vs []T) Set[T] {
	for _, v := range vs {
		s.Add(v)
	}
	return s
}

func (s Set[T]) Delete(sa ...T) Set[T] {
	for _, v := range sa {
		if s.Contains(v) {
			delete(s, v)
		}
	}
	return s
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) ToArray() []T {
	r := []T{}
	for k := range s {
		r = append(r, k)
	}
	return r
}
