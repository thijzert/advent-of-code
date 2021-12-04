package aoc20

import (
	"errors"
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec20a(ctx ch.AOContext) error {
	tiles, err := readSatelliteImagery(ctx, "inputs/2020/dec20.txt")
	if err != nil {
		return err
	}

	ctx.Print(len(tiles))

	twoNeighbours := []satelliteTile{}

	for i, tile := range tiles {
		neighbours := 0
		for j, otherTile := range tiles {
			if j == i {
				continue
			}
			canMeet := false
			for _, edge := range tile.Edges {
				for k, otherEdge := range otherTile.Edges {
					if edge == otherEdge || edge == otherTile.Rdges[k] {
						canMeet = true
					}
				}
			}
			if canMeet {
				neighbours++
			}
		}

		// ctx.Printf("Tile %d has %d neighbours:\n%s", tile.ID, neighbours, tile)
		if neighbours == 2 {
			twoNeighbours = append(twoNeighbours, tile)
		}
	}

	rv := 1
	for _, tile := range twoNeighbours {
		ctx.Printf("Tile %d is probably a corner:\n%s", tile.ID, tile)
		rv *= tile.ID
	}
	if len(twoNeighbours) != 4 {
		return fmt.Errorf("could not find exactly 4 corner pieces")
	}

	ctx.FinalAnswer.Print(rv)
	return nil
}

func Dec20b(ctx ch.AOContext) error {
	return errors.New("not implemented")
}

type satelliteTile struct {
	ID       int
	Contents []int
	Edges    []int
	Rdges    []int
}

func readSatelliteImagery(ctx ch.AOContext, assetName string) ([]satelliteTile, error) {
	lines, err := ctx.DataLines(assetName)
	if err != nil {
		return nil, err
	}

	var rv []satelliteTile
	for len(lines) > 10 {
		tile := satelliteTile{
			Contents: make([]int, 100),
			Edges:    make([]int, 4),
			Rdges:    make([]int, 4),
		}
		fmt.Sscanf(lines[0], "Tile %d:", &tile.ID)
		for i, l := range lines[1:11] {
			for j, c := range l {
				if c == '#' {
					tile.Contents[10*i+j] = 1
				}
			}
		}

		// Calculate edges
		for i := 0; i < 10; i++ {
			// top
			tile.Edges[0] = tile.Edges[0]<<1 | tile.Contents[i]
			tile.Rdges[0] = tile.Rdges[0]<<1 | tile.Contents[9-i]

			// right
			tile.Edges[1] = tile.Edges[1]<<1 | tile.Contents[9+i*10]
			tile.Rdges[3] = tile.Rdges[3]<<1 | tile.Contents[99-(10*i)]

			// bottom
			tile.Edges[2] = tile.Edges[2]<<1 | tile.Contents[99-i]
			tile.Rdges[2] = tile.Rdges[2]<<1 | tile.Contents[90+i]

			// left
			tile.Edges[3] = tile.Edges[3]<<1 | tile.Contents[90-(10*i)]
			tile.Rdges[1] = tile.Rdges[1]<<1 | tile.Contents[i*10]
		}

		rv = append(rv, tile)
		lines = lines[12:]
	}

	return rv, nil
}

func (t satelliteTile) String() string {
	rv := ""
	for i := 0; i < 10; i += 2 {
		if i > 0 {
			rv += "\n"
		}
		for j := 0; j < 10; j++ {
			if t.Contents[10*i+j] == 1 && t.Contents[10+10*i+j] == 1 {
				rv += "\u2588"
			} else if t.Contents[10*i+j] == 1 {
				rv += "\u2580"
			} else if t.Contents[10+10*i+j] == 1 {
				rv += "\u2584"
			} else {
				rv += " "
			}
		}
	}
	return rv
}
