package aoc20

import (
	"fmt"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/data"
)

func Dec16a(ctx ch.AOContext) (interface{}, error) {
	sections, err := ctx.DataSections("inputs/2020/dec16.txt")
	if err != nil {
		return nil, err
	}
	_, intervals := parseIntervalSets(sections[0])

	nearby := data.CSVInts(sections[2][1:])

	rv := 0

	for _, line := range nearby {
		for _, n := range line {
			found := false
			for _, itv := range intervals {
				found = found || itv.In(n)
			}
			if !found {
				rv += n
			}
		}
	}

	return rv, nil
}

func Dec16b(ctx ch.AOContext) (interface{}, error) {
	sections, err := ctx.DataSections("inputs/2020/dec16.txt")
	if err != nil {
		return nil, err
	}
	names, intervals := parseIntervalSets(sections[0])

	myTicket := data.CSVInts(sections[1][1:])[0]

	if len(names) != len(myTicket) {
		return nil, fmt.Errorf("trying to map %d fields onto %d values", len(names), len(myTicket))
	}

	nearby := data.CSVInts(sections[2][1:])
	var validTickets [][]int

	for _, line := range nearby {
		validTicket := true
		for _, n := range line {
			found := false
			for _, itv := range intervals {
				found = found || itv.In(n)
			}
			if !found {
				validTicket = false
			}
		}
		if validTicket {
			validTickets = append(validTickets, line)
		}
	}

	M := len(validTickets)
	ctx.Printf("valid tickets: %d", M)

	fieldMap := make([]int, len(names))
	for i := range fieldMap {
		fieldMap[i] = -1
	}
	validCounter := make([][]int, len(names))
	for i := range validCounter {
		validCounter[i] = make([]int, len(myTicket))
	}

	for _, line := range validTickets {
		for i, n := range line {
			for j, itv := range intervals {
				if itv.In(n) {
					validCounter[j][i]++
				}
			}
		}
	}

	changed := true
	for changed {
		changed = false

		for i, m := range fieldMap {
			if m != -1 {
				continue
			}

			found, idx := 0, 0
			for j, ct := range validCounter {
				if ct[i] == M {
					found++
					idx = j
				}
			}

			if found == 1 {
				fieldMap[i] = idx
				for j, ct := range validCounter {
					if j != idx {
						ct[i] = 0
					}
				}
				for j := range validCounter[idx] {
					if j != i {
						validCounter[idx][j] = 0
					}
				}
				changed = true
			}
		}
	}

	// ctx.Printf("Mapping: %d ; valid counts:\n%3d", fieldMap, validCounter)

	rv := 1
	for col, nameidx := range fieldMap {
		ctx.Printf("Column %2d: %s", col+1, names[nameidx])
		if len(names[nameidx]) > 9 && names[nameidx][0:9] == "departure" {
			rv *= myTicket[col]
		}
	}

	return rv, nil
}

type interval struct {
	Min, Max int
}

type intervalSet []interval

func (i intervalSet) In(v int) bool {
	for _, m := range i {
		if v >= m.Min && v <= m.Max {
			return true
		}
	}

	return false
}

func parseIntervalSets(lines []string) (names []string, intervals []intervalSet) {
	for _, l := range lines {
		lp := strings.SplitN(l, ": ", 2)
		if len(lp) != 2 {
			continue
		}

		iset := make(intervalSet, 0)
		itvs := strings.Split(lp[1], " or ")
		for _, itv := range itvs {
			var i interval
			fmt.Sscanf(itv, "%d-%d", &i.Min, &i.Max)
			iset = append(iset, i)
		}
		names = append(names, lp[0])
		intervals = append(intervals, iset)
	}

	return
}
