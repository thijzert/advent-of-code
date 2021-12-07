package ch

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"

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
			continue
		}
		rv = append(rv, i)
	}
	return rv, nil
}

func (ctx AOContext) DataAsIntLists(assetName string) ([][]int, error) {
	lines, err := ctx.DataLines(assetName)
	if err != nil {
		return nil, err
	}

	var rv [][]int
	for _, line := range lines {
		strs := strings.Split(line, ",")
		iline := make([]int, 0, len(strs))
		for _, str := range strs {
			i, err := strconv.Atoi(str)
			if err != nil {
				continue
			}
			iline = append(iline, i)
		}
		rv = append(rv, iline)
	}
	return rv, nil
}

func (ctx AOContext) Print(v ...interface{}) {
	ctx.Debug.Print(v...)
}

func (ctx AOContext) Printf(format string, v ...interface{}) {
	ctx.Debug.Printf(format, v...)
}

type AdventFunc func(AOContext) error

var errNotImplemented = errors.New("not implemented")

func ExampleChallenge(ctx AOContext) error {
	return errNotImplemented
}
