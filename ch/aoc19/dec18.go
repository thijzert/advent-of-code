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
				path = append(path, keyDiff(prev, p))
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

func (tractorMaze) NoYesIAmVeryVerboseThankYouForAsking() {}

func (b tractorMaze) StartingPositions() []dijkstra.Position {
	rv := []dijkstra.Position{}
	for y, l := range b.Lines {
		for x, c := range l {
			if c == '@' {
				rv = append(rv, pos2d{x, y, 0})
			} else if c == '%' {
				rv = append(rv, pos4d{int8(x + 1), int8(y + 1), int8(x + 1), int8(y - 1), int8(x - 1), int8(y + 1), int8(x - 1), int8(y - 1), 0})

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

func (b tractorMaze) updatedKeys8(x, y int8, currentKeys uint32) uint32 {
	return b.updatedKeys(int(x), int(y), currentKeys)
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

	keydist := make(map[cube.Point]int)
	f := func(p gridPoint, totalCost int) bool {
		c := bb.charAt(p.X, p.Y)
		if c < 'a' || c > 'z' {
			return false
		} else if ((pos.Keys >> int(c-'a')) & 1) == 1 {
			// We already have this key
			return false
		}
		cp := cube.Point{p.X, p.Y}
		if d, ok := keydist[cp]; !ok || d > totalCost {
			keydist[cp] = totalCost
		}
		return false
	}
	bff := BFF{f}
	mtm := metaTractorMaze{
		b:     bb,
		start: gridPoint{pos.X, pos.Y, &bff},
		keys:  pos.Keys,
	}
	dijkstra.ShortestPath(mtm)

	bb.d2kCache[pos] = keydist
	return keydist
}

type metaTractorMaze struct {
	b     tractorMaze
	start gridPoint
	keys  uint32
}

func (b metaTractorMaze) StartingPositions() []dijkstra.Position {
	return []dijkstra.Position{b.start}
}

type BFF struct {
	F func(gridPoint, int) bool
}

type gridPoint struct {
	X, Y int
	F    *BFF
}

func (p gridPoint) HashCode() uint64 {
	return (uint64(uint32(p.X)) << 32) | uint64(uint32(p.Y))
}

func (p gridPoint) Final(b dijkstra.Board, totalCost int) bool {
	return p.F.F(p, totalCost)
}
func (p gridPoint) Adjacent(b dijkstra.Board) dijkstra.AdjacencyIterator {
	bb, ok := b.(metaTractorMaze)
	if !ok {
		return nil
	}

	c := bb.b.charAt(p.X, p.Y)
	if c >= 'a' && c <= 'z' {
		if (bb.keys>>int(c-'a'))&1 != 1 {
			// We don't have this key yet
			return dijkstra.DeadEnd()
		}
	}

	var rv []dijkstra.Adj
	positions := [4]gridPoint{
		gridPoint{p.X + 1, p.Y, p.F},
		gridPoint{p.X, p.Y + 1, p.F},
		gridPoint{p.X - 1, p.Y, p.F},
		gridPoint{p.X, p.Y - 1, p.F},
	}
	for _, pos := range positions {
		c := bb.b.charAt(pos.X, pos.Y)
		if c == '#' || c == '%' {
			continue
		}

		if c >= 'A' && c <= 'Z' {
			if (bb.keys>>int(c-'A'))&1 != 1 {
				continue
			}
		}
		rv = append(rv, dijkstra.Adj{pos, 1})
	}

	return dijkstra.AdjacencyList(rv)
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

func (p pos2d) Final(b dijkstra.Board, totalCost int) bool {
	bb, ok := b.(tractorMaze)
	if !ok {
		return false
	}

	return p.Keys == bb.KeyRequirement
}
func (pos pos2d) Adjacent(b dijkstra.Board) dijkstra.AdjacencyIterator {
	bb, ok := b.(tractorMaze)
	if !ok {
		return nil
	}

	var rv []dijkstra.Adj

	keydist := bb.distanceToKeys(pos)
	for q, tc := range keydist {
		rv = append(rv, dijkstra.Adj{
			Position: bb.posAt(q.X, q.Y, pos.Keys),
			Cost:     tc,
		})
	}

	return dijkstra.AdjacencyList(rv)
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

	return errNotImplemented
	return nil
}

type pos4d struct {
	X1, Y1 int8
	X2, Y2 int8
	X3, Y3 int8
	X4, Y4 int8
	Keys   uint32
}

func (p pos4d) GetKeys() uint32 {
	return p.Keys
}

func (p pos4d) Final(b dijkstra.Board, totalCost int) bool {
	bb, ok := b.(tractorMaze)
	if !ok {
		return false
	}

	return p.Keys == bb.KeyRequirement
}
func (p pos4d) Adjacent(b dijkstra.Board) dijkstra.AdjacencyIterator {
	bb, ok := b.(tractorMaze)
	if !ok {
		return nil
	}

	return &pos4diter{
		positions: [16]pos4d{
			pos4d{p.X1 + 1, p.Y1, p.X2, p.Y2, p.X3, p.Y3, p.X4, p.Y4, bb.updatedKeys8(p.X1+1, p.Y1, p.Keys)},
			pos4d{p.X1, p.Y1 + 1, p.X2, p.Y2, p.X3, p.Y3, p.X4, p.Y4, bb.updatedKeys8(p.X1, p.Y1+1, p.Keys)},
			pos4d{p.X1 - 1, p.Y1, p.X2, p.Y2, p.X3, p.Y3, p.X4, p.Y4, bb.updatedKeys8(p.X1-1, p.Y1, p.Keys)},
			pos4d{p.X1, p.Y1 - 1, p.X2, p.Y2, p.X3, p.Y3, p.X4, p.Y4, bb.updatedKeys8(p.X1, p.Y1-1, p.Keys)},
			pos4d{p.X1, p.Y1, p.X2 + 1, p.Y2, p.X3, p.Y3, p.X4, p.Y4, bb.updatedKeys8(p.X2+1, p.Y2, p.Keys)},
			pos4d{p.X1, p.Y1, p.X2, p.Y2 + 1, p.X3, p.Y3, p.X4, p.Y4, bb.updatedKeys8(p.X2, p.Y2+1, p.Keys)},
			pos4d{p.X1, p.Y1, p.X2 - 1, p.Y2, p.X3, p.Y3, p.X4, p.Y4, bb.updatedKeys8(p.X2-1, p.Y2, p.Keys)},
			pos4d{p.X1, p.Y1, p.X2, p.Y2 - 1, p.X3, p.Y3, p.X4, p.Y4, bb.updatedKeys8(p.X2, p.Y2-1, p.Keys)},
			pos4d{p.X1, p.Y1, p.X2, p.Y2, p.X3 + 1, p.Y3, p.X4, p.Y4, bb.updatedKeys8(p.X3+1, p.Y3, p.Keys)},
			pos4d{p.X1, p.Y1, p.X2, p.Y2, p.X3, p.Y3 + 1, p.X4, p.Y4, bb.updatedKeys8(p.X3, p.Y3+1, p.Keys)},
			pos4d{p.X1, p.Y1, p.X2, p.Y2, p.X3 - 1, p.Y3, p.X4, p.Y4, bb.updatedKeys8(p.X3-1, p.Y3, p.Keys)},
			pos4d{p.X1, p.Y1, p.X2, p.Y2, p.X3, p.Y3 - 1, p.X4, p.Y4, bb.updatedKeys8(p.X3, p.Y3-1, p.Keys)},
			pos4d{p.X1, p.Y1, p.X2, p.Y2, p.X3, p.Y3, p.X4 + 1, p.Y4, bb.updatedKeys8(p.X4+1, p.Y4, p.Keys)},
			pos4d{p.X1, p.Y1, p.X2, p.Y2, p.X3, p.Y3, p.X4, p.Y4 + 1, bb.updatedKeys8(p.X4, p.Y4+1, p.Keys)},
			pos4d{p.X1, p.Y1, p.X2, p.Y2, p.X3, p.Y3, p.X4 - 1, p.Y4, bb.updatedKeys8(p.X4-1, p.Y4, p.Keys)},
			pos4d{p.X1, p.Y1, p.X2, p.Y2, p.X3, p.Y3, p.X4, p.Y4 - 1, bb.updatedKeys8(p.X4, p.Y4-1, p.Keys)},
		},
		idx: 0,
		b:   bb,
	}
}

type pos4diter struct {
	positions [16]pos4d
	idx       int
	b         tractorMaze
}

func (pdi *pos4diter) Next() (dijkstra.Position, int) {
	for pdi.idx < len(pdi.positions) {
		if pdi.isOk(pdi.positions[pdi.idx]) {
			rv := pdi.positions[pdi.idx]
			pdi.idx++
			return rv, 1
		}
		pdi.idx++
	}

	return nil, 0
}

func (pdi *pos4diter) isOk(p pos4d) bool {
	toCheck := [4][2]int{
		{int(p.X1), int(p.Y1)},
		{int(p.X2), int(p.Y2)},
		{int(p.X3), int(p.Y3)},
		{int(p.X4), int(p.Y4)},
	}
	for _, pos := range toCheck {
		c := pdi.b.charAt(pos[0], pos[1])
		if c == '#' || c == '%' {
			return false
		}

		if c >= 'A' && c <= 'Z' {
			if (p.Keys>>int(c-'A'))&1 != 1 {
				return false
			}
		}
	}

	return true
}
