package aoc22

import (
	"fmt"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/dijkstra"
)

func Dec16a(ctx ch.AOContext) (interface{}, error) {
	network, err := readValveNetwork(ctx, "inputs/2022/dec16.txt")
	if err != nil {
		return nil, err
	}

	path, cost, err := dijkstra.ShortestPath(network)
	if err != nil {
		return nil, err
	}

	for _, p := range path {
		ctx.Printf("%s", p)
	}

	return -cost, nil
}

var Dec16b ch.AdventFunc = nil

// func Dec16b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }

type valveNetwork struct {
	NameMask   map[uint16]uint64
	Flow       map[uint16]int
	Neighbours map[uint16][]uint16
}

func (n *valveNetwork) StartingPositions() []dijkstra.Position {
	return []dijkstra.Position{valveConfiguration{
		Position:      0x4141,
		FlowPerMinute: 0,
		Time:          1,
		ValveMask:     0,
	}}
}

type valveConfiguration struct {
	Position      uint16
	FlowPerMinute int
	Time          int8
	ValveMask     uint64
}

func (p valveConfiguration) String() string {
	name := [2]byte{byte(p.Position), byte(p.Position >> 8)}
	return fmt.Sprintf("Minute %d: at valve %s, valves %x are open releasing %d pressure", p.Time, name, p.ValveMask, p.FlowPerMinute)
}

func (p valveConfiguration) Final(b dijkstra.Board) bool {
	return p.Time == 31
}

func (p valveConfiguration) Adjacent(b dijkstra.Board, totalCost int) dijkstra.AdjacencyIterator {
	bb, ok := b.(*valveNetwork)
	if !ok {
		return dijkstra.DeadEnd()
	}
	if p.Time == 31 {
		return dijkstra.DeadEnd()
	}

	cost := -p.FlowPerMinute

	// Make a copy of p with the time advanced
	pp := p
	pp.Time += 1

	adj := []dijkstra.Adj{}

	thisValve := bb.NameMask[p.Position]
	if thisValve != 0 && p.ValveMask&thisValve == 0 && bb.Flow[p.Position] != 0 {
		q := pp
		q.ValveMask = q.ValveMask | thisValve
		q.FlowPerMinute += bb.Flow[p.Position]
		adj = append(adj, dijkstra.Adj{q, cost})
	}

	for _, nb := range bb.Neighbours[p.Position] {
		q := pp
		q.Position = nb
		adj = append(adj, dijkstra.Adj{q, cost})
	}

	return dijkstra.AdjacencyList(adj)
}

func readValveNetwork(ctx ch.AOContext, filename string) (*valveNetwork, error) {
	lines, err := ctx.DataLines(filename)
	if err != nil {
		return nil, err
	}

	rv := &valveNetwork{
		NameMask:   make(map[uint16]uint64),
		Flow:       make(map[uint16]int),
		Neighbours: make(map[uint16][]uint16),
	}

	name := func(s string) uint16 {
		return uint16(s[1])<<8 | uint16(s[0])
	}

	mask := uint64(1)

	for _, line := range lines {
		var n string
		var r int
		_, err := fmt.Sscanf(line, "Valve %2s has flow rate=%d;", &n, &r)
		if err != nil {
			ctx.Printf("Line: '%s'", line)
			return nil, err
		}

		p := name(n)
		rv.NameMask[p] = mask
		mask <<= 1
		rv.Flow[p] = r
		rv.Neighbours[p] = []uint16{}

		j := strings.IndexByte(line, ';')
		for _, nbv := range strings.Split(strings.TrimSpace(line[j+24:]), ", ") {
			rv.Neighbours[p] = append(rv.Neighbours[p], name(nbv))
		}
	}

	return rv, nil
}
