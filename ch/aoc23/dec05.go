package aoc23

import (
	"fmt"
	"sort"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

type AlmanacRange struct {
	Destination int
	Source      int
	Length      int
}

type AlmanacPage struct {
	From, To string
	Ranges   []AlmanacRange
}

func (ap AlmanacPage) Get(n int) int {
	for _, pg := range ap.Ranges {
		if n >= pg.Source && n < pg.Source+pg.Length {
			return pg.Destination + n - pg.Source
		}
	}
	return n
}

func (ap AlmanacPage) Len() int {
	return len(ap.Ranges)
}

func (ap AlmanacPage) Less(i, j int) bool {
	return ap.Ranges[i].Source < ap.Ranges[j].Source
}

func (ap AlmanacPage) Swap(i, j int) {
	ap.Ranges[i], ap.Ranges[j] = ap.Ranges[j], ap.Ranges[i]
}

func (a AlmanacPage) Merge(b AlmanacPage) AlmanacPage {
	rv := AlmanacPage{
		From: a.From,
		To:   b.To,
	}

	var joina, joinb func(rng AlmanacRange)
	joina = func(rng AlmanacRange) {
		if rng.Length <= 0 {
			return
		}
		for _, brn := range b.Ranges {
			offset := brn.Source - rng.Destination
			if rng.Destination <= brn.Source && rng.Length > offset {
				extent := rng.Destination + rng.Length - brn.Source
				// rng  [        o                      ]
				// brn           [               ]      e
				joina(AlmanacRange{rng.Destination, rng.Source, offset})
				if extent <= brn.Length {
					rv.Ranges = append(rv.Ranges, AlmanacRange{brn.Destination, rng.Source + offset, extent})
				} else {
					rv.Ranges = append(rv.Ranges, AlmanacRange{brn.Destination, rng.Source + offset, brn.Length})
					l := offset + brn.Length
					joina(AlmanacRange{rng.Destination + l, rng.Source + l, rng.Length - l})
				}
				return
			} else if brn.Source <= rng.Destination && brn.Length > -offset {
				extent := brn.Source + brn.Length - rng.Destination
				// rng  o      [          e   ]
				// brn  [     -o          ]
				if extent >= rng.Length {
					rv.Ranges = append(rv.Ranges, AlmanacRange{brn.Destination - offset, rng.Source, rng.Length})
				} else {
					rv.Ranges = append(rv.Ranges, AlmanacRange{brn.Destination - offset, rng.Source, extent})
					joina(AlmanacRange{rng.Destination + extent, rng.Source + extent, rng.Length - extent})
				}
				return
			}
		}
		rv.Ranges = append(rv.Ranges, rng)
	}
	joinb = func(brn AlmanacRange) {
		if brn.Length <= 0 {
			return
		}
		for _, rng := range a.Ranges {
			offset := brn.Source - rng.Destination
			extent := rng.Destination + rng.Length - brn.Source
			if rng.Destination <= brn.Source && rng.Length > offset {
				//   brn          [   e     ]
				//   rng [        o   ]
				joinb(AlmanacRange{brn.Destination + extent, brn.Source + extent, brn.Length - extent})
				return
			} else if brn.Source <= rng.Destination && brn.Length > -offset {
				//   brn [     -o            e         ]
				//   rng o      [            ]
				joinb(AlmanacRange{brn.Destination, brn.Source, -offset})
				joinb(AlmanacRange{brn.Destination + extent, brn.Source + extent, brn.Length - extent})
				return
			}
		}
		rv.Ranges = append(rv.Ranges, brn)
	}

	for _, rng := range a.Ranges {
		joina(rng)
	}

	for _, brn := range b.Ranges {
		joinb(brn)
	}

	sort.Sort(rv)

	return rv
}

type Almanac []AlmanacPage

func (a Almanac) String() string {
	if len(a) == 0 {
		return ""
	}

	rv := ""
	for _, pg := range a {
		rv += fmt.Sprintf("\n\n%s-to-%s map:", pg.From, pg.To)
		for _, rng := range pg.Ranges {
			rv += fmt.Sprintf("\n[%d:%d] → [%d:%d]", rng.Source, rng.Source+rng.Length-1, rng.Destination, rng.Destination+rng.Length-1)
		}
	}
	return rv[2:]
}

func Dec05Almanac(ctx ch.AOContext) (Almanac, []int, error) {
	sections, err := ctx.DataSections("inputs/2023/dec05.txt")
	if err != nil {
		return nil, nil, err
	}

	var rv Almanac
	var seeds []int

	if sections[0][0][:6] == "seeds:" {
		for _, s := range strings.Split(sections[0][0][6:], " ") {
			if s != "" {
				seeds = append(seeds, atoid(s, 0))
			}
		}
	}

	for _, sect := range sections[1:] {
		var pg AlmanacPage
		fmt.Sscanf(strings.ReplaceAll(sect[0], "-", " "), "%s to %s map:", &pg.From, &pg.To)
		for _, line := range sect[1:] {
			var rng AlmanacRange
			fmt.Sscanf(line, "%d %d %d", &rng.Destination, &rng.Source, &rng.Length)
			pg.Ranges = append(pg.Ranges, rng)
		}
		rv = append(rv, pg)
	}

	return rv, seeds, nil
}

func Dec05Lowest(ctx ch.AOContext, almanac Almanac, seeds [][2]int) (any, error) {
	lowest := 0x7fffffff
	for _, rng := range seeds {
		for j := 0; j < rng[1]; j++ {
			n := rng[0] + j
			//ctx.Printf("Seed %d:", n)
			m := "seed"
			for m != "location" {
				found := false
				for _, pg := range almanac {
					if pg.From == m {
						m = pg.To
						n = pg.Get(n)
						//ctx.Printf("   - %s %d", m, n)
						found = true
						break
					}
				}
				if !found {
					ctx.Printf("Can't find '%s' in the almanac", m)
					return nil, errFailed
				}
			}
			if n < lowest {
				lowest = n
			}
		}
	}

	return lowest, nil
}

func Dec05a(ctx ch.AOContext) (interface{}, error) {
	almanac, seeds, err := Dec05Almanac(ctx)
	if err != nil {
		return nil, err
	}

	seedRanges := make([][2]int, len(seeds))
	for i, n := range seeds {
		seedRanges[i][0] = n
		seedRanges[i][1] = 1
	}

	return Dec05Lowest(ctx, almanac, seedRanges)
}

func Dec05b(ctx ch.AOContext) (interface{}, error) {
	almanac, seeds, err := Dec05Almanac(ctx)
	if err != nil {
		return nil, err
	}

	almanacA := make(Almanac, 1)
	almanacA[0] = AlmanacPage{"seed", "seed", nil}
	for _, pg := range almanac {
		almanacA[0] = almanacA[0].Merge(pg)
	}
	almanacA[0].To = "location"
	ctx.Printf("Almanac A: %s", almanacA)

	almanacB := make(Almanac, 1)
	almanacB[0] = AlmanacPage{"location", "location", nil}
	for i := range almanac {
		pg := almanac[len(almanac)-i-1]
		almanacB[0] = pg.Merge(almanacB[0])
	}
	ctx.Printf("Almanac B: %s", almanacB)
	ctx.Printf("Are A and B the same: %v", almanacA.String() == almanacB.String())

	totalSeeds := 0
	seedRanges := make([][2]int, len(seeds)/2)
	for i, n := range seeds {
		seedRanges[i/2][i%2] = n
		if i%2 == 1 {
			totalSeeds += n
		}
	}

	// All almanac pages are in ascending order, so only check the first seed of each range
	var cheatyRanges [][2]int
	for _, rn := range seedRanges {
		cheatyRanges = append(cheatyRanges, [2]int{rn[0], 1})
		for _, page := range almanacA {
			for _, rng := range page.Ranges {
				if rng.Source > rn[0] && rng.Source-rn[0] <= rn[1] {
					cheatyRanges = append(cheatyRanges, [2]int{rng.Source, 1})
				}
			}
		}
	}

	ctx.Printf("Need to check %d seeds, but only checking %d", totalSeeds, len(cheatyRanges))
	if totalSeeds < 1000 {
		ans, err := Dec05Lowest(ctx, almanac, seedRanges)
		if err == nil {
			ctx.Printf("(actual answer: %d)", ans)
		}
	}

	// 13505501: too high
	// 14636072: too high
	// 15811003 (too high, presumably)
	// 18499111 (too high, presumably)
	// 56017390: too high

	return Dec05Lowest(ctx, almanacA, cheatyRanges)
}
