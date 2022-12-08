package dijkstra

import (
	"testing"
)

func TestRiverCrossingProblem(t *testing.T) {
	steps, cost, err := ShortestPath(rivercrossing{})
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("Found a solution in %d steps", cost)

	for _, st := range steps {
		t.Logf("   - %s", st)
	}

	if cost != 7 {
		t.Fail()
	}
}

type riverbank int

type rivercrossing struct {
	Boatman riverbank
	Carrot  riverbank
	Rabbit  riverbank
	Fox     riverbank
}

func (rivercrossing) StartingPositions() []Position {
	return []Position{rivercrossing{0, 0, 0, 0}}
}

func (p rivercrossing) String() string {
	rv := ""
	if p.Carrot == 0 {
		rv += "\U0001f955"
	} else {
		rv += "  "
	}
	if p.Rabbit == 0 {
		rv += "\U0001f407"
	} else {
		rv += "  "
	}
	if p.Fox == 0 {
		rv += "\U0001f98a"
	} else {
		rv += "  "
	}
	if p.Boatman == 0 {
		rv += "\U0001f6a3"
	} else {
		rv += "  "
	}

	rv += " \U0001f30a\U0001f30a\U0001f30a "

	if p.Boatman == 1 {
		rv += "\U0001f6a3"
	} else {
		rv += "  "
	}
	if p.Fox == 1 {
		rv += "\U0001f98a"
	} else {
		rv += "  "
	}
	if p.Rabbit == 1 {
		rv += "\U0001f407"
	} else {
		rv += "  "
	}
	if p.Carrot == 1 {
		rv += "\U0001f955"
	} else {
		rv += "  "
	}

	return rv
}

func (p rivercrossing) Final(b Board) bool {
	return p == rivercrossing{1, 1, 1, 1}
}

func (p rivercrossing) Eek() bool {
	if p.Carrot == p.Rabbit {
		return p.Boatman != p.Carrot
	}
	if p.Rabbit == p.Fox {
		return p.Boatman != p.Fox
	}
	return false
}

func (p rivercrossing) Adjacent(b Board, totalCost int) AdjacencyIterator {
	pos := []rivercrossing{
		rivercrossing{1 - p.Boatman, p.Carrot, p.Rabbit, p.Fox},
	}
	if p.Boatman == p.Carrot {
		pos = append(pos, rivercrossing{1 - p.Boatman, 1 - p.Carrot, p.Rabbit, p.Fox})
	}
	if p.Boatman == p.Rabbit {
		pos = append(pos, rivercrossing{1 - p.Boatman, p.Carrot, 1 - p.Rabbit, p.Fox})
	}
	if p.Boatman == p.Fox {
		pos = append(pos, rivercrossing{1 - p.Boatman, p.Carrot, p.Rabbit, 1 - p.Fox})
	}

	var rv []Adj
	for _, q := range pos {
		if !q.Eek() {
			rv = append(rv, Adj{q, 1})
		}
	}
	return AdjacencyList(rv)
}
