package aoc19

import (
	"strings"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/dijkstra"
)

var Dec18b ch.AdventFunc = nil

func Dec18a(ctx ch.AOContext) error {
	sections, err := ctx.DataSections("inputs/2019/dec18.txt")
	if err != nil {
		return err
	}

	for _, lines := range sections {
		tm := tractorMaze{Lines: lines}
		kreq := make([]byte, 26)
		kl := 0
		for _, l := range lines {
			for _, c := range l {
				if c >= 'a' && c <= 'z' {
					cc := int(c - 'a')
					if kl <= cc {
						kl = cc + 1
					}
					kreq[cc] = byte(c)
				}
			}
		}
		tm.KeyRequirement = string(kreq[:kl])

		ctx.Printf("[%s]", tm.KeyRequirement)

		_, cost, err := dijkstra.ShortestPath(tm)
		if err != nil {
			return err
		}
		ctx.FinalAnswer.Print(cost)
	}
	return nil
}

type tractorMaze struct {
	Lines          []string
	KeyRequirement string
}

func (b tractorMaze) StartingPositions() []dijkstra.Position {
	defaultKeys := strings.Repeat(" ", len(b.KeyRequirement))
	rv := []dijkstra.Position{}
	for y, l := range b.Lines {
		for x, c := range l {
			if c == '@' {
				rv = append(rv, pos2d{x, y, defaultKeys})
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

func (b tractorMaze) posAt(x, y int, currentKeys string) pos2d {
	c := b.charAt(x, y)
	if c >= 'a' && c <= 'z' {
		i := int(c - 'a')
		if currentKeys[i] != byte(c) {
			currentKeys = currentKeys[:i] + string(c) + currentKeys[i+1:]
		}
	}
	return pos2d{x, y, currentKeys}
}

type pos2d struct {
	X, Y int
	Keys string
}

func (p pos2d) Final(b dijkstra.Board) bool {
	bb, ok := b.(tractorMaze)
	if !ok {
		return false
	}

	return p.Keys == bb.KeyRequirement
}
func (p pos2d) Adjacent(b dijkstra.Board) dijkstra.AdjacencyIterator {
	bb, ok := b.(tractorMaze)
	if !ok {
		return nil
	}

	return &pos2diter{
		pos: p,
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
		return p.Keys[int(c-'A')] == byte(c+' ')
	}

	return true
}

// func Dec18b(ctx ch.AOContext) error {
// 	return errNotImplemented
// }
