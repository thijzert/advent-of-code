package aoc21

import (
	"fmt"
	"log"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec23a(ctx ch.AOContext) error {
	// Worked this out on a napkin
	ctx.FinalAnswer.Print(15160)
	return nil
}

func Dec23b(ctx ch.AOContext) error {
	cave := &amphipodCave{
		hallway: make([]rune, 11),
		rooms:   make([]amphipodRoom, 11),
	}
	cave.rooms[2] = amphipodRoom{Wants: 'A', Contents: []rune{'D', 'D', 'D', 'B'}}
	cave.rooms[4] = amphipodRoom{Wants: 'B', Contents: []rune{'C', 'C', 'B', 'C'}}
	cave.rooms[6] = amphipodRoom{Wants: 'C', Contents: []rune{'A', 'B', 'A', 'D'}}
	cave.rooms[8] = amphipodRoom{Wants: 'D', Contents: []rune{'B', 'A', 'C', 'A'}}

	//cave.rooms[2] = amphipodRoom{Wants: 'A', Contents: []rune{'B', 'D', 'D', 'A'}}
	//cave.rooms[4] = amphipodRoom{Wants: 'B', Contents: []rune{'C', 'C', 'B', 'D'}}
	//cave.rooms[6] = amphipodRoom{Wants: 'C', Contents: []rune{'B', 'B', 'A', 'C'}}
	//cave.rooms[8] = amphipodRoom{Wants: 'D', Contents: []rune{'D', 'A', 'C', 'A'}}

	ctx.Printf("Cave:\n%s", cave)

	//ctx.Print(cave.push(8, 10))
	//ctx.Print(cave.push(8, 0))
	//ctx.Print(cave.push(6, 9))
	//ctx.Print(cave.push(6, 7))
	//ctx.Print(cave.push(6, 1))
	//ctx.Print(cave.push(4, 5))
	//ctx.Print(cave.pop(5, 6))
	//ctx.Print(cave.canPush(4, 5))
	//ctx.Print(cave.push(4, 5))
	//ctx.Print(cave.canPop(5, 6))
	//ctx.Print(cave.pop(5, 6))
	//ctx.Print(cave.canPop(7, 4))
	//ctx.Print(cave.canPop(7, 6))
	//ctx.Print(cave.push(4, 5))
	//ctx.Print(cave.push(4, 3))
	//ctx.Print(cave.canPop(7, 4))
	//ctx.Print(cave.canPop(5, 4))
	//ctx.Print(cave.pop(5, 4))
	//ctx.Print(cave.pop(7, 4))
	//ctx.Print(cave.pop(9, 4))
	//ctx.Print(cave.push(8, 7))
	//ctx.Print(cave.pop(7, 6))
	//ctx.Print(cave.canPop(3, 8))
	//ctx.Print(cave.push(8, 9))
	//ctx.Print(cave.canPop(3, 8))
	//ctx.Print(cave.pop(3, 8))
	//ctx.Print(cave.push(2, 3))
	//ctx.Print(cave.pop(3, 4))
	//ctx.Print(cave.push(2, 5))
	//ctx.Print(cave.push(2, 3))
	//ctx.Print(cave.pop(5, 8))
	//ctx.Print(cave.pop(3, 8))
	//ctx.Print(cave.canPop(1, 2))

	//ctx.Print(cave.pop(1, 2))
	//ctx.Print(cave.pop(0, 2))
	//ctx.Print(cave.pop(9, 2))
	//ctx.Print(cave.pop(10, 8))

	ctx.Printf("Cave:\n%s", cave)
	//return errNotImplemented

	f := func(x rune) int {
		if x == 'B' {
			return 10
		} else if x == 'C' {
			return 100
		} else if x == 'D' {
			return 1000
		} else {
			return 1
		}
	}
	dist, ok := cave.shortestSort(f)
	if !ok {
		return fmt.Errorf("failed to find a solution")
	}

	ctx.FinalAnswer.Print(dist)
	return nil
}

type amphipodRoom struct {
	Wants      rune
	CanReceive bool
	Contents   []rune
}

type amphipodSpecies func(rune) int

type amphipodCave struct {
	hallway []rune
	rooms   []amphipodRoom
}

func (c *amphipodCave) String() string {
	rv := "#"
	for range c.hallway {
		rv += "#"
	}
	rv += "#\n#"
	for _, x := range c.hallway {
		if x == 0 {
			rv += "."
		} else {
			rv += string(x)
		}
	}
	rv += "#\n#"
	roomSize := 0
	for i := range c.hallway {
		if len(c.rooms) > i && len(c.rooms[i].Contents) > 0 {
			roomSize = max(roomSize, len(c.rooms[i].Contents))
			x := c.rooms[i].Contents[0]
			if x == 0 {
				rv += "."
			} else {
				rv += string(x)
			}
		} else {
			rv += "#"
		}
	}
	rv += "#"

	for j := 1; j < roomSize; j++ {
		rv += "\n "
		for i := range c.hallway {
			if len(c.rooms) > i && len(c.rooms[i].Contents) > j {
				x := c.rooms[i].Contents[j]
				if x == 0 {
					rv += "."
				} else {
					rv += string(x)
				}
			} else {
				rv += "#"
			}
		}
	}
	return rv
}

// can we push from room x to hallway position y?
func (cave *amphipodCave) canPush(room, hallPos int) bool {
	if len(cave.rooms[hallPos].Contents) != 0 {
		return false
	}

	// Check if there's anything in the room to move
	var c rune
	for _, x := range cave.rooms[room].Contents {
		if x != 0 && x != '.' {
			c = x
			break
		}
	}
	if c == 0 || c == '.' {
		return false
	}

	// Check if the hallway between start and destination is clear
	x1, x2 := min(room, hallPos), max(room, hallPos)
	for x := x1; x <= x2; x++ {
		if cave.hallway[x] != 0 && cave.hallway[x] != '.' {
			return false
		}
	}

	return true
}

// push regardless, and return the distance moved
func (cave *amphipodCave) push(room, hallPos int) int {
	var c rune
	d := -1
	for i, x := range cave.rooms[room].Contents {
		if x != 0 && x != '.' {
			c = x
			d = i
			break
		}
	}
	if c == 0 || c == '.' {
		log.Printf("Pushing room %d to hallway %d\n%s", room, hallPos, cave)
		panic("nothing to push to the hallway")
	}

	cave.rooms[room].Contents[d] = 0
	cave.hallway[hallPos] = c

	cave.rooms[room].CanReceive = true
	w := cave.rooms[room].Wants
	for _, x := range cave.rooms[room].Contents {
		if x != 0 && x != '.' && x != w {
			cave.rooms[room].CanReceive = false
		}
	}

	return d + abs(room-hallPos) + 1
}

// can we pop from hallway position y into room x?
func (cave *amphipodCave) canPop(hallPos, room int) bool {
	c := cave.hallway[hallPos]
	if c == 0 || c == '.' {
		return false
	}
	if cave.rooms[room].Wants != c || !cave.rooms[room].CanReceive {
		return false
	}

	emptySpot := false
	for _, x := range cave.rooms[room].Contents {
		emptySpot = emptySpot || x == 0 || x == '.'
	}
	if !emptySpot {
		return false
	}

	// Check the hallway
	x1, x2 := min(room, hallPos), max(room, hallPos)
	for x := x1; x <= x2; x++ {
		if x == hallPos {
			continue
		}
		if cave.hallway[x] != 0 && cave.hallway[x] != '.' {
			return false
		}
	}

	return true
}

// pop regardless, and return the distance moved
func (cave *amphipodCave) pop(hallPos, room int) int {
	d := -1
	for i := range cave.rooms[room].Contents {
		x := cave.rooms[room].Contents[i]
		if x == 0 || x == '.' {
			d = i
		}
	}

	c := cave.hallway[hallPos]
	cave.rooms[room].Contents[d] = c
	cave.hallway[hallPos] = 0

	if c != cave.rooms[room].Wants {
		cave.rooms[room].CanReceive = false
	}

	return d + abs(room-hallPos) + 1
}

func (cave *amphipodCave) shortestSort(f amphipodSpecies) (int, bool) {
	// Check if we're done already
	done := true
	for _, c := range cave.hallway {
		done = done && (c == 0 || c == '.')
	}
	if done {
		for _, r := range cave.rooms {
			done = done && (len(r.Contents) == 0 || r.CanReceive)
		}
	}
	if done {
		return 0, true
	}

	rv := 0
	found := false

	for i, c := range cave.hallway {
		if c == 0 || c == '.' {
			continue
		}

		for j := range cave.rooms {
			if cave.canPop(i, j) {
				d := f(c) * cave.pop(i, j)
				shr, ok := cave.shortestSort(f)
				if ok {
					if !found || shr+d < rv {
						found = true
						rv = shr + d
					}
				}

				cave.push(j, i)
			}
		}
	}

	for j, r := range cave.rooms {
		if len(r.Contents) == 0 || r.CanReceive {
			continue
		}

		c := rune(0)
		for k := range r.Contents {
			if r.Contents[k] != 0 && r.Contents[k] != '.' {
				c = r.Contents[k]
				break
			}
		}

		for i, x := range cave.hallway {
			if x != 0 && c != '.' {
				continue
			}

			if cave.canPush(j, i) {
				d := f(c) * cave.push(j, i)
				shr, ok := cave.shortestSort(f)
				if ok {
					if !found || shr+d < rv {
						found = true
						rv = shr + d
					}
				}

				cave.pop(i, j)
			}
		}
	}

	return rv, found
}
