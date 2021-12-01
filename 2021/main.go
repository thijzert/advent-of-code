package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/thijzert/advent-of-code/2021/ch"
)

var allFuncs []ch.AdventFunc

func init() {
	allFuncs = []ch.AdventFunc{
		ch.Dec01a,
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
