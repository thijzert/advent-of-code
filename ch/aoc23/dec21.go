package aoc23

import (
	"fmt"
	"os"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
	"github.com/thijzert/advent-of-code/lib/image"
)

func Dec21a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec21.txt")
	if err != nil {
		return nil, err
	}
	start := []cube.Point{}
	for y, line := range lines {
		for x, c := range line {
			if c == 'S' {
				start = append(start, cube.Point{x, y})
			}
		}
	}
	answer, double, img := dec21BFSflood(lines, start, 64)
	ctx.Printf("Garden: (%d)\n%s", double, img)

	return answer, nil
}
func dec21BFSflood(lines []string, start []cube.Point, steps int) (int, int, *image.Image) {
	img := image.ReadImage(lines, image.Octothorpe)
	img.Default = -1

	reach := make(map[cube.Point]bool)
	for _, pt := range start {
		reach[pt] = true
		img.Set(pt.X, pt.Y, 2)
	}

	for step := 0; step < steps; step++ {
		newReach := make(map[cube.Point]bool)
		for pos := range reach {
			img.Set(pos.X, pos.Y, 0)
			for _, add := range cube.Cardinal2D {
				np := pos.Add(add)
				if img.At(np.X, np.Y) == 0 {
					newReach[cube.Point{np.X, np.Y}] = true
					img.Set(np.X, np.Y, 2)
				}
			}
		}
		reach = newReach
	}

	// Count the number of lit cells in the top row and left column
	// These are counted double when tiling the garden
	topleft := 0
	size := len(lines) - 1
	for i := 0; i <= size; i++ {
		topleft += img.At(i, 0) / 2
		topleft += img.At(0, i) / 2
	}
	topleft -= img.At(0, 0) / 2

	img.Default = 0
	return len(reach) - topleft, topleft, img
}

func dec21bGenerateTestData(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec21.txt")
	if err != nil {
		return nil, err
	}
	for i, line := range lines {
		lines[i] = "." + strings.Repeat(line[1:], 9)
	}
	newLines := []string{lines[0]}
	for i := 0; i < 9; i++ {
		newLines = append(newLines, lines[1:]...)
	}
	lines = newLines
	size := len(lines) - 1

	// Assumption: the start point is in the exact centre
	start := cube.Point{X: size / 2, Y: size / 2}
	if lines[start.Y][start.X] != 'S' {
		return nil, errFailed
	}

	// 501 steps with actual input: 217868
	// 69 steps with synthetic 19x19 input: 3964
	const STEPS = 501

	answer, _, img := dec21BFSflood(lines, []cube.Point{start}, STEPS)
	ctx.Printf("With %d steps: %d", STEPS, answer)
	for x := 0; x < img.Width; x++ {
		if img.At(x, 0) != 0 {
			return nil, errFailed
		}
	}
	//ctx.Printf("Garden:\n%s", img)
	row := make([]byte, img.Width+1)
	row[img.Width] = '\n'
	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			row[x] = ".o#"[img.At(x, y)]
		}
		os.Stdout.Write(row)
	}

	return answer, nil
}

