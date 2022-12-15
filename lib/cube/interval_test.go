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
		{"{[-419580 4455895]}", NewIntervalSet(
			Interval{2626362, 2698718},
			Interval{3446645, 4455895},
			Interval{2743455, 3530103},
			Interval{2984131, 3008329},
			Interval{2743455, 2765321},
			Interval{-419580, 1468740},
			Interval{2718599, 2718599},
			Interval{2221087, 2605813},
			Interval{132734, 2626362},
			Interval{3018417, 3943501},
			Interval{3635521, 4108239},
			Interval{2718599, 2931813},
			Interval{2743455, 2862295},
			Interval{2640977, 2718599})},
	}

	for _, tc := range cases {
		t.Logf("intervalset %s", tc.IVS)
		if tc.IVS.String() != tc.Expected {
			t.Errorf("Expected: %s, got %s", tc.Expected, tc.IVS)
		}
	}
}
