package main

import (
	"log"
	"os"
	"strconv"

	"github.com/thijzert/advent-of-code/2021/ch"
)

type AdventFunc func(args []string) error

var allFuncs []AdventFunc

func init() {
	allFuncs = []AdventFunc{
		ch.Dec01a,
	}
}

func main() {
	argIdx := len(os.Args)
	for i, arg := range os.Args {
		if arg == "--" {
			argIdx = i + 1
			break
		}
	}

	funcIdx := len(allFuncs) - 1
	if argIdx > 1 {
		if i, err := strconv.Atoi(os.Args[1]); err == nil {
			if i > 0 && i < len(allFuncs) {
				funcIdx = i
			}
		}
	}

	err := allFuncs[funcIdx](os.Args[argIdx:])
	if err != nil {
		log.Fatal(err)
	}
}
