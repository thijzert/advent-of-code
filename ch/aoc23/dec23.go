package aoc23

import (
	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
)

func Dec23a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec23.txt")
	if err != nil {
		return nil, err
	}
	return dec23(ctx, lines)
}

func Dec23b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec23.txt")
	if err != nil {
		return nil, err
	}
	for i, line := range lines {
		buf := []byte(line)
		for j, c := range buf {
			if c != '#' {
				buf[j] = '.'
			}
		}
		lines[i] = string(buf)
	}
	return dec23(ctx, lines)
}

func dec23(ctx ch.AOContext, lines []string) (interface{}, error) {
	start, finish := cube.Pt(1, 0), cube.Pt(len(lines[0])-2, len(lines)-1)

	graph := dec23ReadGraph(lines, start, finish)
	for pt, otherCorners := range graph {
		ctx.Printf("From %v: %v", pt, otherCorners)
	}

	corners := []cube.Point{}
	cornerIdx := make(map[cube.Point]int)
	for pt := range graph {
		cornerIdx[pt] = len(corners)
		corners = append(corners, pt)
	}

	type corn struct {
		Location int
		Mask     int64
	}

	front := make(map[corn]int)
	front[corn{
		Location: cornerIdx[start],
		Mask:     1 << cornerIdx[start],
	}] = 0
	answer, gen := 0, 0
	for len(front) > 0 {
		gen++
		if gen > 1000 {
			return nil, errFailed
		}
		newFront := make(map[corn]int)
		for crn, dist := range front {
			pt := corners[crn.Location]
			for pt1, dist1 := range graph[pt] {
				d := dist + dist1
				if pt1 == finish && d > answer {
					answer = d
				}
				j := cornerIdx[pt1]
				if crn.Mask&(1<<j) != 0 {
					continue
				}
				crn1 := corn{
					Location: j,
					Mask:     crn.Mask | (1 << j),
				}
				if d > newFront[crn1] {
					newFront[crn1] = d
				}
			}
		}
		ctx.Printf("front: %d", len(newFront))
		front = newFront
	}

	return answer, nil
}

func dec23ReadGraph(lines []string, start, finish cube.Point) map[cube.Point]map[cube.Point]int {
	corners := make(map[cube.Point]bool)
	corners[start] = true
	corners[finish] = true
	for y, line := range lines {
		if y == 0 || y == len(lines)-1 {
			continue
		}
		for x, c := range line {
			if c != '.' {
				continue
			}
			pt := cube.Pt(x, y)
			con := 0
			for _, dir := range cube.Cardinal2D {
				pt1 := pt.Add(dir)
				if lines[pt1.Y][pt1.X] != '#' {
					con++
				}
			}
			if con > 2 {
				corners[pt] = true
			}
		}
	}

	rv := make(map[cube.Point]map[cube.Point]int)
	for pt := range corners {
		rv[pt] = dec23bfs(lines, pt, corners)
	}
	return rv
}

func dec23bfs(lines []string, start cube.Point, corners map[cube.Point]bool) map[cube.Point]int {
	rv := make(map[cube.Point]int)
	visited := make(map[cube.Point]bool)
	visited[start] = true

	const slopes string = ">v<^"

	dist := 0
	heads := []cube.Point{start}
	for len(heads) > 0 {
		dist++
		newHeads := make([]cube.Point, 0, len(heads))

		for _, pt := range heads {
			for i, dir := range cube.Cardinal2D {
				pt1 := pt.Add(dir)
				if visited[pt1] {
				} else if corners[pt1] {
					rv[pt1] = dist
				} else if pt1.Y < 0 || pt1.Y >= len(lines) {
				} else if pt1.X < 0 || pt1.X >= len(lines[0]) {
				} else if lines[pt1.Y][pt1.X] == '.' || lines[pt1.Y][pt1.X] == slopes[i] {
					newHeads = append(newHeads, pt1)
					visited[pt1] = true
				}
			}
		}

		heads = newHeads
	}

	return rv
}
