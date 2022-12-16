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

func Dec16b(ctx ch.AOContext) (interface{}, error) {
	network, err := readValveNetwork(ctx, "inputs/2022/dec16.txt")
	if err != nil {
		return nil, err
	}

	network.ElephantInTheRoom = true

	path, _, err := dijkstra.ShortestPath(network)
	if err != nil {
		return nil, err
	}

	tlast, plast := 0, 0
	pressureRelieved := 0
	for _, p := range path {
		ctx.Printf("%s", p)
		pp, ok := p.(elephantInTheRoom)
		if !ok {
			panic("wrong position type")
		}
		pressureRelieved += (pp.FlowPerMinute) * (pp.Time - tlast)
		tlast = pp.Time
		plast = pp.FlowPerMinute
	}
	pressureRelieved -= plast

	return pressureRelieved, nil
}

type valveNetwork struct {
	NameMask          map[uint16]uint64
	Flow              map[uint16]int
	Neighbours        map[uint16][]uint16
	ValveRequirement  uint64
	MaxFlow           int
	ElephantInTheRoom bool
}

func (n *valveNetwork) StartingPositions() []dijkstra.Position {
	if n.ElephantInTheRoom {
		return []dijkstra.Position{elephantInTheRoom{
			Position:      0x4141,
			Elephant:      0x4141,
			FlowPerMinute: 0,
			Time:          1,
			ValveMask:     0,
		}}
	}
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

type elephantInTheRoom struct {
	Position      uint16
	Elephant      uint16
	FlowPerMinute int
	Time          int
	ValveMask     uint64
}

func (p elephantInTheRoom) String() string {
	name := [2]byte{byte(p.Position), byte(p.Position >> 8)}
	ename := [2]byte{byte(p.Elephant), byte(p.Elephant >> 8)}
	return fmt.Sprintf("t=%d: %s/%s, %x → %d", p.Time, name, ename, p.ValveMask, p.FlowPerMinute)
	return fmt.Sprintf("Minute %d: you are at valve %s, the elephant is at valve %s, valves %x are open releasing %d pressure", p.Time, name, ename, p.ValveMask, p.FlowPerMinute)
}

func (p elephantInTheRoom) Final(b dijkstra.Board) bool {
	return p.Time == 27
}

func (p elephantInTheRoom) Adjacent(b dijkstra.Board, totalCost int) dijkstra.AdjacencyIterator {
	bb, ok := b.(*valveNetwork)
	if !ok {
		return dijkstra.DeadEnd()
	}
	TIME := 27
	if p.Time == TIME {
		return dijkstra.DeadEnd()
	}

	//cost := bb.MaxFlow - p.FlowPerMinute

	// Make a copy of p with the time advanced
	pp := p
	pp.Time += 1

	if p.ValveMask == bb.ValveRequirement {
		pp.Time = TIME
		return dijkstra.AdjacencyList([]dijkstra.Adj{{pp, 0}})
	}

	var youMoves, eMoves []elephantInTheRoom

	// Positions if either you or the elephant opened a valve
	adj := []dijkstra.Adj{}

	thisValve := bb.NameMask[p.Position]
	if thisValve != 0 && p.ValveMask&thisValve == 0 && bb.Flow[p.Position] != 0 {
		q := pp
		q.ValveMask = q.ValveMask | thisValve
		q.FlowPerMinute += bb.Flow[p.Position]
		youMoves = append(youMoves, q)
	}
	thisValve = bb.NameMask[p.Elephant]
	if thisValve != 0 && p.ValveMask&thisValve == 0 && bb.Flow[p.Elephant] != 0 {
		q := pp
		q.ValveMask = q.ValveMask | thisValve
		q.FlowPerMinute += bb.Flow[p.Elephant]
		eMoves = append(eMoves, q)
	}

	for _, nb := range bb.Neighbours[p.Position] {
		q := pp
		q.Position = nb
		youMoves = append(youMoves, q)
	}
	for _, nb := range bb.Neighbours[p.Elephant] {
		q := pp
		q.Elephant = nb
		eMoves = append(eMoves, q)
	}

	for _, you := range youMoves {
		for _, eleph := range eMoves {
			if you.Position == eleph.Elephant && you.ValveMask == eleph.ValveMask {
				continue
			}
			q := you
			q.Elephant = eleph.Elephant
			q.ValveMask = q.ValveMask | eleph.ValveMask
			q.FlowPerMinute += eleph.FlowPerMinute - pp.FlowPerMinute
			cost := (TIME - p.Time) * (bb.MaxFlow - q.FlowPerMinute)
			adj = append(adj, dijkstra.Adj{q, cost})
		}
	}

	return dijkstra.AdjacencyList(adj)
}

func (p elephantInTheRoom) Hashcode() [4]uint64 {
	rv := [4]uint64{
		uint64(p.Position)<<32 | uint64(p.Elephant),
		p.ValveMask,
		0,
		0,
	}

	if p.Position < p.Elephant {
		rv[0] = uint64(p.Elephant)<<32 | uint64(p.Position)
	}

	return rv
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
		if r > 0 {
			rv.ValveRequirement |= mask
			rv.MaxFlow += r
			mask <<= 1
		}
		rv.Flow[p] = r
		rv.Neighbours[p] = []uint16{}

		j := strings.IndexByte(line, ';')
		for _, nbv := range strings.Split(strings.TrimSpace(line[j+24:]), ", ") {
			rv.Neighbours[p] = append(rv.Neighbours[p], name(nbv))
		}
	}

	return rv, nil
}
