package aoc21

import (
	"fmt"
	"sort"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec14a(ctx ch.AOContext) (interface{}, error) {
	sections, err := ctx.DataSections("inputs/2021/dec14.txt")
	if err != nil {
		return nil, err
	}

	template := readPolymer(sections[0][0])
	ctx.Printf("Polymer: %s", template)

	rules := readPolymerInsertions(sections[1])
	for i := 0; i < 10; i++ {
		template.ApplyRules(rules)
		if i < 4 {
			ctx.Printf("After step %d: %s", i+1, template)
		}
	}

	occ := template.Occurrences()
	rv := occ[len(occ)-1] - occ[0]

	return rv, nil
}

func Dec14b(ctx ch.AOContext) (interface{}, error) {
	sections, err := ctx.DataSections("inputs/2021/dec14.txt")
	if err != nil {
		return nil, err
	}

	template := readPolymerHisto(sections[0][0])

	rules := readPolymerInsertions(sections[1])
	for i := 0; i < 40; i++ {
		template.ApplyRules(rules)
	}

	occ := template.Occurrences()
	rv := occ[len(occ)-1] - occ[0]

	return rv, nil
}

type polymerInsertion struct {
	Left, Right rune
	Insert      rune
}

func readPolymerInsertions(rules []string) []polymerInsertion {
	var rv []polymerInsertion

	for _, line := range rules {
		var ins polymerInsertion
		if _, err := fmt.Sscanf(line, "%c%c -> %c", &ins.Left, &ins.Right, &ins.Insert); err == nil {
			rv = append(rv, ins)
		}
	}

	return rv
}

type polymerNode struct {
	Prev *polymerNode
	V    rune
	Next *polymerNode
}

func readPolymer(polymer string) *polymerNode {
	var last, rv *polymerNode

	for _, c := range polymer {
		n := &polymerNode{
			Prev: last,
			V:    c,
		}
		if last != nil {
			last.Next = n
		} else {
			rv = n
		}
		last = n
	}

	return rv
}

func (k *polymerNode) Insert(b *polymerNode) {
	after := k.Next
	end := b.Prev

	k.Next = b
	b.Prev = k

	after.Prev = end
	end.Next = after
}

func (k *polymerNode) Remove(until *polymerNode) {
	before := k.Prev
	after := until.Next

	before.Next = after
	after.Prev = before

	k.Prev = until
	until.Next = k
}

func (k *polymerNode) String() string {
	rv := fmt.Sprintf("%c", k.V)
	nd := k.Next
	for nd != nil {
		rv += fmt.Sprintf("%c", nd.V)
		nd = nd.Next
	}
	return rv
}

func (k *polymerNode) Contains(v rune) bool {
	if k.V == v {
		return true
	}
	// There's probably some double pointer magic to be done here
	nd := k.Next
	for nd != k {
		if nd.V == v {
			return true
		}
		nd = nd.Next
	}
	return false
}

func (k *polymerNode) ApplyRules(rules []polymerInsertion) {
	right := k.Next

	for right != nil {
		left := right.Prev

		for _, r := range rules {
			if left.V == r.Left && right.V == r.Right {
				n := &polymerNode{
					Prev: left,
					V:    r.Insert,
					Next: right,
				}
				left.Next = n
				right.Prev = n
			}
		}

		right = right.Next
	}
}

func (k *polymerNode) Occurrences() []int {
	occ := make(map[rune]int)

	r := k
	for r != nil {
		occ[r.V]++
		r = r.Next
	}

	rv := make([]int, 0, len(occ))
	for _, o := range occ {
		rv = append(rv, o)
	}
	sort.Ints(rv)
	return rv
}

type polytwo struct {
	left, right rune
}

type polymerHisto map[polytwo]int

func readPolymerHisto(polymer string) polymerHisto {
	rv := make(polymerHisto)
	var last rune
	for i, c := range polymer {
		if i > 0 {
			rv[polytwo{last, c}]++
		}
		rv[polytwo{0, c}]++
		last = c
	}
	return rv
}

func (hist polymerHisto) ApplyRules(rules []polymerInsertion) {
	toadd := make(polymerHisto)
	for _, r := range rules {
		m := hist[polytwo{r.Left, r.Right}]
		toadd[polytwo{r.Left, r.Right}] += -m
		toadd[polytwo{r.Insert, r.Right}] += m
		toadd[polytwo{r.Left, r.Insert}] += m
		toadd[polytwo{0, r.Insert}] += m
	}

	for k, m := range toadd {
		hist[k] += m
	}
}

func (hist polymerHisto) Occurrences() []int {
	rv := make([]int, 0)
	for k, o := range hist {
		if k.left == 0 {
			rv = append(rv, o)
		}
	}
	sort.Ints(rv)
	return rv
}
