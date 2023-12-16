package aoc23

import (
	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
	"github.com/thijzert/advent-of-code/lib/image"
)

type Beamfron struct {
	Pos, Dir cube.Point
}

func Dec16a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec16.txt")
	if err != nil {
		return nil, err
	}
	img := image.NewImage(len(lines[0]), len(lines), func(_, _ int) int {
		return 0
	})

	type bfm struct {
		Dir  cube.Point
		Char byte
	}
	mirrors := make(map[bfm][]cube.Point)
	right, up, left, down, stopped := cube.Cardinal2D[0], cube.Cardinal2D[3], cube.Cardinal2D[2], cube.Cardinal2D[1], cube.Point{0, 0}
	mirrors[bfm{right, '.'}] = []cube.Point{right}
	mirrors[bfm{up, '.'}] = []cube.Point{up}
	mirrors[bfm{left, '.'}] = []cube.Point{left}
	mirrors[bfm{down, '.'}] = []cube.Point{down}
	mirrors[bfm{right, '\\'}] = []cube.Point{down}
	mirrors[bfm{up, '\\'}] = []cube.Point{left}
	mirrors[bfm{left, '\\'}] = []cube.Point{up}
	mirrors[bfm{down, '\\'}] = []cube.Point{right}
	mirrors[bfm{right, '/'}] = []cube.Point{up}
	mirrors[bfm{up, '/'}] = []cube.Point{right}
	mirrors[bfm{left, '/'}] = []cube.Point{down}
	mirrors[bfm{down, '/'}] = []cube.Point{left}
	mirrors[bfm{right, '-'}] = []cube.Point{right}
	mirrors[bfm{up, '-'}] = []cube.Point{left, right}
	mirrors[bfm{left, '-'}] = []cube.Point{left}
	mirrors[bfm{down, '-'}] = []cube.Point{left, right}
	mirrors[bfm{right, '|'}] = []cube.Point{up, down}
	mirrors[bfm{up, '|'}] = []cube.Point{up}
	mirrors[bfm{left, '|'}] = []cube.Point{up, down}
	mirrors[bfm{down, '|'}] = []cube.Point{down}

	beams := []Beamfron{Beamfron{cube.Point{-1, 0}, right}}
	newBeams := []Beamfron{}
	moving := 1
	for moving > 0 {
		moving = 0
		for i, beam := range beams {
			if beam.Dir == stopped {
				continue
			}
			seen := img.At(beam.Pos.X, beam.Pos.Y)
			dirmask := 1 << (3*(beam.Dir.X+1) + beam.Dir.Y + 1)
			if seen&dirmask != 0 {
				beam.Dir = stopped
				beams[i] = beam
				continue
			}
			img.Set(beam.Pos.X, beam.Pos.Y, seen|dirmask)

			beam.Pos = beam.Pos.Add(beam.Dir)
			if !img.Inside(beam.Pos.X, beam.Pos.Y) {
				beam.Dir = stopped
				beams[i] = beam
				continue
			}
			moving++

			dirs := mirrors[bfm{beam.Dir, lines[beam.Pos.Y][beam.Pos.X]}]
			if len(dirs) == 0 {
				ctx.Printf("%d %d: '%c'", beam.Pos, beam.Dir, lines[beam.Pos.Y][beam.Pos.X])
				return nil, errFailed
			}
			beam.Dir = dirs[0]
			beams[i] = beam
			for _, d := range dirs[1:] {
				newBeams = append(newBeams, Beamfron{beam.Pos, d})
			}
		}
		beams = append(beams, newBeams...)
		newBeams = newBeams[:0]
	}

	ctx.Printf("energised:\n%s", img)
	answer := 0
	for y := range lines {
		for x := range lines[0] {
			if img.At(x, y) != 0 {
				answer++
			}
		}
	}
	return answer, nil
}

var Dec16b ch.AdventFunc = nil

// func Dec16b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }
