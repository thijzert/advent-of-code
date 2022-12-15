package aoc22

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

type distressPacket struct {
	Int  int
	List []distressPacket
}

func parseDistressPacket(buf []byte) (distressPacket, int) {
	if buf[0] >= '0' && buf[0] <= '9' {
		i := 1
		for buf[i] >= '0' && buf[i] <= '9' {
			i++
		}
		var rv distressPacket
		fmt.Sscan(string(buf[:i]), &rv.Int)
		return rv, i
	}
	if buf[0] == '[' {
		rv := distressPacket{List: make([]distressPacket, 0)}
		if buf[1] == ']' {
			return rv, 2
		}

		i := 1
		for len(buf[i:]) > 0 {

			p, n := parseDistressPacket(buf[i:])
			rv.List = append(rv.List, p)
			i += n
			if buf[i] == ',' {
				i++
			} else if buf[i] == ']' {
				return rv, i + 1
			}
		}
	}
	panic("cannot parse distress packet " + string(buf))
	return distressPacket{}, 0
}

func (p distressPacket) String() string {
	if p.List == nil {
		return fmt.Sprintf("%d", p.Int)
	}
	if len(p.List) == 0 {
		return "[]"
	}
	rv, sep := "", "["
	for _, q := range p.List {
		rv += sep + q.String()
		sep = ","
	}
	return rv + "]"
}

var indent = ""

func (p distressPacket) compare(q distressPacket) int {
	//log.Printf("%sCompare %s and %s", indent, p, q)
	if p.List == nil && q.List == nil {
		if p.Int > q.Int {
			return 1
		} else if p.Int < q.Int {
			return -1
		}
		return 0
	}
	if p.List != nil && q.List != nil {
		for i, n := range p.List {
			if len(q.List) <= i {
				//log.Printf("%s%s is shorter", indent, q)
				return 1
			}
			indent = indent + "  "
			c := n.compare(q.List[i])
			indent = indent[:len(indent)-2]
			if c != 0 {
				return c
			}
		}
		if len(q.List) > len(p.List) {
			return -1
		}
		return 0
	}
	if p.List == nil {
		return distressPacket{List: []distressPacket{p}}.compare(q)
	}
	if q.List == nil {
		return p.compare(distressPacket{List: []distressPacket{q}})
	}
	panic(fmt.Sprintf("cannot compare: %s %s", p, q))
	return 0
}

func Dec13a(ctx ch.AOContext) (interface{}, error) {
	pairs, err := ctx.DataSections("inputs/2022/dec13.txt")
	if err != nil {
		return nil, err
	}

	ooo := 0
	for j, pair := range pairs {
		ctx.Printf("")
		ctx.Printf("pair %d", j+1)
		left, _ := parseDistressPacket([]byte(pair[0]))
		right, _ := parseDistressPacket([]byte(pair[1]))

		ctx.Printf("%s  %s", pair[0], pair[1])
		ctx.Printf("%s  %s", left, right)
		if left.compare(right) <= 0 {
			ooo += j + 1
			ctx.Printf("packets are in the right order")
		} else {
			ctx.Printf("packets are out of order: %v is more", left)
		}
	}

	return ooo, nil
}

func Dec13b(ctx ch.AOContext) (interface{}, error) {
	pairs, err := ctx.DataSections("inputs/2022/dec13.txt")
	if err != nil {
		return nil, err
	}

	idxA, idxB := 1, 2
	sepA, _ := parseDistressPacket([]byte("[[2]]"))
	sepB, _ := parseDistressPacket([]byte("[[6]]"))

	for _, pair := range pairs {
		for _, line := range pair {
			pkt, _ := parseDistressPacket([]byte(line))
			if pkt.compare(sepA) < 0 {
				idxA++
			}
			if pkt.compare(sepB) < 0 {
				idxB++
			}
		}
	}

	return idxA * idxB, nil
}
