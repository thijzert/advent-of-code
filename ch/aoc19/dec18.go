package aoc19

import (
	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
	"github.com/thijzert/advent-of-code/lib/dijkstra"
)

func Dec18a(ctx ch.AOContext) error {
	sections, err := ctx.DataSections("inputs/2019/dec18.txt")
	if err != nil {
		return err
	}

	for _, lines := range sections {
		tm := readTractorMaze(lines)

		steps, cost, err := dijkstra.ShortestPath(tm)
		if err != nil {
			return err
		}
		path := make([]byte, 0, len(steps)-1)
		var prev dijkstra.Position = nil
		for _, p := range steps {
			if prev != nil {
				c := keyDiff(prev, p)
				if c != '-' {
					path = append(path, c)
				}
			}
			prev = p
		}
		ctx.Printf("Path taken: %s", path)
		ctx.FinalAnswer.Print(cost)
	}
	return nil
}

type tractorMaze struct {
	Lines          []string
	KeyRequirement uint32
	d2kCache       map[pos2d]map[cube.Point]int
}

func readTractorMaze(lines []string) tractorMaze {
	tm := tractorMaze{
		Lines:    lines,
		d2kCache: make(map[pos2d]map[cube.Point]int),
	}
	for _, l := range lines {
		for _, c := range l {
			if c >= 'a' && c <= 'z' {
				tm.KeyRequirement = tm.KeyRequirement | (1 << int(c-'a'))
			}
		}
	}
	return tm
}

func (b tractorMaze) StartingPositions() []dijkstra.Position {
	rv := []dijkstra.Position{}
	for y, l := range b.Lines {
		for x, c := range l {
			if c == '@' {
				rv = append(rv, pos2d{x, y, 0})
			} else if c == '%' {
				pos := pos4d{
					Robots: [4]cube.Point{
						{x - 1, y - 1},
						{x + 1, y - 1},
						{x - 1, y + 1},
						{x + 1, y + 1},
					},
					Keys: 0,
				}
				rv = append(rv, pos.WithUpdatedMask(b))
			}
		}
	}
	return rv
}

func (b tractorMaze) charAt(x, y int) rune {
	if y >= 0 && y < len(b.Lines) {
		if x >= 0 && x < len(b.Lines[y]) {
			return rune(b.Lines[y][x])
		}
	}
	return '#'
}

func (b tractorMaze) updatedKeys(x, y int, currentKeys uint32) uint32 {
	c := b.charAt(x, y)
	if c >= 'a' && c <= 'z' {
		currentKeys = currentKeys | (1 << int(c-'a'))
	}
	return currentKeys
}
func (b tractorMaze) posAt(x, y int, currentKeys uint32) pos2d {
	return pos2d{x, y, b.updatedKeys(x, y, currentKeys)}
}

func (bb tractorMaze) distanceToKeys(pos pos2d) map[cube.Point]int {
	if rv, ok := bb.d2kCache[pos]; ok {
		return rv
	}

	// This map contains the distance to all reachable keys.
	// We fill it from the final() function
	keydist := make(map[cube.Point]int)

	valid := func(x0, y0, x, y, totalCost int) bool {
		c := bb.charAt(x, y)
		if c == '#' || c == '%' {
			return false
		} else if c >= 'A' && c <= 'Z' {
			return (pos.Keys>>int(c-'A'))&1 == 1
		} else if c >= 'a' && c <= 'z' {
			if ((pos.Keys >> int(c-'a')) & 1) == 1 {
				// We already have this key
				return true
			}
			cp := cube.Point{x, y}
			if d, ok := keydist[cp]; !ok || d > totalCost {
				keydist[cp] = totalCost
			}
			// This key is new, so don't continue. We can't return a dead end,
			// so let's just pretend this tile wasn't valid either
			return false
		}
		return true
	}
	final := func(x, y int) bool {
		return false
	}
	dijkstra.ShortestPath(dijkstra.GridWalker(valid, final, pos.X, pos.Y))

	bb.d2kCache[pos] = keydist
	return keydist
}

type Keyer interface {
	GetKeys() uint32
}

func keyDiff(a, b dijkstra.Position) byte {
	var p, q Keyer
	var ok bool
	if p, ok = a.(Keyer); !ok {
		return ' '
	}
	if q, ok = b.(Keyer); !ok {
		return ' '
	}

	diff := p.GetKeys() ^ q.GetKeys()
	if diff == 0 {
		return '-'
	}
	for c := byte('a'); c <= 'z'; c++ {
		if (diff >> int(c-'a')) == 1 {
			return c
		}
	}
	return '?'
}

type pos2d struct {
	X, Y int
	Keys uint32
}

func (p pos2d) GetKeys() uint32 {
	return p.Keys
}

func (p pos2d) Final(b dijkstra.Board) bool {
	bb, ok := b.(tractorMaze)
	if !ok {
		return false
	}

	return p.Keys == bb.KeyRequirement
}
func (pos pos2d) Adjacent(b dijkstra.Board, totalCost int) dijkstra.AdjacencyIterator {
	bb, ok := b.(tractorMaze)
	if !ok {
		return nil
	}

	return &pos2diter{
		pos: pos,
		b:   bb,
		idx: 0,
	}
}

