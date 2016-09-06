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
	// consecuence, the code below fells awkard. Also the tips in the
	// exercise don't allow to process the input as a stream, which can
	// lead to running out of memory if the input string is very long.

	var ii uint32
	if _, err := fmt.Scanf("%d\n", &ii); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%d\n", i+ii)

	var dd float32
	if _, err := fmt.Scanf("%f\n", &dd); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("%.1f\n", d+dd)

	fmt.Print(s)
	// We could substitute this whole for loop with a simple
	// io.Copy(os.Stdout, os.Stdin) if we were allowed to use
	// the io package.
	for {
		chunk, moreInput, err := scanner.ReadLine()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Print(string(chunk))

		if !moreInput {
			break
		}
	}

	fmt.Println()
}
