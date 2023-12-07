package aoc23

import (
	"log"
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

func CamelCardLess(a, b byte) bool {
	aa, bb := -1, -1
	for i, c := range "AKQJT98765432" {
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

type CamelPoker []CamelHand

func (cp CamelPoker) Len() int {
	return len(cp)
}

func (cp CamelPoker) Swap(i, j int) {
	cp[i], cp[j] = cp[j], cp[i]
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
	for k, f := range labels {
		aa, bb := f(cp[i].Counts), f(cp[j].Counts)
		if aa && !bb {
			log.Printf("Hand %s has label %v", cp[i].Hand, k)
			return false
		} else if !aa && bb {
			log.Printf("Hand %s has label %v", cp[j].Hand, k)
			return true
		} else if aa && bb {
			log.Printf("Hand %s and %s have label %v", cp[i].Hand, cp[j].Hand, k)
			for k, c := range cp[i].Hand {
				if CamelCardLess(c, cp[j].Hand[k]) {
					return false
				} else if CamelCardLess(cp[j].Hand[k], c) {
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

	var cp CamelPoker
	for _, line := range lines {
		var p CamelHand
		for i, c := range line[:5] {
			p.Hand[i] = byte(c)
		}
		p.Counts = Dec7CountCards(p.Hand)
		p.Bid = atoid(line[6:], -1)
		cp = append(cp, p)
		ctx.Printf("Hand: %s, counts: %v", p.Hand, p.Counts)
	}

	sort.Sort(cp)

	answer := 0
	for i, hand := range cp {
		ctx.Printf("%3d: %s (%d)", i+1, hand.Hand, hand.Bid)
		answer += (i + 1) * hand.Bid
	}
	return answer, nil
}

var Dec07b ch.AdventFunc = nil

// func Dec07b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }
