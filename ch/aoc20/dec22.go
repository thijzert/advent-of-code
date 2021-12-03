package aoc20

import (
	"fmt"
	"log"

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

	score := combatScore(playerA) + combatScore(playerB)

	ctx.FinalAnswer.Print(score)
	return nil
}

func takeTopCard(deck *[]int) int {
	rv := (*deck)[0]

	copy(*deck, (*deck)[1:])
	*deck = (*deck)[:len(*deck)-1]

	return rv
}

func combatScore(deck []int) int {
	score := 0
	for i, card := range deck {
		score += card * (len(deck) - i)
	}
	return score
}

func Dec22b(ctx ch.AOContext) error {
	decks, err := ctx.DataAsInts("inputs/2020/dec22.txt")
	if err != nil {
		return err
	}

	cutoff := len(decks) / 2

	rkk := &recursiveKrabKombat{
		L:           ctx.Debug,
		DeckHistory: make(map[string]bool),
	}

	_, score := rkk.Combat(decks[:cutoff], decks[cutoff:])

	ctx.FinalAnswer.Print(score)
	return nil
}

type recursiveKrabKombat struct {
	L           *log.Logger
	Games       int
	DeckHistory map[string]bool
}

func (k *recursiveKrabKombat) Combat(deckA, deckB []int) (aWins bool, score int) {
	k.Games++
	gameNo := k.Games

	prefix := k.L.Prefix()
	defer func() {
		k.L.SetPrefix(prefix)
	}()

	playerA := make([]int, len(deckA))
	playerB := make([]int, len(deckB))
	copy(playerA, deckA)
	copy(playerB, deckB)

	k.L.Printf("===== game %d =====", gameNo)

	k.L.SetPrefix("    " + prefix)

	i := 0
	for len(playerA) > 0 && len(playerB) > 0 {
		i++
		k.L.Printf("--- round %d ---", i)
		k.L.Printf("Player A: %d", playerA)
		k.L.Printf("Player B: %d", playerB)

		// Check if this configuration has occurred before
		deckState := fmt.Sprintf("Game %d, Player A: %d; Player B: %d", gameNo, playerA, playerB)
		if _, ok := k.DeckHistory[deckState]; ok {
			k.L.Printf("I've seen this configuration before in this game! A wins.")
			return true, -1
		}
		k.DeckHistory[deckState] = true

		cardA := takeTopCard(&playerA)
		cardB := takeTopCard(&playerB)
		k.L.Printf("A gets %d, B gets %d", cardA, cardB)

		if len(playerA) >= cardA && len(playerB) >= cardB {
			k.L.Printf("Playing a sub-game to figure out the winner... ")

			aWins, score = k.Combat(playerA[:cardA], playerB[:cardB])

			k.L.Printf("... but back to game %d", gameNo)
		} else {
			aWins = cardA > cardB
		}

		if aWins {
			k.L.Printf("Player A wins round %d of game %d", i+1, gameNo)
			playerA = append(playerA, cardA, cardB)
		} else {
			k.L.Printf("Player B wins round %d of game %d", i+1, gameNo)
			playerB = append(playerB, cardB, cardA)
		}
	}

	k.L.Printf("Player A: %d", playerA)
	k.L.Printf("Player B: %d", playerB)

	score = combatScore(playerA) + combatScore(playerB)
	aWins = len(playerA) > 0
	return
}
