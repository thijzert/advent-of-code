package dijkstra

import (
	"errors"

	"github.com/thijzert/advent-of-code/lib/pq"
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

	// Pack packs the position into an integer index. Different positions should
	// not pack to the same index
	Pack() int
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
func (hashCodePos) Pack() int                                         { return -1 }

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
		Visited: make([]dijkHead, 250),
		VisMap:  make(map[Position]dijkHead),
	}
	for _, pos := range starts {
		dijk.Heads.Push(pos, 0)
		dijk.setVisited(pos, dijkHead{
			Position:  startingPoint{},
			TotalCost: 0,
		})
	}

	//log.Printf("initial: %v", dijk.Heads)
	for i := 1; i < 10000; i++ {
		dijk.Step()
		//log.Printf("after step %d: %v", i, dijk.Heads)
		if dijk.Heads.Len() == 0 {
			break
		}
		//pos, tc, _ := dijk.Heads.Peek()
		//log.Printf("after step %d: %d heads: %v (%d) %d, ...", i, dijk.Heads.Len(), pos, pos.Pack(), tc)
	}

	if dijk.Shortest.Position == nil {
		return nil, 0, errors.New("Failed to find the shortest path")
	}

	// Create list of steps
	rv := []Position{}
	end := dijk.Shortest.Position
	for end != nil {
		rv = append(rv, end)
		dh, _ := dijk.haveVisited(end)
		end = dh.Position
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
	Heads    pq.PriorityQueue[Position]
	Shortest dijkHead
	Visited  []dijkHead
	VisMap   map[Position]dijkHead
}

func (d *dijkstra) Step() {
	length := d.Heads.Len()

	for i := 0; i < length; i++ {
		position, totalCost, _ := d.Heads.Pop()
		it := position.Adjacent(d.Board, totalCost)
		for {
			n, cost := it.Next()
			if n == nil {
				break
			}

			newCost := totalCost + cost
			if d.Shortest.Position != nil && d.Shortest.TotalCost < newCost {
				// The current partial path is already longer than the shortest path that finishes
				continue
			}

			val, ok := d.haveVisited(n)
			if !ok || val.TotalCost > newCost {
				d.setVisited(n, dijkHead{
					Position:  position,
					TotalCost: newCost,
				})

				d.Heads.Push(n, newCost)
			}

			if n.Final(d.Board) {
				if d.Shortest.Position == nil || d.Shortest.TotalCost > newCost {
					d.Shortest.Position = n
					d.Shortest.TotalCost = newCost
				}
			}
		}
	}
}

type startingPoint struct{}

func (startingPoint) Final(b Board) bool                                { return false }
func (startingPoint) Adjacent(b Board, totalCost int) AdjacencyIterator { return nil }
func (startingPoint) Pack() int                                         { return -1 }

func (d *dijkstra) haveVisited(p Position) (dijkHead, bool) {
	idx := p.Pack()
	rv, ok := dijkHead{}, false
	if idx < 0 {
		rv, ok = d.VisMap[p]
	} else if idx < len(d.Visited) {
		rv = d.Visited[idx]
		ok = rv.Position != nil
	}

	if ok && rv.TotalCost == 0 {
		if _, ok2 := rv.Position.(startingPoint); ok2 {
			rv.Position = nil
		}
	}
	return rv, ok
}

func (d *dijkstra) setVisited(p Position, head dijkHead) {
	idx := p.Pack()
	if idx < 0 {
		d.VisMap[p] = head
		return
	}
	if idx >= len(d.Visited) {
		nl := len(d.Visited) * 2
		for idx >= nl {
			nl *= 2
		}
		nv := make([]dijkHead, nl)
		copy(nv, d.Visited)
		d.Visited = nv
	}
	d.Visited[idx] = head
}
