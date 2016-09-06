package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var i uint32 = 4
	var d float32 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewReader(os.Stdin)

	// The above code is pretty limiting (in terms of available imports)
	// and quite confusing as there is no need to buffer the stdin. As a
	// consecuence, the code below fells awkard.

	var ii uint32
	if _, err := fmt.Scanf("%d\n", &ii); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var dd float32
	if _, err := fmt.Scanf("%f\n", &dd); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// read a string in chunks, no matter how long it is, as long as
	// there is enough memory for it.
	chunks := [][]byte{}
	for {
		part, moreInput, err := scanner.ReadLine()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		clone := make([]byte, len(part))
		copy(clone, part)
		chunks = append(chunks, clone)

		if !moreInput {
			break
		}
	}

	fmt.Printf("%d\n", i+ii)
	fmt.Printf("%.1f\n", d+dd)
	fmt.Print(s)
	for _, c := range chunks {
		fmt.Print(string(c))
	}
	fmt.Println()
}
