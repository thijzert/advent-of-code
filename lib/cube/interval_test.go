package cube

import "testing"

func TestIntervalOverlap(t *testing.T) {
}

func TestIntervalFullyContain(t *testing.T) {
}

func TestIntervalSet(t *testing.T) {
	cases := []struct {
		Expected string
		IVS      *IntervalSet
	}{
		{"{[1 3]}", NewIntervalSet(Interval{1, 3})},
		{"{[1 3], [5 7]}", NewIntervalSet(Interval{1, 3}, Interval{5, 7})},
		{"{[1 3], [5 7]}", NewIntervalSet(Interval{5, 7}, Interval{1, 3})},
		{"{[1 7]}", NewIntervalSet(Interval{1, 3}, Interval{5, 7}, Interval{2, 6})},
		{"{[1 8]}", NewIntervalSet(Interval{1, 3}, Interval{5, 7}, Interval{2, 8})},
		{"{[1 3], [5 8]}", NewIntervalSet(Interval{1, 3}, Interval{5, 7}, Interval{5, 8})},
		{"{[1 7]}", NewIntervalSet(Interval{1, 3}, Interval{5, 7}, Interval{4, 6})},
	}

	for _, tc := range cases {
		t.Logf("intervalset %s", tc.IVS)
		if tc.IVS.String() != tc.Expected {
			t.Errorf("Expected: %s, got %s", tc.Expected, tc.IVS)
		}
	}
}
