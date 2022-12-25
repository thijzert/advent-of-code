package aoc22

import (
	"github.com/thijzert/advent-of-code/ch"
)

func Dec25a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2022/dec25.txt")
	if err != nil {
		return nil, err
	}

	rv := 0
	for _, l := range lines {
		n := fromSNAFU(l)
		s := toSNAFU(n)
		ctx.Printf("%s: %d → %s", l, n, s)
		if l != s {
			return nil, errFailed
		}
		rv += n
	}

	return toSNAFU(rv), nil
}

func Dec25b(ctx ch.AOContext) (interface{}, error) {
	return "You make a smoothie with all fifty stars and deliver it to the reindeer!", nil
}

func fromSNAFU(s string) int {
	rv := 0
	five := 1
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		if c == '=' {
			rv -= 2 * five
		} else if c == '-' {
			rv -= five
		} else if c == '1' {
			rv += five
		} else if c == '2' {
			rv += 2 * five
		}
		five = 5 * five
	}
	return rv
}

func toSNAFU(n int) string {
	if n == 0 {
		return "0"
	} else if n < 0 {
		panic(n)
	}

	digits := [5]string{"0", "1", "2", "=", "-"}
	rv := ""
	for n > 0 {
		m := n % 5
		rv = digits[m] + rv
		n /= 5
		if m >= 3 {
			n++
		}
	}
	return rv
}
