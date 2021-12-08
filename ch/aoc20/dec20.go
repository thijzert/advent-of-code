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

	twoNeighbours := []*satelliteTile{}

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
	img, err := formSatelliteImage(ctx, "inputs/2020/dec20.txt")
	if img.Tiles != nil && len(img.Tiles) == img.Size*img.Size {
		ctx.Printf("Tile ID's:")
		for i := 0; i < len(img.Tiles); i += img.Size {
			ids := make([]int, img.Size)
			for j, n := range img.Tiles[i : i+img.Size] {
				if n.Tile != nil {
					ids[j] = n.Tile.ID
				}
			}
			ctx.Printf("    %4d", ids)
		}
		ctx.Printf("Image: (%d×%d)\n%s", img.Size, img.Size, img)
	}
	if err != nil {
		return err
	}

	ctx.Printf("Image: (%d×%d)\n%s", img.Size, img.Size, img)

	return errors.New("not implemented")
}

type satelliteImage struct {
	Size  int
	Tiles []satelliteImageTile
}

func formSatelliteImage(ctx ch.AOContext, assetName string) (satelliteImage, error) {
	tiles, err := readSatelliteImagery(ctx, assetName)
	if err != nil {
		return satelliteImage{}, err
	}

	size := 0
	for size*size < len(tiles) {
		size++
	}
	if size*size > len(tiles) {
		return satelliteImage{}, fmt.Errorf("was expecting a %d×%d image, but have %d tiles", size, size, len(tiles))
	}

	ctx.Print(len(tiles))

	type tileNeighbours struct {
		Tile        *satelliteTile
		Neighbours  [4][]*tileNeighbours
		NNeighbours int
		Used        bool
	}
	allNeighbours := make([]*tileNeighbours, len(tiles))
	for i, tile := range tiles {
		allNeighbours[i] = &tileNeighbours{
			Tile: tile,
		}
	}

	for i, tile := range allNeighbours {
		for k := 0; k < 4; k++ {
			edge := tile.Tile.Edges[k]
			for j, otherTile := range allNeighbours {
				if j == i {
					continue
				}
				fits := false
				for l, otherEdge := range otherTile.Tile.Rdges {
					if edge == otherEdge || edge == otherTile.Tile.Edges[l] {
						fits = true
					}
				}
				if fits {
					tile.Neighbours[k] = append(tile.Neighbours[k], otherTile)
					tile.NNeighbours++
				}
			}
		}
	}

	type protoTile struct {
		TN       *tileNeighbours
		Rotation int
		Flipped  bool
	}
	pNeighbours := func(p protoTile, side int) (int, []*tileNeighbours) {
		if p.TN == nil {
			return -1, nil
		}

		n := (side + p.Rotation) % 4
		if p.Flipped {
			return p.TN.Tile.Edges[4+n], p.TN.Neighbours[(8-n)%4]
		}
		return p.TN.Tile.Edges[n], p.TN.Neighbours[n]
	}
	picture := make([]protoTile, size*size)

	rv := satelliteImage{
		Size:  size,
		Tiles: make([]satelliteImageTile, size*size),
	}

	defer func() {
		for i, ptile := range picture {
			if ptile.TN != nil {
				rv.Tiles[i].Tile = ptile.TN.Tile
				rv.Tiles[i].Rotation = ptile.Rotation
				rv.Tiles[i].Flipped = ptile.Flipped
			}
		}
	}()

	var iTopLeft int
	cornerPieces := 0
	for i, tile := range allNeighbours {
		if tile.NNeighbours == 2 {
			// Put one of the corner pieces in the top left
			if picture[0].TN == nil {
				picture[0].TN = tile
				tile.Used = true
				iTopLeft = i
				ctx.Printf("Can put tile %d on %d,%d", tile.Tile.ID, 0, 0)
			}

			cornerPieces++
			// ctx.Printf("Tile %d is probably a corner:\n%s", tile.Tile.ID, tile.Tile)
		}
	}
	if cornerPieces == 0 {
		return satelliteImage{}, fmt.Errorf("could not find exactly 4 corner pieces")
	}

	// Try to find the correct orientation for this first tile.
	// It should have one neighbour to the right, and one below
	for i := 0; i < 4; i++ {
		if len(allNeighbours[iTopLeft].Neighbours[(i+1)%4]) > 0 && len(allNeighbours[iTopLeft].Neighbours[(i+2)%4]) > 0 {
			picture[0].Rotation = i
			break
		}
	}

	// HACK: fudge the data so that it's identical to the example
	picture[0].Flipped = true
	picture[0].Rotation = 2

	for k := 0; k < 4; k++ {
		c, poss := pNeighbours(picture[0], k)
		id := 0
		if len(poss) > 0 {
			id = poss[0].Tile.ID
		}
		ctx.Printf("Side %d: %010b; %d neighbour, tile %d", k, c, len(poss), id)
	}

	tilesRemaining := len(tiles) - 1
	for tilesRemaining > 0 {
		changed := false
		for y := 0; y < rv.Size; y++ {
			for x := 0; x < rv.Size; x++ {
				i := rv.Size*y + x
				if picture[i].TN != nil {
					continue
				}
				adj := []struct {
					X, Y int
				}{{x, y + 1}, {x - 1, y}, {x, y - 1}, {x + 1, y}}
				poss := make(map[*tileNeighbours]bool)
				constraints := []int{-1, -1, -1, -1}
				for k, coord := range adj {
					if coord.X < 0 || coord.X >= size || coord.Y < 0 || coord.Y >= size {
						continue
					}
					n := size*coord.Y + coord.X
					if n < 0 || n >= len(rv.Tiles) || picture[n].TN == nil {
						continue
					}
					c, neigh := pNeighbours(picture[n], k)
					constraints[k] = c
					for _, tn := range neigh {
						if !tn.Used {
							poss[tn] = true
						}
					}
				}

				ctx.Printf("[%d,%d] Constraints: %010b; %d possibilities", x, y, constraints, len(poss))
				if len(poss) == 1 {
					for tn := range poss {
						ctx.Printf("Can put tile %d on %d,%d", tn.Tile.ID, x, y)
						picture[i].TN = tn
						tn.Used = true
						tilesRemaining--
						changed = true
						correct := false

						// Figure out the orientation
						for r := 0; r < 8; r++ {
							picture[i].Rotation = r % 4
							picture[i].Flipped = r >= 4
							correct = true

							for k, c := range constraints {
								if c == -1 {
									continue
								}
								c0, _ := pNeighbours(picture[i], (k+2)%4)
								ctx.Printf("Rotated %d times, side %d is: %010b - looking for %010b", r, k, c0, c)
								correct = correct && (c == c0)
							}
							if correct {
								break
							}
						}
						if !correct {
							return rv, fmt.Errorf("there's no correct orientation for tile %d at %d,%d", tn.Tile.ID, x, y)
						}
					}
				}
			}
		}
		if !changed {
			return rv, fmt.Errorf("got stuck with %d of %d tiles remaining", tilesRemaining, len(tiles))
		}
	}

	ctx.Printf("Tiles remaining: %d of %d", tilesRemaining, len(tiles))

	return rv, nil
}

