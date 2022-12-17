package aoc22

import (
	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
	"github.com/thijzert/advent-of-code/lib/image"
)

func Dec17a(ctx ch.AOContext) (interface{}, error) {
	return dec17(ctx, 2022)
	return dec17(ctx, 1000000000)
}

var Dec17b ch.AdventFunc = nil

// func Dec17b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }

func dec17(ctx ch.AOContext, MAXROCKS int) (interface{}, error) {
	gasjets, err := ctx.DataLines("inputs/2022/dec17.txt")
	if err != nil {
		return nil, err
	}
	//gasjets = []string{">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"}
	gasjet := gasjets[0]

	rocks := fallingRocks()
	for _, b := range rocks {
		ctx.Printf("\n%s", b)
	}

	HEIGHT, WIDTH := 80, 9
	tunnel := image.NewImage(WIDTH, HEIGHT, func(x, y int) int {
		if x == 0 || x == WIDTH-1 || y == HEIGHT-1 {
			return 2
		}
		return 0
	})

	towerHeight := 0
	j := 0
	for r := 0; r < MAXROCKS; r++ {
		rock := rocks[r%len(rocks)]
		offset := cube.Point{3, towerHeight + rock.Height + 3}

		for {
			tryX := offset.X + 1
			jet := gasjet[j]
			j = (j + 1) % len(gasjet)
			if jet == '<' {
				tryX = offset.X - 1
			}

			// Move the rock to the side if possible
			if tunnel.MaskAt(rock, tryX, HEIGHT-offset.Y-1) == 0 {
				offset.X = tryX
			}

			// Try to drop it down
			if tunnel.MaskAt(rock, offset.X, HEIGHT-offset.Y) != 0 {
				break
			}
			offset.Y--
		}

		tunnel.Sprite(rock, rock, offset.X, HEIGHT-offset.Y-1)
		if offset.Y > towerHeight {
			towerHeight = offset.Y
		}
		if r&0xfff == 0 {
			ctx.Printf("block y: %d, offset y: %d, tower height: %d", offset.Y, tunnel.OffsetY, towerHeight)
		}

		if towerHeight+tunnel.OffsetY > HEIGHT-20 {
			shiftLines := 3
			tunnel.OffsetY -= shiftLines
			copy(tunnel.Contents[shiftLines*WIDTH:], tunnel.Contents[:len(tunnel.Contents)-shiftLines*WIDTH])
			//ctx.Printf("\n%s", tunnel)
		}
	}
	ctx.Printf("\n%s", tunnel)

	return towerHeight, nil
}

func fallingRocks() []*image.Image {
	return []*image.Image{
		image.ReadImage([]string{"####"}, image.Octothorpe),
		image.ReadImage([]string{".#.", "###", ".#."}, image.Octothorpe),
		image.ReadImage([]string{"..#", "..#", "###"}, image.Octothorpe),
		image.ReadImage([]string{"#", "#", "#", "#"}, image.Octothorpe),
		image.ReadImage([]string{"##", "##"}, image.Octothorpe),
	}
}
