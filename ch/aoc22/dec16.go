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

	for route, step := range network.Roadmap {
		if route>>16 == 0x4141 {
			ctx.Printf("From AA to %c%c: via %c%c", byte(route), byte(route>>8), byte(step), byte(step>>8))
		}
	}

	network.ElephantInTheRoom = true

	path, cost, err := dijkstra.ShortestPath(network)
	if err != nil {
		return nil, err
	}
	ctx.Printf("Pressure relieved: %d", 26*network.MaxFlow-cost)

	tlast := 0
	pressureRelieved := 0
	for _, p := range path {
		ctx.Printf("%s", p)
		pp, ok := p.(elephantInTheRoom)
		if !ok {
			panic("wrong position type")
		}
		pressureRelieved += (pp.FlowPerMinute) * (pp.Time - tlast)
		tlast = pp.Time
	}

	return pressureRelieved, nil
}

type valveNetwork struct {
	NameMask          map[uint16]uint64
	Flow              map[uint16]int
	Neighbours        map[uint16][]uint16
	Destinations      []uint16
	Roadmap           map[uint32]uint16
	ValveRequirement  uint64
	MaxFlow           int
	ElephantInTheRoom bool
	from, to          uint16
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
		Roadmap:    make(map[uint32]uint16),
	}

	name := func(s string) uint16 {
		return uint16(s[1])<<8 | uint16(s[0])
	}

	rv.Destinations = []uint16{0x4141}

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
		if r > 0 {
			rv.NameMask[p] = mask
			rv.ValveRequirement |= mask
			rv.MaxFlow += r
			mask <<= 1
			rv.Destinations = append(rv.Destinations, p)
		}
		rv.Flow[p] = r
		rv.Neighbours[p] = []uint16{}

		j := strings.IndexByte(line, ';')
		for _, nbv := range strings.Split(strings.TrimSpace(line[j+24:]), ", ") {
			rv.Neighbours[p] = append(rv.Neighbours[p], name(nbv))
		}
	}

	// Build road maps
	for _, rv.from = range rv.Destinations {
		for _, rv.to = range rv.Destinations {
			if rv.to == rv.from {
				continue
			}
			path, _, err := dijkstra.ShortestPath(rv)
			if err != nil {
				return nil, err
			}
			last := rv.from
			for _, p := range path {
				if pp, ok := p.(roadmapper); ok {
					rv.Roadmap[uint32(last)<<16|uint32(rv.to)] = uint16(pp)
					last = uint16(pp)
				}
			}
		}
	}
	rv.from, rv.to = 0, 0

	return rv, nil
}