func (i satelliteImage) String() string {
	rv := ""

	for y := 0; y < 8*i.Size; y += 2 {
		if y > 0 {
			rv += "\n"
		}
		for x := 0; x < 8*i.Size; x++ {
			j := i.Size*(y/8) + (x / 8)
			t, b := 0, 0
			if i.Tiles[j].Tile != nil {
				t, b = i.Tiles[j].At(x%8, y%8), i.Tiles[j].At(x%8, 1+y%8)
			}

			if t != 0 && b != 0 {
				rv += "\u2588"
			} else if t != 0 {
				rv += "\u2580"
			} else if b != 0 {
				rv += "\u2584"
			} else {
				rv += " "
			}
		}
	}

	return rv
}

func (i satelliteImage) At(x, y int) int {
	j := i.Size*(y/8) + (x / 8)
	if i.Tiles[j].Tile == nil {
		return 0
	}

	return i.Tiles[j].At(x%8, y%8)
}

type satelliteImageTile struct {
	Tile     *satelliteTile
	Flipped  bool
	Rotation int
}

func (t satelliteImageTile) Edge(side int) int {
	edges := t.Tile.Edges
	if t.Flipped {
		side = (8 - side) % 4
		edges = t.Tile.Rdges
	}

	return edges[side]
}

func (t satelliteImageTile) At(x, y int) int {
	if t.Flipped {
		x = 7 - x
	}

	if t.Rotation == 1 {
		x, y = 7-y, x
	} else if t.Rotation == 2 {
		x, y = 7-x, 7-y
	} else if t.Rotation == 3 {
		x, y = y, 7-x
	}

	return t.Tile.Contents[10*y+x+11]
}

type satelliteTile struct {
	ID       int
	Contents []int
	Edges    []int
	Rdges    []int
}

func readSatelliteImagery(ctx ch.AOContext, assetName string) ([]*satelliteTile, error) {
	lines, err := ctx.DataLines(assetName)
	if err != nil {
		return nil, err
	}

	var rv []*satelliteTile
	for len(lines) > 10 {
		tile := &satelliteTile{
			Contents: make([]int, 100),
			Edges:    make([]int, 8),
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
			tile.Edges[4] = tile.Edges[4] | tile.Contents[i]<<i

			// right
			tile.Edges[1] = tile.Edges[1]<<1 | tile.Contents[9+i*10]
			tile.Rdges[3] = tile.Rdges[3]<<1 | tile.Contents[99-(10*i)]
			tile.Edges[7] = tile.Edges[7] | tile.Contents[9+i*10]<<i

			// bottom
			tile.Edges[2] = tile.Edges[2]<<1 | tile.Contents[99-i]
			tile.Rdges[2] = tile.Rdges[2]<<1 | tile.Contents[90+i]
			tile.Edges[6] = tile.Edges[6] | tile.Contents[99-i]<<i

			// left
			tile.Edges[3] = tile.Edges[3]<<1 | tile.Contents[90-(10*i)]
			tile.Rdges[1] = tile.Rdges[1]<<1 | tile.Contents[i*10]
			tile.Edges[5] = tile.Edges[5] | tile.Contents[90-(10*i)]<<i
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

// func (t satelliteTile) FindOrientation(direction, edge int) (rotation int, flipped bool, nWays int) {
// 	for i := 0; i < 4; i++ {
//
// 	}
// }
