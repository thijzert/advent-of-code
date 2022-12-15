package ch

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/thijzert/advent-of-code/data"
)

type AOContext struct {
	Ctx   context.Context
	Args  []string
	Debug *log.Logger
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

var answersInData map[int][50]string

func getAnswersInData(year int) [50]string {
	if answersInData == nil {
		answersInData = make(map[int][50]string)
	}
	var rv [50]string
	var ok bool
	if rv, ok = answersInData[year]; ok {
		return rv
	}

	lines, err := data.GetLines("results.txt")
	if err != nil {
		log.Fatalf("Error opening \"results.txt\": %v", err)
	}
	for _, line := range lines {
		var y, d int
		var c rune
		_, err := fmt.Sscanf(line, "%d %d-%c:", &y, &d, &c)
		if err != nil || y != year || d < 1 || d > 25 {
			continue
		}
		i := 2 * (d - 1)
		if c == 'B' {
			i++
		}
		j := strings.IndexByte(line, ':')
		rv[i] = strings.TrimSpace(line[j+1:])
	}

	answersInData[year] = rv
	return rv
}

func (ctx AOContext) CheckAnswer(year, challengeIndex int, ans interface{}) error {
	if ans == nil {
		return fmt.Errorf("empty answer")
	}

	answers := getAnswersInData(year)
	ex := answers[challengeIndex]
	if ex == "" {
		// TBD
		return nil
	}
	ob := fmt.Sprint(ans)
	if ob != ex {
		return errIncorrect{
			Expected: ex,
			Observed: ob,
		}
	}
	return nil
}

type errIncorrect struct {
	Expected, Observed string
}

func (e errIncorrect) Error() string {
	return fmt.Sprintf("incorrect answer: expected '%s', got '%s'", e.Expected, e.Observed)
}

type AdventFunc func(AOContext) (interface{}, error)

type Advent [50]AdventFunc

func (ad Advent) Stars(ctx AOContext, year int, answerOut io.Writer, onError func(error)) string {
	rv := ""

	for i := 0; i < 50; i += 2 {
		published := false
		n := 0
		for j := 0; j < 2; j++ {
			if ad[i+j] == nil {
				continue
			}
			published = true
			ans, err := ad[i+j](ctx)
			if ans != nil {
				fmt.Fprintf(answerOut, "%d-%d-%c: %v\n", year, 1+i/2, 'A'+rune(j), ans)
			}
			if err != nil {
				onError(err)
			} else {
				err = ctx.CheckAnswer(year, i+j, ans)
				if err != nil {
					onError(err)
				} else {
					n++
				}
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
