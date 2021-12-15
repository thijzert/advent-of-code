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

var allFuncs []ch.AdventFunc

func init() {
	allFuncs = []ch.AdventFunc{
		aoc20.Dec09a,
		aoc20.Dec09b,
		aoc20.Dec10a,
		aoc20.Dec10b,
		aoc20.Dec11a,
		aoc20.Dec11b,
		aoc20.Dec12a,
		aoc20.Dec12b,
		aoc20.Dec13a,
		aoc20.Dec13b,
		aoc20.Dec14a,
		aoc20.Dec14b,
		aoc20.Dec15a,
		aoc20.Dec15b,
		aoc20.Dec16a,
		aoc20.Dec16b,
		aoc20.Dec17a,
		aoc20.Dec17b,
		aoc20.Dec18a,
		aoc20.Dec18b,
		aoc20.Dec19a,
		aoc20.Dec19b,
		aoc20.Dec20a,
		aoc20.Dec20b,
		aoc20.Dec21a,
		aoc20.Dec21b,
		aoc20.Dec22a,
		aoc20.Dec22b,
		aoc20.Dec23a,
		aoc20.Dec23b,
		aoc20.Dec24a,
		aoc20.Dec24b,
		aoc20.Dec25a,

		aoc21.Dec01a,
		aoc21.Dec01b,
		aoc21.Dec02a,
		aoc21.Dec02b,
		aoc21.Dec03a,
		aoc21.Dec03b,
		aoc21.Dec04a,
		aoc21.Dec04b,
		aoc21.Dec05a,
		aoc21.Dec05b,
		aoc21.Dec06a,
		aoc21.Dec06b,
		aoc21.Dec07a,
		aoc21.Dec07b,
		aoc21.Dec08a,
		aoc21.Dec08b,
		aoc21.Dec09a,
		aoc21.Dec09b,
		aoc21.Dec10a,
		aoc21.Dec10b,
		aoc21.Dec11a,
		aoc21.Dec11b,
		aoc21.Dec12a,
		aoc21.Dec12b,
		aoc21.Dec13a,
		aoc21.Dec13b,
		aoc21.Dec14a,
		aoc21.Dec14b,
		aoc21.Dec15a,
		aoc21.Dec15b,
	}
}

func main() {
	var funcIdx int = len(allFuncs) - 1
	var quiet bool

	flag.IntVar(&funcIdx, "f", funcIdx, fmt.Sprintf("Index to challenge (0-%d)", funcIdx))
	flag.BoolVar(&quiet, "q", false, "Suppress debug output")
	flag.Parse()

	var debugOut io.Writer = os.Stdout
	if quiet {
		debugOut = io.Discard
	}

	ctx := ch.AOContext{
		Ctx:         context.Background(),
		Args:        flag.Args(),
		Debug:       log.New(debugOut, "", log.Lshortfile),
		FinalAnswer: log.New(os.Stdout, "final answer: ", log.Lshortfile),
	}

	err := allFuncs[funcIdx](ctx)
	if err != nil {
		log.Fatal(err)
	}
}
