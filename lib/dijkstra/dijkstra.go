package dijkstra

import (
	"errors"
	"log"
)

type Board interface {
	// StartingPositions returns all possible starting positions from which to find the shortest path
	StartingPositions() []Position
}

type Position interface {
	Final(b Board, totalCost int) bool
	Adjacent(b Board) AdjacencyIterator
}

type AdjacencyIterator interface {
	Next() (Position, int)
}

type deadEndIter struct{}

func (deadEndIter) Next() (Position, int) {
	return nil, 0
}

func DeadEnd() AdjacencyIterator {
	return deadEndIter{}
}

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

func AdjacencyList(positions []Adj) AdjacencyIterator {
	return &adjList{
		Positions: positions,
		Idx:       0,
	}
}

type ChattyBoi interface {
	YesIAmVeryVerboseThankYouForAsking()
}

func ShortestPath(b Board) ([]Position, int, error) {
	starts := b.StartingPositions()

	dijk := &dijkstra{
		Board:   b,
		Visited: make(map[Position]dijkHead),
	}
	for _, pos := range starts {
		dijk.Heads = append(dijk.Heads, dijkHead{
			Position:  pos,
			TotalCost: 0,
		})
		dijk.Visited[pos] = dijkHead{
			Position:  nil,
			TotalCost: 0,
		}
	}

	if _, ok := b.(ChattyBoi); ok {
		log.Printf("Step 0; heads: %v", dijk.Heads)
	}
	for i := 1; i < 10000; i++ {
		dijk.Step()
		if _, ok := b.(ChattyBoi); ok {
			log.Printf("Step %d; heads: %v", i, dijk.Heads)
		}
		if len(dijk.Heads) == 0 {
			break
		}
	}

	if dijk.Shortest.Position == nil {
		return nil, 0, errors.New("Failed to find the shortest path")
	}

	// Create list of steps
	rv := []Position{}
	end := dijk.Shortest.Position
	for end != nil {
		rv = append(rv, end)
		end = dijk.Visited[end].Position
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
	newHeads := []dijkHead{}

	for i, h := range d.Heads {
		it := h.Position.Adjacent(d.Board)
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

			val, ok := d.Visited[n]
			if !ok || val.TotalCost > newCost {
				val = dijkHead{
					Position:  h.Position,
					TotalCost: newCost,
				}
				d.Visited[n] = val

				if first {
					d.Heads[i].Position = n
					d.Heads[i].TotalCost = newCost
					first = false
				} else {
					newHeads = append(newHeads, dijkHead{n, newCost})
				}
			}

			if n.Final(d.Board, newCost) {
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

	for i := len(d.Heads) - 1; i >= 0; i-- {
		if d.Heads[i].Position == nil {
			d.Heads = append(d.Heads[:i], d.Heads[i+1:]...)
		}
	}
	d.Heads = append(d.Heads, newHeads...)
}
