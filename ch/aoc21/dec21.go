package aoc21

import (
	"github.com/thijzert/advent-of-code/ch"
)

func Dec21a(ctx ch.AOContext) (interface{}, error) {
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

	return die.NRolls() * min(playerScore...), nil
}

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

func Dec21b(ctx ch.AOContext) (interface{}, error) {
	const toWin int = 21
	const positions int = 10
	const players int = 2
	const minPerTurn int = 1
	const maxPerTurn int = positions
	const maxTurns int = 2*toWin/minPerTurn - 1
	const maxTotalScore int = 2*(toWin-1) + positions
	const maxPlayerScore int = toWin + positions - 1

	die := make([]int, 0, 27)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				die = append(die, i+j+k)
			}
		}
	}
	ctx.Printf("%d", die)

	var gameSpace [maxTurns + 1][maxPlayerScore + 1][maxPlayerScore + 1][positions][positions][players]uint64
	for i := 0; i <= maxTurns; i++ {
		for sW := toWin; sW <= maxPlayerScore; sW++ {
			for sL := 0; sL < toWin; sL++ {
				for a := 0; a < positions; a++ {
					for b := 0; b < positions; b++ {
						if i%2 == 0 {
							gameSpace[i][sL][sW][a][b][1] = 1
						} else {
							gameSpace[i][sW][sL][a][b][0] = 1
						}
					}
				}
			}
		}
	}

	for i := maxTurns; i > 0; i-- {
		for sA := 0; sA <= maxPlayerScore; sA++ {
			for sB := 0; sB <= maxPlayerScore; sB++ {
				for pos := 0; pos < positions; pos++ {
					psA := sA - pos - 1
					psB := sB - pos - 1

					for pO := 0; pO < positions; pO++ {
						for _, d := range die {
							prevPos := (pos + 20 - d) % positions

							if i%2 == 0 {
								if psB >= 0 && psB < toWin && sA < toWin {
									gameSpace[i-1][sA][psB][pO][prevPos][0] += gameSpace[i][sA][sB][pO][pos][0]
									gameSpace[i-1][sA][psB][pO][prevPos][1] += gameSpace[i][sA][sB][pO][pos][1]
								}
							} else {
								if psA >= 0 && psA < toWin && sB < toWin {
									gameSpace[i-1][psA][sB][prevPos][pO][0] += gameSpace[i][sA][sB][pos][pO][0]
									gameSpace[i-1][psA][sB][prevPos][pO][1] += gameSpace[i][sA][sB][pos][pO][1]
								}
							}
						}
					}
				}
			}
		}
	}

	ctx.Printf("There's no way this worked: %d / %d", gameSpace[0][0][0][3][7][0], gameSpace[0][0][0][3][7][1])
	ctx.Printf("There's no fucking way this worked: %d / %d", gameSpace[0][0][0][1][4][0], gameSpace[0][0][0][1][4][1])

	ctx.Printf("Why'd I need %.1f fucking megabytes of memory for this", float64(8*(maxTurns+1)*(maxPlayerScore+1)*(maxPlayerScore+1)*(positions)*(positions)*(players))/(1024.0*1024.0))
	return gameSpace[0][0][0][1][4][0], nil
}

func almostDec21b(ctx ch.AOContext) (interface{}, error) {
	const toWin int = 21
	const positions int = 10
	const players int = 2
	const maxTotalScore int = 2*(toWin-1) + positions
	const maxPlayerScore int = toWin + positions - 1

	die := make([]int, 0, 27)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				die = append(die, i+j+k)
			}
		}
	}

	var dirac [maxPlayerScore + 1][maxPlayerScore + 1][positions][positions][players]uint64

	// Set up: one player just won
	for sW := toWin; sW <= maxPlayerScore; sW++ {
		for sL := 0; sL < toWin; sL++ {
			// player won with final score sW, against score sL
			for pA := 0; pA < 10; pA++ {
				for pB := 0; pB < 10; pB++ {
					dirac[sW][sL][pA][pB][0] = 1
					dirac[sL][sW][pA][pB][1] = 1
				}
			}
		}
	}

	for sT := maxTotalScore; sT >= 0; sT-- {
		for sA := 0; sA <= maxPlayerScore; sA++ {
			sB := sT - sA
			if sB < 0 || sB > maxPlayerScore {
				continue
			}

			for pos := 0; pos < 10; pos++ {
				psA := sA - pos - 1
				psB := sB - pos - 1

				for pO := 0; pO < positions; pO++ {
					for _, d := range die {
						prevPos := (pos + 20 - d) % positions

						if psA >= 0 && psA < toWin && sB < toWin {
							dirac[psA][sB][prevPos][pO][1] += dirac[sA][sB][pos][pO][0]
						}
						if psB >= 0 && psB < toWin && sA < toWin {
							dirac[sA][psB][pO][prevPos][0] += dirac[sA][sB][pO][pos][1]
						}
					}
				}
			}
		}
	}

	ctx.Printf("There's no way this worked: %d / %d", dirac[0][0][3][7][0], dirac[0][0][3][7][1])
	ctx.Printf("There's no way this worked: %d / %d", dirac[0][0][7][3][0], dirac[0][0][7][3][1])
	for i, l := range dirac[0][0] {
		for j, st := range l {
			for a, v := range st {
				if v == 444356092776315 || v == 341960390180808 {
					ctx.Printf("a = %d, i = %d, j = %d", a, i, j)
				}
			}
		}
	}

	return nil, errNotImplemented
}

