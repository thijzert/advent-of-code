package dijkstra

import (
	"errors"
	"sort"
)

// A Board represents the space through which we'd like to find short paths.
type Board interface {
	// StartingPositions returns all possible starting positions from which to find the shortest path
	StartingPositions() []Position
}

// A Position encompasses the state of an agent en route. Typically this boils
// down to a grid position, but it can also encode the state of the agent's
// inventory, for instance if the agent is carrying a carrot, rabbit, or fox
// across a bridge. (Or none of the above.)
type Position interface {
	// Final reports whether or not a path can end in this particular
	// configuration.
	Final(b Board) bool

	// Adjacent returns an iterator with all positions that can be reached
	// directly from this position. The total cost of the path that led here is
	// passed in the second parameter
	Adjacent(b Board, totalCost int) AdjacencyIterator
}

// A Hashcoder is a special type of position that exerts influence over how
// it's represented in a hash map.
type Hashcoder interface {
	Position
	Hashcode() [4]uint64
}

type hashCodePos [4]uint64

func (hashCodePos) Final(b Board) bool                                { return false }
func (hashCodePos) Adjacent(b Board, totalCost int) AdjacencyIterator { return nil }

// An AdjacencyIterator is an abstraction over an Adjacency slice that can be
// ranged over. This allows one to generate infinitely many adjacent tiles
// without allocating all the RAM in the known universe.
type AdjacencyIterator interface {
	// Next advances the iterator and returns the next Position in the series,
	// as well as the distance from the previous position. If there are no more
	// Positions, Next should return nil.
	Next() (Position, int)
}

type deadEndIter struct{}

func (deadEndIter) Next() (Position, int) {
	return nil, 0
}

// DeadEnd returns an empty iterator with no neighbouring Positions.
func DeadEnd() AdjacencyIterator {
	return deadEndIter{}
}

// An Adj represents an adjacent Position along with its step size
type Adj struct {
	Position Position
	Cost     int
}

type adjList struct {
	Positions []Adj
	Idx       int
}

func (al *adjList) Next() (Position, int) {
	if al.Idx >= len(al.Positions) {
		return nil, 0
	}
	rv := al.Positions[al.Idx]
	al.Idx++
	return rv.Position, rv.Cost
}

// AdjacencyList turns a list of adjacent positions into an AdjacencyIterator
func AdjacencyList(positions []Adj) AdjacencyIterator {
	return &adjList{
		Positions: positions,
		Idx:       0,
	}
}

// ShortestPath calculates the shortest path (i.e., the path with the lowest
// overall cost) from any of the Board's starting positions to any Positions
// that are Final. It returns the list of steps taken, along with the total
// cost for this path, or an error.
func ShortestPath(b Board) ([]Position, int, error) {
	starts := b.StartingPositions()

	dijk := &dijkstra{
		Board:   b,
		Visited: make(map[Position]dijkHead),
	}
	for _, pos := range starts {
		nn := pos
		if hc, ok := pos.(Hashcoder); ok {
			nn = hashCodePos(hc.Hashcode())
		}
		dijk.Heads = append(dijk.Heads, dijkHead{
			Position:  pos,
			TotalCost: 0,
		})
		dijk.Visited[nn] = dijkHead{
			Position:  nil,
			TotalCost: 0,
		}
	}

	//log.Printf("initial: %v", dijk.Heads)
	for i := 1; i < 10000; i++ {
		dijk.Step()
		//log.Printf("after step %d: %v", i, dijk.Heads)
		if len(dijk.Heads) == 0 {
			break
		}
		//log.Printf("after step %d: %d heads: %v...", i, len(dijk.Heads), dijk.Heads[0])
	}

	if dijk.Shortest.Position == nil {
		return nil, 0, errors.New("Failed to find the shortest path")
	}

	// Create list of steps
	rv := []Position{}
	end := dijk.Shortest.Position
	for end != nil {
		rv = append(rv, end)
		nn := end
		if hc, ok := end.(Hashcoder); ok {
			nn = hashCodePos(hc.Hashcode())
		}
		end = dijk.Visited[nn].Position
	}
	l := len(rv)
	for i := range rv[:l/2] {
		rv[i], rv[l-i-1] = rv[l-i-1], rv[i]
	}

	return rv, dijk.Shortest.TotalCost, nil
}

type dijkHead struct {
	Position  Position
	TotalCost int
}

type dijkstra struct {
	Board    Board
	Heads    []dijkHead
	Shortest dijkHead
	Visited  map[Position]dijkHead
}

func (d *dijkstra) Step() {
	length := len(d.Heads)
	CUTOFF := 5000
	if length > CUTOFF {
		sort.Slice(d.Heads, func(i, j int) bool {
			return d.Heads[i].TotalCost < d.Heads[j].TotalCost
		})
		length = CUTOFF
	}

	newHeads := []dijkHead{}
	for i, h := range d.Heads[:length] {
		it := h.Position.Adjacent(d.Board, h.TotalCost)
		first := true
		for {
			n, cost := it.Next()
			if n == nil {
				break
			}

			newCost := h.TotalCost + cost
			if d.Shortest.Position != nil && d.Shortest.TotalCost < newCost {
				// The current partial path is already longer than the shortest path that finishes
				continue
			}

			nn := n
			if hc, ok := n.(Hashcoder); ok {
				nn = hashCodePos(hc.Hashcode())
			}
			val, ok := d.Visited[nn]
			if !ok || val.TotalCost > newCost {
				val = dijkHead{
					Position:  h.Position,
					TotalCost: newCost,
				}
				d.Visited[nn] = val

				if first {
					d.Heads[i].Position = n
					d.Heads[i].TotalCost = newCost
					first = false
				} else {
					newHeads = append(newHeads, dijkHead{n, newCost})
				}
			}

			if n.Final(d.Board) {
				if d.Shortest.Position == nil || d.Shortest.TotalCost > newCost {
					d.Shortest.Position = n
					d.Shortest.TotalCost = newCost
				}
			}
		}
		if first {
			d.Heads[i].Position = nil
		}
	}

	if d.Shortest.Position != nil {
		// Prune paths that are longer than the shortest finishing path, but were below the cutoff
		pruned := 0
		for i, h := range d.Heads[length:] {
			if h.TotalCost > d.Shortest.TotalCost {
				d.Heads[length+i].Position = nil
				pruned++
			}
		}
		if pruned > 0 {
			//log.Printf("%d paths pruned; shortest path is %d", pruned, d.Shortest.TotalCost)
		}
	}

	for i := len(d.Heads) - 1; i >= 0; i-- {
		if d.Heads[i].Position == nil {
			d.Heads = append(d.Heads[:i], d.Heads[i+1:]...)
		}
	}
	d.Heads = append(d.Heads, newHeads...)
}
