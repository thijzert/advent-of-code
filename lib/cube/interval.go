package cube

import (
	"fmt"
)

// Interval represent an inclusive integer interval, with A <= B
type Interval struct {
	A, B int
}

func (a Interval) Length() int {
	return a.B - a.A + 1
}

func (a Interval) Contains(x int) bool {
	return x >= a.A && x <= a.B
}

func (a Interval) Overlap(b Interval) (Interval, bool) {
	if b.A >= a.A && b.A <= a.B {
		if b.B < a.B {
			return b, true
		} else {
			return Interval{b.A, a.B}, true
		}
	} else if b.B >= a.A && b.B <= a.B {
		if a.A < b.A {
			return b, true
		} else {
			return Interval{a.A, b.B}, true
		}
	} else if b.A <= a.A && b.B >= a.B {
		return a, true
	}
	return Interval{}, false
}

func (a Interval) Union(b Interval) (Interval, bool) {
	if a.B+1 == b.A {
		return Interval{a.A, b.B}, true
	} else if b.B+1 == a.A {
		return Interval{b.A, a.B}, true
	} else if b.A >= a.A && b.A <= a.B {
		if b.B < a.B {
			return a, true
		} else {
			return Interval{a.A, b.B}, true
		}
	} else if b.B >= a.A && b.B <= a.B {
		if a.A < b.A {
			return a, true
		} else {
			return Interval{b.A, a.B}, true
		}
	} else if b.A <= a.A && b.B >= a.B {
		return b, true
	}
	return Interval{}, false
}

func (a Interval) FullyContains(b Interval) bool {
	return b.A >= a.A && b.A <= a.B && b.B >= a.A && b.B <= a.B
}

func iabs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type IntervalSet struct {
	I []Interval
}

func NewIntervalSet(intervals ...Interval) *IntervalSet {
	rv := &IntervalSet{}
	for _, x := range intervals {
		rv.Add(x)
	}
	return rv
}

func (s *IntervalSet) Add(x Interval) {
	if len(s.I) == 0 {
		s.I = append(s.I, x)
		return
	}

	idx := -1
	for j, y := range s.I {
		if x.A < y.A {
			continue
		}
		idx = j
		break
	}

	if idx == -1 {
		s.I = append(s.I, x)
		copy(s.I[1:], s.I[:len(s.I)-1])
		s.I[0] = x
		idx = 0
	} else if y, overlap := x.Union(s.I[idx]); overlap {
		s.I[idx] = y
	} else {
		idx++
		s.I = append(s.I, x)
		copy(s.I[idx+1:], s.I[idx:len(s.I)-1])
		s.I[idx] = x
	}

	for len(s.I) > idx+1 {
		y, overlap := s.I[idx].Union(s.I[idx+1])
		if !overlap {
			break
		}
		s.I[idx] = y
		s.I = append(s.I[:idx+1], s.I[idx+2:]...)
	}
}

func (s *IntervalSet) String() string {
	if len(s.I) == 0 {
		return "{}"
	}

	rv, sep := "", "{"
	for _, i := range s.I {
		rv += fmt.Sprintf("%s[%d %d]", sep, i.A, i.B)
		sep = ", "
	}
	return rv + "}"
}

func (s *IntervalSet) Length() int {
	rv := 0
	for _, iv := range s.I {
		rv += iv.Length()
	}
	return rv
}

func (s *IntervalSet) FullyContains(v Interval) bool {
	for _, a := range s.I {
		if a.FullyContains(v) {
			return true
		}
	}
	return false
}
