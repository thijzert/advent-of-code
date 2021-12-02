package ch

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/thijzert/advent-of-code/data"
)

type AOContext struct {
	Ctx         context.Context
	Args        []string
	Debug       *log.Logger
	FinalAnswer *log.Logger
}

func (ctx AOContext) DataLines(assetName string) ([]string, error) {
	if len(ctx.Args) == 0 {
		return data.GetLines(assetName)
	}

	return ctx.Args, nil
}

func (ctx AOContext) DataAsInts(assetName string) ([]int, error) {
	if len(ctx.Args) == 0 {
		return data.GetInts(assetName)
	}

	var rv []int
	for _, str := range ctx.Args {
		i, err := strconv.Atoi(str)
		if err != nil {
			return rv, err
		}
		rv = append(rv, i)
	}
	return rv, nil
}

type AdventFunc func(AOContext) error

var errNotImplemented = errors.New("not implemented")

func ExampleChallenge(ctx AOContext) error {
	return errNotImplemented
}
