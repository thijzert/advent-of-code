package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/ch/aoc20"
	"github.com/thijzert/advent-of-code/ch/aoc21"
)

func main() {
	minYear := 2020
	allYears := []ch.Advent{
		aoc20.Advent,
		aoc21.Advent,
	}

	var yearParam int = -1
	var yearIdx int = -1
	var funcIdx int = -1
	var quiet bool
	var veryQuiet bool
	var runAll bool

	flag.IntVar(&yearParam, "y", yearParam, "Year")
	flag.IntVar(&funcIdx, "f", funcIdx, "Index to challenge (0-49)")
	flag.BoolVar(&quiet, "q", false, "Suppress debug output")
	flag.BoolVar(&veryQuiet, "qq", false, "Suppress most output")
	flag.BoolVar(&runAll, "a", false, "Run all challenges")
	flag.Parse()

	yearIdx = yearParam
	if yearIdx == -1 {
		yearIdx = len(allYears) - 1
	} else if yearIdx >= len(allYears) {
		yearIdx -= minYear
	}

	if funcIdx == -1 {
		for i := range allYears[yearIdx] {
			if f := allYears[yearIdx][49-i]; f != nil {
				funcIdx = 49 - i
				break
			}
		}
	}

	var answerOut io.Writer = os.Stdout
	var debugOut io.Writer = os.Stdout
	if quiet || veryQuiet {
		debugOut = io.Discard
	}
	if veryQuiet {
		answerOut = io.Discard
	}

	ctx := ch.AOContext{
		Ctx:         context.Background(),
		Args:        flag.Args(),
		Debug:       log.New(debugOut, "", log.Lshortfile),
		FinalAnswer: log.New(answerOut, "final answer: ", log.Lshortfile),
	}

	if runAll {
		if yearParam != -1 {
			allYears = allYears[yearIdx : yearIdx+1]
			minYear += yearIdx
		}

		exitStatus := 0
		for i, ann := range allYears {
			if !veryQuiet {
				fmt.Printf("%d: \n", minYear+i)
			}
			st := ann.Stars(ctx, func(err error) {
				log.Print(err)
				exitStatus = 1
			})
			fmt.Printf("%d: %s\n", minYear+i, st)
		}
		os.Exit(exitStatus)
	} else {
		err := allYears[yearIdx][funcIdx](ctx)
		if err != nil {
			log.Fatal(err)
		}
	}
}
