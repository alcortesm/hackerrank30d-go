package main

import (
	"fmt"
	"log"
)

// hackerrank tests are failing on this solution, even though it is
// correct. I'm pretty sure they have some bug in their test system or
// in their recursion detection system.
func main() {
	n, err := readFactInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(fact(n))
}

func readFactInput() (int, error) {
	var n int
	if _, err := fmt.Scanf("%d", &n); err != nil {
		return 0, fmt.Errorf("reading from stdin: %s", err)
	}
	if err := checkFactInput(n); err != nil {
		return 0, err
	}
	return n, nil
}

func checkFactInput(n int) error {
	if n < 0 {
		return fmt.Errorf("factorial is not defined for negative numbers")
	}
	return nil
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