func Dec21b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec21.txt")
	if err != nil {
		return nil, err
	}
	size := len(lines) - 1

	// Assumption: the start point is in the exact centre
	start := cube.Point{size / 2, size / 2}
	if lines[start.Y][start.X] != 'S' {
		return nil, errFailed
	}

	const STEPS = 501 // 26501365
	const EXPECT = 237868

	chunks := (STEPS - size/2) / size
	stepsLeft := (STEPS - start.Y) % size
	ctx.Printf("Chunks: %d; steps left at corners: %d", chunks, stepsLeft)

	answer := 0

	// Top to right
	n, _, img := dec21BFSflood(lines, []cube.Point{{X: start.X, Y: size}}, stepsLeft)
	ctx.Printf("top corner: (%d)\n%s", n, img)
	answer += n
	n, _, img = dec21BFSflood(lines, []cube.Point{{X: start.X, Y: size}}, stepsLeft+size)
	ctx.Printf("below that: (%d)\n%s", n, img)
	answer += n
	n, _, img = dec21BFSflood(lines, []cube.Point{{X: 0, Y: size - 1}, {X: 1, Y: size}}, stepsLeft-1+size/2)
	ctx.Printf("just to the right: (%d)\n%s", n, img)
	answer += n * chunks
	n, _, img = dec21BFSflood(lines, []cube.Point{{X: 0, Y: size - 1}, {X: 1, Y: size}}, stepsLeft-1+size+size/2)
	ctx.Printf("just below that: (%d)\n%s", n, img)
	answer += n * (chunks - 1)

	// Right to bottom
	n, _, img = dec21BFSflood(lines, []cube.Point{{X: 0, Y: start.Y}}, stepsLeft)
	ctx.Printf("right corner: (%d)\n%s", n, img)
	answer += n
	n, _, img = dec21BFSflood(lines, []cube.Point{{X: 0, Y: start.Y}}, stepsLeft+size)
	ctx.Printf("left of that: (%d)\n%s", n, img)
	answer += n
	n, _, img = dec21BFSflood(lines, []cube.Point{{X: 0, Y: 1}, {X: 1, Y: 0}}, stepsLeft-1+size/2)
	ctx.Printf("just below: (%d)\n%s", n, img)
	answer += n * chunks
	n, _, img = dec21BFSflood(lines, []cube.Point{{X: 0, Y: 1}, {X: 1, Y: 0}}, stepsLeft-1+size+size/2)
	ctx.Printf("just left of that: (%d)\n%s", n, img)
	answer += n * (chunks - 1)

	// Bottom to left
	n, _, img = dec21BFSflood(lines, []cube.Point{{X: start.X, Y: 0}}, stepsLeft)
	ctx.Printf("bottom corner: (%d)\n%s", n, img)
	answer += n
	n, _, img = dec21BFSflood(lines, []cube.Point{{X: start.X, Y: 0}}, stepsLeft+size)
	ctx.Printf("above that: (%d)\n%s", n, img)
	answer += n
	n, _, img = dec21BFSflood(lines, []cube.Point{{X: size, Y: 1}, {X: size - 1, Y: 0}}, stepsLeft-1+size/2)
	ctx.Printf("just left of that: (%d)\n%s", n, img)
	answer += n * chunks
	n, _, img = dec21BFSflood(lines, []cube.Point{{X: size, Y: 1}, {X: size - 1, Y: 0}}, stepsLeft-1+size+size/2)
	ctx.Printf("just above: (%d)\n%s", n, img)
	answer += n * (chunks - 1)

	// Left to top
	n, _, img = dec21BFSflood(lines, []cube.Point{{X: size, Y: start.Y}}, stepsLeft)
	ctx.Printf("left corner: (%d)\n%s", n, img)
	answer += n
	n, _, img = dec21BFSflood(lines, []cube.Point{{X: size, Y: start.Y}}, stepsLeft+size)
	ctx.Printf("next to that: (%d)\n%s", n, img)
	answer += n
	n, _, img = dec21BFSflood(lines, []cube.Point{{X: size, Y: size - 1}, {X: size - 1, Y: size}}, stepsLeft-1+size/2)
	ctx.Printf("just above: (%d)\n%s", n, img)
	answer += n * chunks
	n, _, img = dec21BFSflood(lines, []cube.Point{{X: size, Y: size - 1}, {X: size - 1, Y: size}}, stepsLeft-1+size+size/2)
	ctx.Printf("just right of that: (%d)\n%s", n, img)
	answer += n * (chunks - 1)

	// Full chunks
	n, _, img = dec21BFSflood(lines, []cube.Point{{X: start.X, Y: 0}, {X: size, Y: start.Y}, {X: start.X, Y: size}, {X: 0, Y: start.Y}}, size)
	ctx.Printf("full chunk: (%d)\n%s", n, img)
	answer += n * (2*(chunks-1)*(chunks-2) + 4*(chunks-1) + 1)

	if answer == EXPECT {
		return answer, nil
	} else if answer <= 607072529781072 {
		return answer, fmt.Errorf("your answer is too low")
	}

	return answer, errNotImplemented
}

func falseStartDec21b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec21.txt")
	if err != nil {
		return nil, err
	}

	plotsPerChunk := 0
	var start cube.Point
	for y, line := range lines {
		for x, c := range line {
			if c == 'S' {
				start.X, start.Y = x, y
			} else if c == '.' {
				plotsPerChunk++
			}
		}
	}

	const STEPS = 260 + 85 // 26501365
	const STEPSM = STEPS % 2
	size := len(lines) - 1
	chunks := (STEPS + size - 1) / size
	ctx.Printf("Chunk distance: %d", chunks)

	answer := 0
	for k := 0; k < 2; k++ {
		ctx.Printf("Jagged edge %d...", k+1)
		for i := -chunks; i <= chunks; i++ {
			j := -chunks + iabs(i)
			for jk := 0; (j != 0 && jk < 2) || jk == 0; jk++ {
				// ctx.Printf("Handle chunk %d,%d", i, j)
				thischunk := 0
				for x := 0; x < size; x++ {
					for y := 0; y < size; y++ {
						if lines[y][x] == '#' {
							continue
						}
						pt := cube.Point{size*i + x, size*j + y}
						mh := pt.Sub(start).Manhattan()
						if mh <= STEPS && mh%2 == STEPSM {
							// ctx.Printf("   can reach %v", pt)
							thischunk++
						}
					}
				}
				ctx.Printf("Chunk %d,%d: add %d", i, j, thischunk)
				answer += thischunk
				j = -j
			}
			// for j := ; j <= chunks-iabs(i); j++ {
			// }
		}
		chunks--
	}

	ctx.Printf("Chunk distance left: %d", chunks)
	for i := -chunks; i <= chunks; i++ {
		for j := -chunks + iabs(i); j <= chunks-iabs(i); j++ {
			answer += plotsPerChunk
		}
	}

	return answer, nil
}
