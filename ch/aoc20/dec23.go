package aoc20

import (
	"errors"
	"fmt"
	"log"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec23a(ctx ch.AOContext) error {
	input := "643719258"
	if len(ctx.Args) > 0 {
		input = ctx.Args[0]
	}

	cc := newCrabCup(input)
	for i := 0; i < 100; i++ {
		cc.Move(ctx.Debug)
	}

	ctx.Debug.Print(cc.Result())

	return errors.New("not implemented")
}

type crabCup struct {
	Cups    []int
	Grab    []int
	Current int
}

func newCrabCup(labeling string) *crabCup {
	rv := &crabCup{
		Cups:    make([]int, len(labeling)),
		Grab:    make([]int, 3),
		Current: 0,
	}
	for i, c := range labeling {
		rv.Cups[i] = int(c - '0')
	}
	return rv
}

func (cc *crabCup) String() string {
	rv := "cups: "
	for i, l := range cc.Cups {
		if i == cc.Current {
			rv += fmt.Sprintf("(%d)", l)
		} else {
			rv += fmt.Sprintf(" %d ", l)
		}
	}
	return rv
}

func (cc *crabCup) Result() string {
	for i, l := range cc.Cups {
		if l == 1 {
			rv := ""
			for j := range cc.Cups[1:] {
				rv += fmt.Sprintf("%d", cc.Cups[(i+j+1)%len(cc.Cups)])
			}
			return rv
		}
	}
	return "?"
}

func (cc *crabCup) Move(l *log.Logger) {
	l.Print(cc)

	L := len(cc.Cups)
	buf := make([]int, L)
	copy(buf, cc.Cups)

	for i := range cc.Grab {
		j := (cc.Current + i + 1) % L
		cc.Grab[i] = buf[j]
		buf[j] = 0
	}

	l.Printf("pick up: %d", cc.Grab)

	dest := cc.Cups[cc.Current]
	found := true
	for found {
		dest--
		if dest == 0 {
			dest += L
		}
		found = false
		for _, p := range cc.Grab {
			if p == dest {
				found = true
			}
		}
	}

	l.Printf("destination: %d", dest)

	isrc := cc.Current
	idest := cc.Current
	for range cc.Cups {
		if buf[isrc%L] == dest {
			cc.Cups[idest%L] = dest
			idest++
			for _, p := range cc.Grab {
				cc.Cups[idest%L] = p
				idest++
			}
			isrc++
		} else if buf[isrc%L] == 0 {
			isrc++
		} else {
			cc.Cups[idest%L] = buf[isrc%L]
			isrc++
			idest++
		}
	}

	cc.Current = (cc.Current + 1) % L
}

func Dec23b(ctx ch.AOContext) error {
	return errors.New("not implemented")
}
