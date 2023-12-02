package aoc23

import (
	"fmt"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

type RGB struct {
	R, G, B int
}

func (a RGB) LE(b RGB) bool {
	return a.R <= b.R && a.G <= b.G && a.B <= b.B
}

func (a RGB) Power() int {
	return a.R * a.G * a.B
}

type cubeGame struct {
	ID   int
	Sets []RGB
}

func dec2Cubegames(ctx ch.AOContext) ([]cubeGame, error) {
	lines, err := ctx.DataLines("inputs/2023/dec02.txt")
	if err != nil {
		return nil, err
	}
	//lines = []string{
	//	"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	//	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	//	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	//	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	//	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	//}

	var rv []cubeGame
	for _, line := range lines {
		var g cubeGame
		fmt.Sscanf(line, "Game %d: ", &g.ID)
		setss := strings.Split(line, ": ")
		sets := strings.Split(setss[1], "; ")
		for _, s := range sets {
			var rgb RGB
			counts := strings.Split(s, ", ")
			var n int
			var c string
			for _, count := range counts {
				fmt.Sscanf(count, "%d %s", &n, &c)
				if c == "red" {
					rgb.R = n
				} else if c == "green" {
					rgb.G = n
				} else if c == "blue" {
					rgb.B = n
				}
			}
			g.Sets = append(g.Sets, rgb)
		}
		rv = append(rv, g)
	}

	return rv, nil
}

func Dec02a(ctx ch.AOContext) (interface{}, error) {
	games, err := dec2Cubegames(ctx)
	if err != nil {
		return nil, err
	}

	maxRGB := RGB{12, 13, 14}
	answer := 0
	for _, game := range games {
		possible := true
		for _, set := range game.Sets {
			possible = possible && set.LE(maxRGB)
		}
		if possible {
			answer += game.ID
		}
	}

	return answer, nil
}

func Dec02b(ctx ch.AOContext) (interface{}, error) {
	games, err := dec2Cubegames(ctx)
	if err != nil {
		return nil, err
	}

	answer := 0
	for _, game := range games {
		maxRGB := RGB{0, 0, 0}
		for _, set := range game.Sets {
			maxRGB.R = max(maxRGB.R, set.R)
			maxRGB.G = max(maxRGB.G, set.G)
			maxRGB.B = max(maxRGB.B, set.B)
		}
		answer += maxRGB.Power()
	}

	return answer, nil
}