func (n *valveNetwork) StartingPositions() []dijkstra.Position {
	if n.from != 0 && n.to != 0 {
		return []dijkstra.Position{roadmapper(n.from)}
	}

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

type roadmapper uint16

func (p roadmapper) Final(b dijkstra.Board) bool {
	bb, ok := b.(*valveNetwork)
	if !ok {
		return false
	}
	return uint16(p) == bb.to
}

func (p roadmapper) Adjacent(b dijkstra.Board, totalCost int) dijkstra.AdjacencyIterator {
	bb, ok := b.(*valveNetwork)
	if !ok {
		return dijkstra.DeadEnd()
	}

	adj := []dijkstra.Adj{}
	for _, nb := range bb.Neighbours[uint16(p)] {
		adj = append(adj, dijkstra.Adj{roadmapper(nb), 1})
	}
	return dijkstra.AdjacencyList(adj)
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
	Plan          uint16
	Elephant      uint16
	ElephantPlan  uint16
	FlowPerMinute int
	Time          int
	ValveMask     uint64
}

func (p elephantInTheRoom) String() string {
	name := [2]byte{byte(p.Position), byte(p.Position >> 8)}
	ename := [2]byte{byte(p.Elephant), byte(p.Elephant >> 8)}
	plan := [2]byte{byte(p.Plan), byte(p.Plan >> 8)}
	eplan := [2]byte{byte(p.ElephantPlan), byte(p.ElephantPlan >> 8)}
	return fmt.Sprintf("t=%d: %s→%s/%s→%s, %x → %d", p.Time, name, plan, ename, eplan, p.ValveMask, p.FlowPerMinute)
	return fmt.Sprintf("t=%d: %s/%s, %x → %d", p.Time, name, ename, p.ValveMask, p.FlowPerMinute)
	return fmt.Sprintf("Minute %d: you are at valve %s, the elephant is at valve %s, valves %x are open releasing %d pressure", p.Time, name, ename, p.ValveMask, p.FlowPerMinute)
}

func (p elephantInTheRoom) Final(b dijkstra.Board) bool {
	return p.Time == 26
}

func (p elephantInTheRoom) Adjacent(b dijkstra.Board, totalCost int) dijkstra.AdjacencyIterator {
	bb, ok := b.(*valveNetwork)
	if !ok {
		return dijkstra.DeadEnd()
	}
	TIME := 26
	if p.Time > TIME {
		return dijkstra.DeadEnd()
	}

	cost := bb.MaxFlow - p.FlowPerMinute

	// Make a copy of p with the time advanced
	pp := p
	pp.Time += 1

	if p.ValveMask == bb.ValveRequirement {
		//pp.Time = TIME
		//return dijkstra.AdjacencyList([]dijkstra.Adj{{pp, 0}})
	}

	// Positions if either you or the elephant opened a valve
	adj := []dijkstra.Adj{}

	var youMoves, eMoves []elephantInTheRoom

	thisValve := bb.NameMask[p.Position]
	if p.Plan != 0 && p.Position == p.Plan && thisValve != 0 && p.ValveMask&thisValve == 0 {
		q := pp
		q.Plan = 0
		q.ValveMask = q.ValveMask | thisValve
		q.FlowPerMinute += bb.Flow[p.Position]
		youMoves = append(youMoves, q)
	} else if p.Plan != 0 && p.Position != p.Plan {
		q := pp
		q.Position = bb.Roadmap[uint32(p.Position)<<16|uint32(p.Plan)]
		youMoves = append(youMoves, q)
	} else {
		for _, dst := range bb.Destinations {
			if dst == pp.Position {
				continue
			}
			if p.ValveMask&bb.NameMask[dst] != 0 {
				continue
			}
			q := pp
			q.Plan = dst
			q.Position = bb.Roadmap[uint32(q.Position)<<16|uint32(q.Plan)]
			youMoves = append(youMoves, q)
		}
		if len(youMoves) == 0 {
			q := pp
			q.Plan = 0
			youMoves = append(youMoves, pp)
		}
	}

	thisValve = bb.NameMask[p.Elephant]
	if p.ElephantPlan != 0 && p.Elephant == p.ElephantPlan && thisValve != 0 && p.ValveMask&thisValve == 0 {
		q := pp
		q.ElephantPlan = 0
		q.ValveMask = q.ValveMask | thisValve
		q.FlowPerMinute += bb.Flow[p.Elephant]
		eMoves = append(eMoves, q)
	} else if p.ElephantPlan != 0 && p.Elephant != p.ElephantPlan {
		q := pp
		q.Elephant = bb.Roadmap[uint32(p.Elephant)<<16|uint32(p.ElephantPlan)]
		eMoves = append(eMoves, q)
	} else {
		for _, dst := range bb.Destinations {
			if dst == pp.Elephant {
				continue
			}
			if p.ValveMask&bb.NameMask[dst] != 0 {
				continue
			}
			q := pp
			q.ElephantPlan = dst
			q.Elephant = bb.Roadmap[uint32(q.Elephant)<<16|uint32(q.ElephantPlan)]
			eMoves = append(eMoves, q)
		}
		if len(eMoves) == 0 {
			q := pp
			q.ElephantPlan = 0
			eMoves = append(eMoves, q)
		}
	}
	//if p.Time == 22 {
	//	log.Printf("from %v (%d)", p, bb.ValveRequirement)
	//	log.Printf("umoves: %v", youMoves)
	//	log.Printf("emoves: %v", eMoves)
	//}

	for _, you := range youMoves {
		for _, eleph := range eMoves {
			//if p.Time == 22 {
			//	log.Printf("combine {%v} and {%v}", you, eleph)
			//}
			if you.Position == eleph.Elephant && you.ValveMask == eleph.ValveMask {
				//continue
			}
			q := you
			q.Elephant = eleph.Elephant
			q.ElephantPlan = eleph.ElephantPlan
			q.ValveMask = q.ValveMask | eleph.ValveMask
			//q.FlowPerMinute += eleph.FlowPerMinute - pp.FlowPerMinute
			q.FlowPerMinute = 0
			for nd, m := range bb.NameMask {
				if q.ValveMask&m != 0 {
					q.FlowPerMinute += bb.Flow[nd]
				}
			}
			adj = append(adj, dijkstra.Adj{q, cost})
			//if p.Time == 22 {
			//	log.Printf("   → into {%v}", q)
			//}
		}
	}

	return dijkstra.AdjacencyList(adj)
}

func (p elephantInTheRoom) Hushcode() [4]uint64 {
	rv := [4]uint64{
		uint64(p.Position)<<32 | uint64(p.Elephant),
		p.ValveMask,
		uint64(p.Plan)<<32 | uint64(p.ElephantPlan),
		0,
	}

	if p.Position < p.Elephant {
		rv[0] = uint64(p.Elephant)<<32 | uint64(p.Position)
	}

	return rv
}
