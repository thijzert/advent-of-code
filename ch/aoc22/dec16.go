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

	//ctx.Printf("Network: %v", network)
	m := make(maskMap)
	fillMaskMap(m, network, 30, 0, "AA", "AA", 0)
	ctx.Printf("%d paths examined", len(m))

	max := valvePath{}
	for _, vp := range m {
		if vp.PressureRelieved > max.PressureRelieved {
			max = vp
		}
	}
	ctx.Printf("Best path: %s", max.Path)

	return max.PressureRelieved, nil
}

func Dec16b(ctx ch.AOContext) (interface{}, error) {
	return nil, errNotImplemented
}

type valvePath struct {
	Path             string
	PressureRelieved int
}

type maskMap map[uint32]valvePath

func fillMaskMap(m maskMap, network *valveNetwork, timeRemaining int, mask uint32, position string, path string, relieved int) {
	if relieved > m[mask].PressureRelieved {
		m[mask] = valvePath{
			Path:             path,
			PressureRelieved: relieved,
		}
	}

	for dest, destm := range network.Mask {
		if mask&destm != 0 {
			continue
		}
		ntr := timeRemaining - network.Distances[position+"-"+dest] - 1
		if ntr > 0 {
			fillMaskMap(m, network, ntr, mask|destm, dest, path+"-"+dest, relieved+ntr*network.Flow[dest])
		}
	}
}

type valveNetwork struct {
	Mask       map[string]uint32
	Flow       map[string]int
	Distances  map[string]int
	Neighbours map[string][]string
	to, from   string
}

func readValveNetwork(ctx ch.AOContext, name string) (*valveNetwork, error) {
	lines, err := ctx.DataLines(name)
	if err != nil {
		return nil, err
	}

	rv := &valveNetwork{
		Mask:       make(map[string]uint32),
		Flow:       make(map[string]int),
		Distances:  make(map[string]int),
		Neighbours: make(map[string][]string),
	}

	var rate int
	mask := uint32(1)
	for _, line := range lines {
		var name string
		_, err := fmt.Sscanf(line, "Valve %2s has flow rate=%d;", &name, &rate)
		if err != nil {
			ctx.Printf("Line: '%s'", line)
			return nil, err
		}

		if rate > 0 {
			rv.Mask[name] = mask
			mask <<= 1
		}
		rv.Flow[name] = rate
		rv.Neighbours[name] = []string{}

		j := strings.IndexByte(line, ';')
		for _, nb := range strings.Split(strings.TrimSpace(line[j+24:]), ", ") {
			rv.Neighbours[name] = append(rv.Neighbours[name], nb)
		}
	}

	for rv.from, rate = range rv.Flow {
		if rate == 0 && rv.from != "AA" {
			continue
		}
		for rv.to, rate = range rv.Flow {
			if rate == 0 || rv.to == rv.from {
				continue
			}

			_, dist, err := dijkstra.ShortestPath(rv)
			if err != nil {
				return nil, err
			}
			k := rv.from + "-" + rv.to
			if rv.Distances[k] == 0 || dist < rv.Distances[k] {
				rv.Distances[k] = dist
			}
		}
	}

	rv.Neighbours = nil
	rv.to, rv.from = "", ""

	return rv, nil
}

func (n *valveNetwork) StartingPositions() []dijkstra.Position {
	return []dijkstra.Position{roadmapper(n.from)}
}

type roadmapper string

func (p roadmapper) Final(b dijkstra.Board) bool {
	bb, ok := b.(*valveNetwork)
	if !ok {
		return false
	}
	return string(p) == bb.to
}

func (p roadmapper) Adjacent(b dijkstra.Board, totalCost int) dijkstra.AdjacencyIterator {
	bb, ok := b.(*valveNetwork)
	if !ok {
		return dijkstra.DeadEnd()
	}

	adj := []dijkstra.Adj{}
	for _, nb := range bb.Neighbours[string(p)] {
		adj = append(adj, dijkstra.Adj{roadmapper(nb), 1})
	}
	return dijkstra.AdjacencyList(adj)
}
