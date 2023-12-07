package aoc23

import (
	"sort"

	"github.com/thijzert/advent-of-code/ch"
)

type CamelFunc func(hand [5]CamelCardCount) bool

type CamelCardCount struct {
	Card  byte
	Count int
}

func Dec7CountCards(hand [5]byte) [5]CamelCardCount {
	var rv [5]CamelCardCount
	for _, c := range hand {
		for i, cc := range rv {
			if cc.Card == c || cc.Card == 0 {
				rv[i].Card = c
				rv[i].Count++
				break
			}
		}
	}

	// Bubblesort the result, so that the most frequent card is at the top
	if rv[1].Count > rv[0].Count {
		rv[0], rv[1] = rv[1], rv[0]
	}
	if rv[2].Count > rv[1].Count {
		rv[1], rv[2] = rv[2], rv[1]
	}
	if rv[1].Count > rv[0].Count {
		rv[0], rv[1] = rv[1], rv[0]
	}
	if rv[3].Count > rv[2].Count {
		rv[2], rv[3] = rv[3], rv[2]
	}
	if rv[2].Count > rv[1].Count {
		rv[1], rv[2] = rv[2], rv[1]
	}
	if rv[1].Count > rv[0].Count {
		rv[0], rv[1] = rv[1], rv[0]
	}

	return rv
}

func CamelCardLess(a, b byte, cardOrder string) bool {
	aa, bb := -1, -1
	for i, c := range cardOrder {
		if byte(c) == a {
			aa = i
		}
		if byte(c) == b {
			bb = i
		}
	}
	return aa < bb
}

func Dec7Fiveof(hand [5]CamelCardCount) bool {
	return hand[0].Count == 5
}
func Dec7Fourof(hand [5]CamelCardCount) bool {
	return hand[0].Count == 4
}
func Dec7FullHouse(hand [5]CamelCardCount) bool {
	return hand[0].Count == 3 && hand[1].Count == 2
}
func Dec7Threeof(hand [5]CamelCardCount) bool {
	return hand[0].Count == 3
}
func Dec7TwoPair(hand [5]CamelCardCount) bool {
	return hand[0].Count == 2 && hand[1].Count == 2
}
func Dec7OnePair(hand [5]CamelCardCount) bool {
	return hand[0].Count == 2
}
func Dec7HighCard(hand [5]CamelCardCount) bool {
	return true
}

type CamelHand struct {
	Hand   [5]byte
	Bid    int
	Counts [5]CamelCardCount
}

type CamelPoker struct {
	Hands     []CamelHand
	CardOrder string
}

func (cp CamelPoker) Len() int {
	return len(cp.Hands)
}

func (cp CamelPoker) Swap(i, j int) {
	cp.Hands[i], cp.Hands[j] = cp.Hands[j], cp.Hands[i]
}

func (cp CamelPoker) Less(i, j int) bool {
	labels := [7]CamelFunc{
		Dec7Fiveof,
		Dec7Fourof,
		Dec7FullHouse,
		Dec7Threeof,
		Dec7TwoPair,
		Dec7OnePair,
		Dec7HighCard,
	}
	for _, f := range labels {
		aa, bb := f(cp.Hands[i].Counts), f(cp.Hands[j].Counts)
		if aa && !bb {
			return false
		} else if !aa && bb {
			return true
		} else if aa && bb {
			for k, c := range cp.Hands[i].Hand {
				if CamelCardLess(c, cp.Hands[j].Hand[k], cp.CardOrder) {
					return false
				} else if CamelCardLess(cp.Hands[j].Hand[k], c, cp.CardOrder) {
					return true
				}
			}
		}
	}
	return false
}

func Dec07a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec07.txt")
	if err != nil {
		return nil, err
	}

	cp := CamelPoker{
		CardOrder: "AKQJT98765432",
	}
	for _, line := range lines {
		var p CamelHand
		for i, c := range line[:5] {
			p.Hand[i] = byte(c)
		}
		p.Counts = Dec7CountCards(p.Hand)
		p.Bid = atoid(line[6:], -1)
		cp.Hands = append(cp.Hands, p)
	}

	sort.Sort(cp)

	answer := 0
	for i, hand := range cp.Hands {
		ctx.Printf("%3d: %s (%d)", i+1, hand.Hand, hand.Bid)
		answer += (i + 1) * hand.Bid
	}
	return answer, nil
}

func Dec07b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec07.txt")
	if err != nil {
		return nil, err
	}

	cp := CamelPoker{
		CardOrder: "AKQT98765432J",
	}
	for _, line := range lines {
		var p CamelHand
		for i, c := range line[:5] {
			p.Hand[i] = byte(c)
		}
		p.Counts = Dec7CountCards(p.Hand)
		p.Bid = atoid(line[6:], -1)

		for i, ct := range p.Counts {
			if ct.Card == 'J' {
				jokers := ct.Count
				copy(p.Counts[i:], p.Counts[i+1:])
				p.Counts[4].Count = 0
				p.Counts[0].Count += jokers
				break
			}
		}

		cp.Hands = append(cp.Hands, p)
	}

	sort.Sort(cp)

	answer := 0
	for i, hand := range cp.Hands {
		ctx.Printf("%3d: %s (%d)", i+1, hand.Hand, hand.Bid)
		answer += (i + 1) * hand.Bid
	}
	return answer, nil
}
