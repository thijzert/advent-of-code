package aoc21

import "testing"

func TestOrientationsAreUnique(t *testing.T) {
	pt := point3{3, 5, 7}
	for a := 0; a < 24; a++ {
		for b := 0; b < 24; b++ {
			if a == b {
				continue
			}

			pa := pt.Tr(orientation(a))
			pb := pt.Tr(orientation(b))
			if pa == pb {
				t.Errorf("Orientations %d and %d both result in %d", a, b, pa)
			}
		}
	}
}

func TestOrientationsAreConsistent(t *testing.T) {
	groundTruth := []point3{
		{-1, -1, 1},
		{-2, -2, 2},
		{-3, -3, 3},
		{-2, -3, 1},
		{5, 6, -4},
		{8, 0, 7},
	}

	transformed := [][]point3{
		{
			point3{1, -1, 1},
			point3{2, -2, 2},
			point3{3, -3, 3},
			point3{2, -1, 3},
			point3{-5, 4, -6},
			point3{-8, -7, 0},
		},
		{
			point3{-1, -1, -1},
			point3{-2, -2, -2},
			point3{-3, -3, -3},
			point3{-1, -3, -2},
			point3{4, 6, 5},
			point3{-7, 0, 8},
		},
		{
			point3{1, 1, -1},
			point3{2, 2, -2},
			point3{3, 3, -3},
			point3{1, 3, -2},
			point3{-4, -6, 5},
			point3{7, 0, 8},
		},
		{
			point3{1, 1, 1},
			point3{2, 2, 2},
			point3{3, 3, 3},
			point3{3, 1, 2},
			point3{-6, -4, -5},
			point3{0, 7, -8},
		},
	}

	for i, pts := range transformed {
		found := false
		for o := orientation(0); o < 24; o++ {
			found = true
			for j, pt := range groundTruth {
				ptb := pt.Tr(o)
				found = found && (ptb == pts[j])
			}
			if found {
				t.Logf("Scanner %d has orientation %d", i+1, o)
				for _, pt := range groundTruth {
					ptb := pt.Tr(o)
					t.Logf("  -  %d → %d", pt, ptb)
				}
				break
			}
		}

		if !found {
			t.Fail()
		}
	}
}
