package cube

// Interval represent an inclusive integer interval, with A <= B
type Interval struct {
	A, B int
}

func (a Interval) Contains(x int) bool {
	return x >= a.A && x <= a.B
}

func (a Interval) Overlap(b Interval) (Interval, bool) {
	if b.A >= a.A && b.A <= a.B {
		if b.B < a.B {
			return Interval{b.A, b.B}, true
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

func (a Interval) FullyContains(b Interval) bool {
	return b.A >= a.A && b.A <= a.B && b.B >= a.A && b.B <= a.B
}

func iabs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
