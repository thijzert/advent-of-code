package ch

import (
	"context"
	"log"
	"strconv"

	"github.com/thijzert/advent-of-code/2021/data"
)

type AOContext struct {
	Ctx         context.Context
	Args        []string
	Debug       *log.Logger
	FinalAnswer *log.Logger
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
