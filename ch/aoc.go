package ch

import (
	"context"
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

func (ctx AOContext) DataSections(assetName string) ([][]string, error) {
	lines := ctx.Args
	if len(ctx.Args) == 0 {
		var err error
		lines, err = data.GetLines(assetName)
		if err != nil {
			return nil, err
		}
	}

	rv := make([][]string, 1)
	for _, l := range lines {
		if l == "" {
			rv = append(rv, []string{})
		} else {
			rv[len(rv)-1] = append(rv[len(rv)-1], l)
		}
	}

	if len(rv[len(rv)-1]) == 0 {
		rv = rv[:len(rv)-1]
	}

	return rv, nil
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

type Advent [50]AdventFunc

func (ad Advent) Stars(ctx AOContext, onError func(error)) string {
	rv := ""

	for i := 0; i < 50; i += 2 {
		published := false
		n := 0
		for j := 0; j < 2; j++ {
			if ad[i+j] == nil {
				continue
			}
			published = true
			err := ad[i+j](ctx)
			if err != nil {
				onError(err)
			} else {
				n++
			}
		}
		if !published {
			rv += " "
		} else if n == 2 {
			rv += "\x1b[38;5;226m*\x1b[0m"
		} else if n == 1 {
			rv += "\x1b[38;5;252m*\x1b[0m"
		} else {
			rv += "\x1b[38;5;238m*\x1b[0m"
		}
	}

	return "[" + rv + "]"
}
