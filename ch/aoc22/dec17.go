package aoc22

import (
	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
	"github.com/thijzert/advent-of-code/lib/image"
)

func Dec17a(ctx ch.AOContext) (interface{}, error) {
	return dec17(ctx, 2022)
}

func Dec17b(ctx ch.AOContext) (interface{}, error) {
	return dec17(ctx, 1000000000000)
}

func dec17(ctx ch.AOContext, MAXROCKS int64) (interface{}, error) {
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

	type lookingBack struct {
		RockIndex    int
		JetIndex     int
		LandingSpots [40]cube.Point
	}
	type rockState struct {
		RocksFallen int64
		TowerHeight int64
	}

	seenBefore := make(map[lookingBack]rockState)
	lb := lookingBack{}
	skip, lengthExtension := int64(0), int64(0)

	towerHeight := int64(0)
	j := 0
	for r := 0; skip+int64(r) < MAXROCKS; r++ {
		rock := rocks[r%len(rocks)]
		offset := cube.Point{3, int(towerHeight) + rock.Height + 3}

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

		lb.RockIndex = r % len(rocks)
		lb.JetIndex = j
		copy(lb.LandingSpots[1:], lb.LandingSpots[:len(lb.LandingSpots)-1])
		lb.LandingSpots[0] = cube.Point{offset.X, offset.Y - int(towerHeight)}

		tunnel.Sprite(rock, rock, offset.X, HEIGHT-offset.Y-1)
		if int64(offset.Y) > towerHeight {
			towerHeight = int64(offset.Y)
		}
		if r&0xff == 0 {
			ctx.Printf("after %d rocks: tower height: %d", r+1, towerHeight)
		}

		if skip != 0 {
		} else if st, ok := seenBefore[lb]; ok {
			num, den := towerHeight-st.TowerHeight, int64(r)-st.RocksFallen
			n := (MAXROCKS - int64(r)) / den
			if n > 0 {
				skip = den * n
				lengthExtension = num * n
				ctx.Printf("I've seen this one before, after rock %d", st.RocksFallen)
				ctx.Printf("Seems that for every %d rocks, the tower gets %d spaces bigger.", den, num)
				ctx.Printf("I should skip %d rocks and add %d to the tower height at the end", skip, lengthExtension)
			}
		} else if r > len(lb.LandingSpots) {
			seenBefore[lb] = rockState{
				RocksFallen: int64(r),
				TowerHeight: towerHeight,
			}
		}

		if towerHeight+int64(tunnel.OffsetY) > int64(HEIGHT-20) {
			shiftLines := 3
			tunnel.OffsetY -= shiftLines
			copy(tunnel.Contents[shiftLines*WIDTH:], tunnel.Contents[:len(tunnel.Contents)-shiftLines*WIDTH])
			//ctx.Printf("\n%s", tunnel)
		}
	}
	ctx.Printf("\n%s", tunnel)

	return towerHeight + lengthExtension, nil
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
