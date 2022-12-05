package cube

import "testing"

func TestOrientationsAreUnique(t *testing.T) {
	pt := Point3{3, 5, 7}
	for a := 0; a < 24; a++ {
		for b := 0; b < 24; b++ {
			if a == b {
				continue
			}

			pa := pt.Tr(Orientation(a))
			pb := pt.Tr(Orientation(b))
			if pa == pb {
				t.Errorf("Orientations %d and %d both result in %d", a, b, pa)
			}
		}
	}
}

func TestOrientationsAreConsistent(t *testing.T) {
	groundTruth := []Point3{
		{-1, -1, 1},
		{-2, -2, 2},
		{-3, -3, 3},
		{-2, -3, 1},
		{5, 6, -4},
		{8, 0, 7},
	}

	transformed := [][]Point3{
		{
			Point3{1, -1, 1},
			Point3{2, -2, 2},
			Point3{3, -3, 3},
			Point3{2, -1, 3},
			Point3{-5, 4, -6},
			Point3{-8, -7, 0},
		},
		{
			Point3{-1, -1, -1},
			Point3{-2, -2, -2},
			Point3{-3, -3, -3},
			Point3{-1, -3, -2},
			Point3{4, 6, 5},
			Point3{-7, 0, 8},
		},
		{
			Point3{1, 1, -1},
			Point3{2, 2, -2},
			Point3{3, 3, -3},
			Point3{1, 3, -2},
			Point3{-4, -6, 5},
			Point3{7, 0, 8},
		},
		{
			Point3{1, 1, 1},
			Point3{2, 2, 2},
			Point3{3, 3, 3},
			Point3{3, 1, 2},
			Point3{-6, -4, -5},
			Point3{0, 7, -8},
		},
	}

	for i, pts := range transformed {
		found := false
		for o := Orientation(0); o < 24; o++ {
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
