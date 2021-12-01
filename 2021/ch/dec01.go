package ch

func Dec01a(ctx AOContext) error {
	depths, err := ctx.DataAsInts("inputs/dec01a.txt")
	if err != nil {
		return err
	}

	rv := 0

	for i, depth := range depths {
		if i == 0 {
			continue
		}
		if depth > depths[i-1] {
			rv++
		}
	}

	ctx.FinalAnswer.Printf("%d\n", rv)
	return nil
}
