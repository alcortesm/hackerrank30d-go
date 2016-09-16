package main

import (
	"fmt"
	"log"
)

func main() {
	n, err := readFactInput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(factorial(n))
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

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
