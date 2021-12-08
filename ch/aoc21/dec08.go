package aoc21

import (
	"errors"
	"fmt"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec08a(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2021/dec08.txt")
	if err != nil {
		return err
	}

	numberHistogram := make([]int, 10)

	for l, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " | ")
		if len(parts) != 2 {
			return fmt.Errorf("line %d: invalid format '%s'", l+1, line)
		}
		digits := strings.Split(parts[1], " ")
		mapping, err := extractMicrowaveMapping(strings.Split(parts[0], " "))
		if err != nil {
			ctx.Printf("Mapping so far: '%c'", mapping)
			return err
		}

		for _, d := range digits {
			n := mapping.GetDigit(d)
			if n >= 0 && n <= 9 {
				numberHistogram[n]++
			} else {
				ctx.Printf("%s is %d", d, n)
			}
		}
	}

	ctx.Printf("Occurrences: %d", numberHistogram)
	ctx.FinalAnswer.Print(numberHistogram[1] + numberHistogram[4] + numberHistogram[7] + numberHistogram[8])
	return nil
}

func Dec08b(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2021/dec08.txt")
	if err != nil {
		return err
	}

	sum := 0

	for l, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " | ")
		if len(parts) != 2 {
			return fmt.Errorf("line %d: invalid format '%s'", l+1, line)
		}

		passcode := 0
		digits := strings.Split(parts[1], " ")
		mapping, err := extractMicrowaveMapping(strings.Split(parts[0], " "))
		if err != nil {
			ctx.Printf("Mapping so far: '%c'", mapping)
			return err
		}

		for _, d := range digits {
			n := mapping.GetDigit(d)
			if n >= 0 && n <= 9 {
				passcode = 10*passcode + n
			} else {
				ctx.Printf("%s is %d", d, n)
			}
		}

		sum += passcode
		ctx.Printf("Passcode: %d", passcode)
	}

	ctx.FinalAnswer.Print(sum)
	return nil
}

var microwaveNumbers [10][7]bool = [10][7]bool{
	{true, true, true, false, true, true, true},
	{false, false, true, false, false, true, false},
	{true, false, true, true, true, false, true},
	{true, false, true, true, false, true, true},
	{false, true, true, true, false, true, false},
	{true, true, false, true, false, true, true},
	{true, true, false, true, true, true, true},
	{true, false, true, false, false, true, false},
	{true, true, true, true, true, true, true},
	{true, true, true, true, false, true, true},
}

type microwaveMapping [7]rune

func extractMicrowaveMapping(uniqueCombinations []string) (microwaveMapping, error) {
	lengths := make([][]string, 8)
	for _, s := range uniqueCombinations {
		if len(s) < 2 || len(s) > 7 {
			return microwaveMapping{}, errors.New("invalid combination '%s'")
		}
		lengths[len(s)] = append(lengths[len(s)], s)
	}
	if len(uniqueCombinations) != 10 ||
		len(lengths[2]) != 1 || len(lengths[3]) != 1 || len(lengths[4]) != 1 || len(lengths[7]) != 1 ||
		len(lengths[5]) != 3 || len(lengths[6]) != 3 {
		return microwaveMapping{}, errors.New("invalid combinations")
	}

	rv := microwaveMapping{}
	for i := range rv {
		rv[i] = ' '
	}

	digits := make([]string, 10)
	digits[1] = lengths[2][0]
	digits[7] = lengths[3][0]
	digits[4] = lengths[4][0]
	digits[8] = lengths[7][0]

	var ovr, diff []rune

	// Use 1 and 7 to find segment a
	_, diff = rv.intersect(digits[7], digits[1])
	if len(diff) != 1 {
		return rv, fmt.Errorf("invalid overlap between 1 ('%s') and 7 ('%s')", digits[1], digits[7])
	}
	rv[0] = diff[0]

	// Find 1 and 6, and map segment c and f
	for _, six := range lengths[6] {
		ovr, diff = rv.intersect(digits[1], six)
		if len(ovr) == 1 && len(diff) == 1 {
			digits[6] = six
			rv[2] = diff[0]
			rv[5] = ovr[0]
			break
		}
	}

	// Use 4 and 1 to find 0, and map segment d
	for _, zero := range lengths[6] {
		_, diff = rv.intersect(digits[4], digits[1], zero)
		if len(diff) == 1 {
			digits[0] = zero
			rv[3] = diff[0]
			break
		}
	}

	// Find 9, and map segment e
	for _, nine := range lengths[6] {
		if nine == digits[6] || nine == digits[0] {
			continue
		}
		_, diff = rv.intersect(digits[6], nine, digits[1])
		if len(diff) == 1 {
			rv[4] = diff[0]
			digits[9] = nine
		}
	}

	// Find 3
	for _, three := range lengths[5] {
		ovr, _ = rv.intersect(three, digits[1])
		if len(ovr) == 2 {
			digits[3] = three
		}
	}
	// Use 3 and 4 to map segment b
	_, diff = rv.intersect(digits[4], digits[3])
	if len(diff) != 1 {
		return rv, fmt.Errorf("invalid difference between 3 ('%s') and 4 ('%s')", digits[3], digits[4])
	}
	rv[1] = diff[0]

	// Use 9 and the amalgamation of 4 and 7 to map segment g
	_, diff = rv.intersect(digits[9], digits[4]+digits[7])
	if len(diff) != 1 {
		return rv, fmt.Errorf("invalid difference between 9 ('%s') and 47 ('%s')", digits[9], digits[4]+digits[7])
	}
	rv[6] = diff[0]

	// Sanity check
	for _, c := range rv {
		if c == ' ' {
			return rv, fmt.Errorf("failed to infer segment mapping")
		}
	}

	return rv, nil
}

func (microwaveMapping) intersect(a string, bs ...string) (overlap, difference []rune) {
	for _, c := range a {
		unique := true
		common := true
		for _, b := range bs {
			found := false
			for _, d := range b {
				found = found || (d == c)
				unique = unique && (d != c)
			}
			common = common && found
			unique = unique && !found
		}
		if common {
			overlap = append(overlap, c)
		}
		if unique {
			difference = append(difference, c)
		}
	}
	return
}

func (m microwaveMapping) GetDigit(segments string) int {
	power := [7]bool{}
	for _, s := range segments {
		for i, c := range m {
			power[i] = power[i] || (c == s)
		}
	}

	for n, state := range microwaveNumbers {
		if state == power {
			return n
		}
	}

	return -1
}
