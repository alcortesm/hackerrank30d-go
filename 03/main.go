package main

import (
	"fmt"
	"log"
)

func main() {
	n, err := readInt()
	if err != nil {
		log.Fatal(err)
	}

	if isWeird(n) {
		fmt.Println("Weird")
	} else {
		fmt.Println("Not Weird")
	}
}

func readInt() (int, error) {
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		return 0, err
	}

	if isInvalid(n) {
		return 0, fmt.Errorf("invalid input, should be >= 1 and <=100, but was %d\n", n)
	}

	return n, nil
}

func isInvalid(n int) bool {
	return n < 1 || n > 100
}

func isWeird(n int) bool {
	return isOdd(n) || (n >= 6 && n <= 20)
}

func isOdd(n int) bool {
	return n%2 != 0
}
