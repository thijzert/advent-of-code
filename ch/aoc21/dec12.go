package aoc21

import (
	"strings"
	"unicode"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec12a(ctx ch.AOContext) error {
	caveA, caveB, caveC := dec12ExampleCaves()

	ctx.Print(caveA.NPathsTo("end", 0))
	ctx.Print(caveB.NPathsTo("end", 0))
	ctx.Print(caveC.NPathsTo("end", 0))

	lines, err := ctx.DataLines("inputs/2021/dec12.txt")
	if err != nil {
		return err
	}
	cave := parseCaveSystem(lines)

	ctx.FinalAnswer.Print(cave.NPathsTo("end", 0))
	return nil
}

func Dec12b(ctx ch.AOContext) error {
	caveA, caveB, caveC := dec12ExampleCaves()

	ctx.Print(caveA.NPathsTo("end", 1))
	ctx.Print(caveB.NPathsTo("end", 1))
	ctx.Print(caveC.NPathsTo("end", 1))

	lines, err := ctx.DataLines("inputs/2021/dec12.txt")
	if err != nil {
		return err
	}
	cave := parseCaveSystem(lines)

	ctx.FinalAnswer.Print(cave.NPathsTo("end", 1))
	return nil
}

func dec12ExampleCaves() (*subcave, *subcave, *subcave) {
	caveA := parseCaveSystem([]string{
		"start-A", "start-b", "A-c", "A-b",
		"b-d", "A-end", "b-end",
	})

	caveB := parseCaveSystem([]string{
		"dc-end", "HN-start", "start-kj", "dc-start", "dc-HN",
		"LN-dc", "HN-end", "kj-sa", "kj-HN", "kj-dc",
	})

	caveC := parseCaveSystem([]string{
		"fs-end", "he-DX", "fs-he", "start-DX", "pj-DX",
		"end-zg", "zg-sl", "zg-pj", "pj-he", "RW-he",
		"fs-DX", "pj-RW", "zg-RW", "start-pj", "he-WI",
		"zg-he", "pj-fs", "start-RW",
	})

	return caveA, caveB, caveC
}

func parseCaveSystem(lines []string) *subcave {
	rv := make(map[string]*subcave)

	for _, l := range lines {
		pts := strings.Split(l, "-")
		if len(pts) != 2 {
			continue
		}

		if rv[pts[0]] == nil {
			rv[pts[0]] = &subcave{Name: pts[0]}
		}
		if rv[pts[1]] == nil {
			rv[pts[1]] = &subcave{Name: pts[1]}
		}

		rv[pts[0]].Neighbours = append(rv[pts[0]].Neighbours, rv[pts[1]])
		rv[pts[1]].Neighbours = append(rv[pts[1]].Neighbours, rv[pts[0]])
	}

	for n, c := range rv {
		c.Large = unicode.IsUpper(rune(n[0]))
	}

	return rv["start"]
}

type subcave struct {
	Name       string
	Large      bool
	Visited    int
	Neighbours []*subcave
}

func (c *subcave) NPathsTo(otherCave string, smallCaveBudget int) int {
	rv := 0
	c.Visited++

	for _, oc := range c.Neighbours {
		if oc.Name == otherCave {
			rv++
		} else if oc.Name == "start" || oc.Name == "end" {
		} else if oc.Large || oc.Visited == 0 {
			rv += oc.NPathsTo(otherCave, smallCaveBudget)
		} else if oc.Visited == 1 && smallCaveBudget > 0 {
			rv += oc.NPathsTo(otherCave, smallCaveBudget-1)
		}
	}

	c.Visited--
	return rv
}