type pos2diter struct {
	pos pos2d
	b   tractorMaze
	idx int
}

func (pdi *pos2diter) Next() (dijkstra.Position, int) {
	pdi.idx++
	found := 0

	rv := pdi.b.posAt(pdi.pos.X+1, pdi.pos.Y, pdi.pos.Keys)
	if pdi.isOk(rv) {
		found++
		if found == pdi.idx {
			return rv, 1
		}
	}

	rv = pdi.b.posAt(pdi.pos.X, pdi.pos.Y+1, pdi.pos.Keys)
	if pdi.isOk(rv) {
		found++
		if found == pdi.idx {
			return rv, 1
		}
	}

	rv = pdi.b.posAt(pdi.pos.X-1, pdi.pos.Y, pdi.pos.Keys)
	if pdi.isOk(rv) {
		found++
		if found == pdi.idx {
			return rv, 1
		}
	}

	rv = pdi.b.posAt(pdi.pos.X, pdi.pos.Y-1, pdi.pos.Keys)
	if pdi.isOk(rv) {
		found++
		if found == pdi.idx {
			return rv, 1
		}
	}

	return nil, 0
}

func (pdi *pos2diter) isOk(p pos2d) bool {
	c := pdi.b.charAt(p.X, p.Y)
	if c == '#' {
		return false
	}

	if c >= 'A' && c <= 'Z' {
		return (p.Keys>>int(c-'A'))&1 == 1
	}

	return true
}

func Dec18b(ctx ch.AOContext) error {
	sections, err := ctx.DataSections("inputs/2019/dec18.txt")
	if err != nil {
		return err
	}
	for _, lines := range sections {

		tm := readTractorMaze(lines)

		// Patch the maze to split it into four chunks
		found := false
		for y, line := range tm.Lines {
			if found {
				break
			}
			for x, c := range line {
				if c == '@' {
					found = true
					tm.Lines[y-1] = tm.Lines[y-1][:x-1] + "*#*" + tm.Lines[y-1][x+2:]
					tm.Lines[y+0] = tm.Lines[y+0][:x-1] + "#%#" + tm.Lines[y+0][x+2:]
					tm.Lines[y+1] = tm.Lines[y+1][:x-1] + "*#*" + tm.Lines[y+1][x+2:]
					break
				}
			}
		}

		//ctx.Printf("\n%s\n", strings.Join(tm.Lines, "\n"))

		_, cost, err := dijkstra.ShortestPath(tm)
		if err != nil {
			return err
		}
		ctx.FinalAnswer.Print(cost)
	}

	return nil
}

type pos4d struct {
	Robots [4]cube.Point
	// KeyDoorMask contains a bit mask for all doors and keys a particular robot can potentially encounter.
	// Any state changes not in this mask will be of no consequence to possible next steps for this robot, and can be ignored.
	KeyDoorMask [4]uint32
	Keys        uint32
}

func (p pos4d) GetKeys() uint32 {
	return p.Keys
}

func (p pos4d) Final(b dijkstra.Board) bool {
	bb, ok := b.(tractorMaze)
	if !ok {
		return false
	}

	return p.Keys == bb.KeyRequirement
}
func (p pos4d) Adjacent(b dijkstra.Board, totalCost int) dijkstra.AdjacencyIterator {
	bb, ok := b.(tractorMaze)
	if !ok {
		return nil
	}

	var rv []dijkstra.Adj

	for i, pos := range p.Robots {
		keydist := bb.distanceToKeys(pos2d{pos.X, pos.Y, p.Keys & p.KeyDoorMask[i]})
		for q, tc := range keydist {
			var np pos4d = p
			np.Robots[i] = q
			np.Keys = bb.updatedKeys(q.X, q.Y, p.Keys)
			rv = append(rv, dijkstra.Adj{
				Position: np,
				Cost:     tc,
			})
		}
	}

	return dijkstra.AdjacencyList(rv)
}

// WithUpdatedMask returns a copy of p with the key/door masks set to the correct value for this board
func (p pos4d) WithUpdatedMask(b tractorMaze) pos4d {
	for i, pos := range p.Robots {
		p.KeyDoorMask[i] = 0
		valid := func(x0, y0, x, y, totalCost int) bool {
			c := b.charAt(x, y)
			return c != '#' && c != '%'
		}
		final := func(x, y int) bool {
			c := b.charAt(x, y)
			if c >= 'A' && c <= 'Z' {
				p.KeyDoorMask[i] = p.KeyDoorMask[i] | (1 << int(c-'A'))
			} else if c >= 'a' && c <= 'z' {
				p.KeyDoorMask[i] = p.KeyDoorMask[i] | (1 << int(c-'a'))
			}
			return false
		}
		dijkstra.ShortestPath(dijkstra.GridWalker(valid, final, pos.X, pos.Y))
	}
	return p
}
