package aoc21

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func signum(a int) int {
	if a < 0 {
		return -1
	} else if a > 0 {
		return 1
	}
	return 0
}

func min(vals ...int) int {
	rv := vals[0]
	for _, v := range vals {
		if v < rv {
			rv = v
		}
	}
	return rv
}

func max(vals ...int) int {
	rv := vals[0]
	for _, v := range vals {
		if v > rv {
			rv = v
		}
	}
	return rv
}
