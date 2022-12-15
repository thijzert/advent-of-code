package aoc19

import (
	"fmt"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
)

func Dec03a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2019/dec03.txt")
	if err != nil {
		return nil, err
	}
	//lines = []string{"R8,U5,L5,D3", "U7,R6,D4,L4"}
	//lines = []string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}
	//lines = []string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}

	visited := make(map[cube.Point]int)
	for wire, line := range lines {
		instrs := strings.Split(line, ",")
		head := cube.Point{0, 0}
		for _, instr := range instrs {
			var a string
			var l int
			fmt.Sscanf(instr, "%1s%d", &a, &l)
			dir, _ := cube.ParseDirection2D(a)
			for i := 0; i < l; i++ {
				head = head.Add(dir)
				if visited[head] == wire {
					visited[head] = wire + 1
				}
			}
		}
	}

	closest := 0xffffff
	for pt, wires := range visited {
		if wires != len(lines) {
			continue
		}
		ctx.Printf("Wires intersect at %v", pt)
		mh := iabs(pt.X) + iabs(pt.Y)
		if mh < closest {
			closest = mh
		}
	}
	return closest, nil
}

func iabs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Dec03b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2019/dec03.txt")
	if err != nil {
		return nil, err
	}

	var lastVisited map[cube.Point]int
	for _, line := range lines {
		visited := make(map[cube.Point]int)
		instrs := strings.Split(line, ",")
		head := cube.Point{0, 0}
		length := 0
		for _, instr := range instrs {
			var a string
			var l int
			fmt.Sscanf(instr, "%1s%d", &a, &l)
			dir, _ := cube.ParseDirection2D(a)
			for i := 0; i < l; i++ {
				head = head.Add(dir)
				length++
				if visited[head] > 0 {
				} else if lastVisited == nil || lastVisited[head] > 0 {
					visited[head] = lastVisited[head] + length
				}
			}
		}
		lastVisited = visited
	}

	closest := 0xffffff
	for pt, lengths := range lastVisited {
		ctx.Printf("Wires intersect at %v", pt)
		if lengths < closest {
			closest = lengths
		}
	}
	return closest, nil
}
