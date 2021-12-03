package aoc20

import (
	"errors"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec22a(ctx ch.AOContext) error {
	decks, err := ctx.DataAsInts("inputs/2020/dec22.txt")
	if err != nil {
		return err
	}

	playerA := make([]int, len(decks)/2, len(decks))
	playerB := make([]int, len(decks)/2, len(decks))

	copy(playerA, decks[:len(decks)/2])
	copy(playerB, decks[len(decks)/2:])

	i := 0
	for len(playerA) > 0 && len(playerB) > 0 {
		i++
		ctx.Debug.Printf("--- round %d ---", i)
		ctx.Debug.Printf("Player A: %d", playerA)
		ctx.Debug.Printf("Player B: %d", playerB)

		cardA := takeTopCard(&playerA)
		cardB := takeTopCard(&playerB)
		ctx.Debug.Printf("A gets %d, B gets %d", cardA, cardB)

		if cardA > cardB {
			playerA = append(playerA, cardA, cardB)
		} else {
			playerB = append(playerB, cardB, cardA)
		}
	}

	ctx.Debug.Printf("Player A: %d", playerA)
	ctx.Debug.Printf("Player B: %d", playerB)

	winner := playerA
	if len(playerA) == 0 {
		winner = playerB
	}

	score := 0
	for i, card := range winner {
		score += card * (len(winner) - i)
	}

	ctx.FinalAnswer.Print(score)
	return nil
}
func Dec22b(ctx ch.AOContext) error {
	return errors.New("not implemented")
}

func takeTopCard(deck *[]int) int {
	rv := (*deck)[0]

	copy(*deck, (*deck)[1:])
	*deck = (*deck)[:len(*deck)-1]

	return rv
}
