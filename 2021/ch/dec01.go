package ch

import "fmt"

func Dec01a(args []string) error {
	fmt.Printf("Args: %d (%v)\n", len(args), args)
	depths, err := dataAsInts(args, "inputs/dec01a.txt")
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

	fmt.Printf("%d\n", rv)
	return nil
}
