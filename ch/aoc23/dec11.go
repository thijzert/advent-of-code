package aoc23

import (
	"github.com/thijzert/advent-of-code/ch"
)

func dec11galaxies(ctx ch.AOContext, dilationFactor int) ([][2]int, error) {
	lines, err := ctx.DataLines("inputs/2023/dec11.txt")
	if err != nil {
		return nil, err
	}
	dilationX, dilationY := make([]int, len(lines[0])), make([]int, len(lines))

	for y, line := range lines {
		dilationY[y] = dilationFactor
		for _, ch := range line {
			if ch == '#' {
				dilationY[y] = 1
			}
		}
	}
	d := 0
	for i, c := range dilationY {
		dilationY[i] = d
		d += c
	}
	for x := range lines[0] {
		dilationX[x] = dilationFactor
		for _, line := range lines {
			if line[x] == '#' {
				dilationX[x] = 1
			}
		}
	}
	d = 0
	for i, c := range dilationX {
		dilationX[i] = d
		d += c
	}

	galaxies := make([][2]int, 0)
	for y, line := range lines {
		for x, ch := range line {
			if ch == '#' {
				galaxies = append(galaxies, [2]int{dilationX[x], dilationY[y]})
			}
		}
	}

	return galaxies, nil
}

func dec11dist(galaxies [][2]int, i, j int) int {
	dx := galaxies[i][0] - galaxies[j][0]
	if dx < 0 {
		dx = -dx
	}
	dy := galaxies[i][1] - galaxies[j][1]
	if dy < 0 {
		dy = -dy
	}
	return dx + dy
}

func Dec11a(ctx ch.AOContext) (interface{}, error) {
	galaxies, err := dec11galaxies(ctx, 2)
	if err != nil {
		return nil, err
	}

	ctx.Printf("The distance between galaxy 1 and 7: %d", dec11dist(galaxies, 0, 6))
	ctx.Printf("The distance between galaxy 3 and 6: %d", dec11dist(galaxies, 2, 5))
	ctx.Printf("The distance between galaxy 8 and 9: %d", dec11dist(galaxies, 7, 8))

	answer := 0
	for i := range galaxies {
		for j := range galaxies[i+1:] {
			answer += dec11dist(galaxies, i, i+j+1)
		}
	}

	return answer, nil
}

func Dec11b(ctx ch.AOContext) (interface{}, error) {
	galaxies, err := dec11galaxies(ctx, 1000000)
	if err != nil {
		return nil, err
	}

	ctx.Printf("The distance between galaxy 1 and 7: %d", dec11dist(galaxies, 0, 6))
	ctx.Printf("The distance between galaxy 3 and 6: %d", dec11dist(galaxies, 2, 5))
	ctx.Printf("The distance between galaxy 8 and 9: %d", dec11dist(galaxies, 7, 8))

	answer := 0
	for i := range galaxies {
		for j := range galaxies[i+1:] {
			answer += dec11dist(galaxies, i, i+j+1)
		}
	}

	return answer, nil
}
