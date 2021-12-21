package aoc21

import (
	"github.com/thijzert/advent-of-code/ch"
)

var Dec21b ch.AdventFunc = nil

func Dec21a(ctx ch.AOContext) error {
	// playerPos := []int{4, 8}
	playerPos := []int{2, 5}
	playerScore := []int{0, 0}

	var die DeterministicDice = &DetD100{}

	playing := true
	for playing {
		for i := range playerPos {
			move := die.Next() + die.Next() + die.Next()
			playerPos[i] = 1 + ((playerPos[i] + move - 1) % 10)
			playerScore[i] += playerPos[i]
			//ctx.Printf("Player %d rolls %d and moves to %d, bringing their score to %d", i+1, move, playerPos[i], playerScore[i])
			if playerScore[i] >= 1000 {
				playing = false
				break
			}
		}
	}

	ctx.FinalAnswer.Print(die.NRolls() * min(playerScore...))
	return nil
}

// func Dec21b(ctx ch.AOContext) error {
// 	return errNotImplemented
// }

type DeterministicDice interface {
	Next() int
	NRolls() int
}

type DetD100 struct {
	current int
	nRolls  int
}

func (d *DetD100) Next() int {
	d.current++
	d.nRolls++
	rv := d.current
	d.current %= 100
	return rv
}

func (d *DetD100) NRolls() int {
	return d.nRolls
}
