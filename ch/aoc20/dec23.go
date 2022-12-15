package aoc20

import (
	"fmt"
	"log"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec23a(ctx ch.AOContext) (interface{}, error) {
	input := "643719258"
	if len(ctx.Args) > 0 {
		input = ctx.Args[0]
	}

	cc := newCrabCup(input, len(input))
	for i := 0; i < 100; i++ {
		cc.Move(ctx.Debug)
	}

	return cc.Result(), nil
}

type crabCup struct {
	Cups    *krabbyNode
	lookup  []*krabbyNode
	Grab    *krabbyNode
	Current *krabbyNode
}

func newCrabCup(labeling string, length int) *crabCup {
	rv := &crabCup{
		Cups:    &krabbyNode{V: 1},
		lookup:  make([]*krabbyNode, length),
		Grab:    nil,
		Current: nil,
	}
	for i := range rv.lookup {
		if i == 0 {
			rv.Current = rv.Cups
		} else {
			nd := &krabbyNode{
				Prev: rv.Current,
				V:    i + 1,
				Next: nil,
			}
			rv.Current.Next = nd
			rv.Current = nd
		}
		rv.lookup[i] = rv.Current
	}
	// Close the circle
	rv.Current.Next = rv.Cups
	rv.Cups.Prev = rv.Current
	rv.Current = rv.Cups

	// Overwrite labeling
	nd := rv.Current
	for _, c := range labeling {
		nd.V = int(c - '0')
		rv.lookup[nd.V-1] = nd
		nd = nd.Next
	}
	return rv
}

func (cc *crabCup) String() string {
	return "cups: " + cc.Current.String()
}

func (cc *crabCup) Result() string {
	start := cc.lookup[0]
	rv := ""
	nd := start.Next
	for nd != start {
		rv += fmt.Sprintf("%d", nd.V)
		nd = nd.Next
	}
	return rv
}

func (cc *crabCup) Move(l *log.Logger) {
	if l != nil {
		l.Print(cc)
	}

	cc.Grab = cc.Current.Next
	cc.Grab.Remove(cc.Grab.Next.Next)

	if l != nil {
		l.Printf("pick up: %s", cc.Grab)
	}

	dest := cc.Current.V
	found := true
	for found {
		dest--
		if dest == 0 {
			dest += len(cc.lookup)
		}
		found = cc.Grab.Contains(dest)
	}

	if l != nil {
		l.Printf("destination: %d", dest)
	}

	cc.lookup[dest-1].Insert(cc.Grab)
	cc.Current = cc.Current.Next
}

func Dec23b(ctx ch.AOContext) (interface{}, error) {
	input := "643719258"
	if len(ctx.Args) > 0 {
		input = ctx.Args[0]
	}

	cc := newCrabCup(input, 1000000)
	for i := 0; i < 10000000; i++ {
		cc.Move(nil)
		if i%1000000 == 0 {
			ctx.Debug.Printf("move %d - current: %d", i+1, cc.Current.V)
		}
	}

	a := cc.lookup[0].Next.V
	b := cc.lookup[0].Next.Next.V

	ctx.Debug.Printf("Value after 1:      %d", a)
	ctx.Debug.Printf("The one after that: %d", b)

	return a * b, nil
}

type krabbyNode struct {
	Prev *krabbyNode
	V    int
	Next *krabbyNode
}

func (k *krabbyNode) Insert(b *krabbyNode) {
	after := k.Next
	end := b.Prev

	k.Next = b
	b.Prev = k

	after.Prev = end
	end.Next = after
}

func (k *krabbyNode) Remove(until *krabbyNode) {
	before := k.Prev
	after := until.Next

	before.Next = after
	after.Prev = before

	k.Prev = until
	until.Next = k
}

func (k *krabbyNode) String() string {
	rv := fmt.Sprintf("%d", k.V)
	nd := k.Next
	for nd != k {
		rv += fmt.Sprintf(" %d", nd.V)
		nd = nd.Next
	}
	return rv
}

func (k *krabbyNode) Contains(v int) bool {
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
