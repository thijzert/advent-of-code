package aoc22

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec20a(ctx ch.AOContext) (interface{}, error) {
	ints, err := ctx.DataAsInts("inputs/2022/dec20.txt")
	if err != nil {
		return nil, err
	}
	//ints = []int{1, 2, -3, 3, -2, 0, 4}

	var head *dll
	nodes := make([]*dll, len(ints))

	for i, v := range ints {
		nd := &dll{V: v}
		nodes[i] = nd
		if i == 0 {
			nd.Prev, nd.Next = nd, nd
		} else {
			head.Append(nd)
		}
		head = nd
	}
	for head.V != 0 {
		head = head.Next
	}

	ctx.Printf("initial arrangement: %s", head)
	for _, nd := range nodes {
		nd.Move(nd.V)
		ctx.Printf("%d moves between %d and %d", nd.V, nd.Prev.V, nd.Next.V)
		//ctx.Printf("%s", head)
	}
	ctx.Printf("final arrangement: %s", head)

	rv := 0
	for _, steps := range []int{1000, 2000, 3000} {
		steps = steps % len(nodes)
		nd := head
		for i := 0; i < steps; i++ {
			nd = nd.Next
		}
		ctx.Printf("Found coordinate component: %d", nd.V)
		rv += nd.V
	}
	return rv, nil
}

var Dec20b ch.AdventFunc = nil

// func Dec20b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }

type dll struct {
	Prev *dll
	V    int
	Next *dll
}

func (n *dll) String() string {
	rv := fmt.Sprintf("%d", n.V)
	i := n.Next
	for i != nil && i != n {
		rv += fmt.Sprintf(", %d", i.V)
		i = i.Next
	}
	return rv
}

func (n *dll) Append(b *dll) {
	b.Prev = n
	b.Next = n.Next
	n.Next.Prev = b
	n.Next = b
}

func (n *dll) Move(spaces int) {
	npr, nnx := n.Prev, n.Next
	npr.Next, nnx.Prev = nnx, npr

	i := npr
	if spaces < 0 {
		for j := 0; j < -spaces; j++ {
			i = i.Prev
		}
	} else {
		for j := 0; j < spaces; j++ {
			i = i.Next
		}
	}

	inx := i.Next
	i.Next, inx.Prev = n, n
	n.Next, n.Prev = inx, i
}