func alsoNotDec21b(ctx ch.AOContext) (interface{}, error) {
	const toWin int = 21
	const positions int = 10
	const players int = 2
	const maxTotalScore int = 2*(toWin-1) + positions
	const maxPlayerScore int = toWin + positions - 1

	die := make([]int, 0, 27)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				die = append(die, i+j+k)
			}
		}
	}

	var dirac [maxTotalScore + 1][maxTotalScore + 1][positions][positions][players]uint64

	// Set up: one player just won
	for sW := toWin; sW <= maxPlayerScore; sW++ {
		for sL := 0; sL < toWin; sL++ {
			// player won with final score sW, against score sL
			for pA := 0; pA < 10; pA++ {
				for pB := 0; pB < 10; pB++ {
					dirac[sW+sL][sW][pA][pB][1] = 1
					dirac[sW+sL][sW][pA][pB][0] = 1
				}
			}
		}
	}

	for sT := maxTotalScore; sT >= 0; sT-- {
		for sW := 0; sW <= sT; sW++ {
			sL := sT - sW
			for pW := 0; pW < positions; pW++ {
				// Previous score of the eventual winner
				psW := sW - pW - 1

				for pL := 0; pL < positions; pL++ {
					// Previous score of the eventual loser
					psL := sL - pL - 1

					for _, d := range die {
						ppA := (40 + pW - d) % 10
						ppB := (40 + pL - d) % 10

						if psW >= 0 && psW < toWin {
							dirac[psW+sL][psW][ppA][pL][0] += dirac[sT][sW][pW][pL][0]
							dirac[psW+sL][psW][ppA][pL][1] += dirac[sT][sW][pW][pL][1]
						}
						if psL >= 0 && psL < toWin {
							dirac[sW+psL][sW][pW][ppB][0] += dirac[sT][sW][pW][pL][0]
							dirac[sW+psL][sW][pW][ppB][1] += dirac[sT][sW][pW][pL][1]
						}
					}
				}
			}
		}
	}

	ctx.Printf("There's no way this worked: %d / %d", dirac[0][0][3][7][0], dirac[0][0][3][7][1])
	for i, l := range dirac[0][0] {
		for j, st := range l {
			for a, v := range st {
				if v == 444356092776315 || v == 341960390180808 {
					ctx.Printf("a = %d, i = %d, j = %d", a, i, j)
				}
			}
		}
	}

	return nil, errNotImplemented
}

func notdec21b(ctx ch.AOContext) (interface{}, error) {
	maxPlayerAScore := 30
	maxTotalScore := 50

	var dirac [51][51][10][10]uint64

	// Set up: player A just won
	for sA := maxPlayerAScore; sA >= 21; sA-- {
		for sB := 0; sB < 21; sB++ {
			for pA := 0; pA < 10; pA++ {
				if sA-pA-1 >= 21 {
					continue
				}

				for pB := 0; pB < 10; pB++ {
					dirac[sA+sB][sA][pA][pB] = 1
				}

			}
		}
	}

	for sT := maxTotalScore; sT >= 0; sT-- {
		for sA := 0; sA <= maxPlayerAScore; sA++ {
			sB := sT - sA
			if sB < 0 {
				continue
			}
			for pA := 0; pA < 10; pA++ {
				if sA-pA-1 < 0 {
					continue
				}

				// previous total score
				psT := sT - pA - 1

				for d1 := 1; d1 <= 3; d1++ {
					for d2 := 1; d2 <= 3; d2++ {
						for d3 := 1; d3 <= 3; d3++ {
							// previous A's position
							ppA := (pA + 40 - d1 - d2 - d3) % 10

							for pB := 0; pB < 10; pB++ {
								dirac[psT][sB][pB][ppA] += dirac[sT][sA][pA][pB]
							}
						}
					}
				}
			}
		}
	}

	ctx.Printf("There's no way this worked: %d / %d", dirac[0][0][3][7], dirac[0][0][7][3])

	return nil, errNotImplemented
}
