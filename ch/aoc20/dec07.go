package aoc20

import (
	"fmt"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec07a(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2020/dec07.txt")
	if err != nil {
		return err
	}

	bgr := parseBagRules(lines)
	trans := whichBagsCanContain(bgr, "shiny gold")

	rv := 0
	for colour, v := range trans {
		if v > 0 && colour != "shiny gold" {
			rv++
		}
	}
	ctx.FinalAnswer.Print(rv)
	return nil
}

func Dec07b(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2020/dec07.txt")
	if err != nil {
		return err
	}

	bgr := parseBagRules(lines)
	trans := totalBagSize(bgr)

	ctx.FinalAnswer.Print(trans["shiny gold"])
	return errNotImplemented
}

type bagRule struct {
	Bag   string
	Count int
}

func parseBagRules(lines []string) map[string][]bagRule {
	rv := make(map[string][]bagRule)

	for _, rule := range lines {
		comp := strings.Split(rule, " bags contain ")
		if len(comp) != 2 {
			continue
		}

		var contains []bagRule
		if comp[1] != "no other bags." {
			for _, rule := range strings.Split(comp[1], ", ") {
				rule = strings.TrimRight(rule, "s.,")
				r := bagRule{}
				var c string
				fmt.Sscanf(rule, "%d %s %s bag", &r.Count, &r.Bag, &c)
				r.Bag += " " + c
				contains = append(contains, r)
			}
		}

		rv[comp[0]] = contains
	}

	return rv
}

func totalBagSize(bagRules map[string][]bagRule) map[string]int {
	rv := make(map[string]int)

	for colour, rules := range bagRules {
		if len(rules) == 0 {
			rv[colour] = 0
		}
	}

	changed := true
	for changed {
		changed = false

		for colour, rules := range bagRules {
			if _, ok := rv[colour]; ok {
				continue
			}

			valid := true
			value := 0
			for _, r := range rules {
				n, ok := rv[r.Bag]
				valid = valid && ok
				value += (n + 1) * r.Count
			}

			if valid {
				changed = true
				rv[colour] = value
			}
		}
	}

	return rv
}

func whichBagsCanContain(bagRules map[string][]bagRule, initialBag string) map[string]int {
	rv := make(map[string]int)

	for colour, rules := range bagRules {
		if len(rules) == 0 {
			rv[colour] = 0
		}
	}

	rv[initialBag] = 1

	changed := true
	for changed {
		changed = false

		for colour, rules := range bagRules {
			if _, ok := rv[colour]; ok {
				continue
			}

			valid := true
			value := 0
			for _, r := range rules {
				n, ok := rv[r.Bag]
				valid = valid && ok
				value += n * r.Count
			}

			if valid {
				changed = true
				rv[colour] = value
			}
		}
	}

	return rv
}
