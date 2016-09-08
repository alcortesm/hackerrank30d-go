package main

import (
	"fmt"
	"log"
)

func main() {
	var n int
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		log.Fatal(err)
	}

	if isInvalid(n) {
		log.Fatalf("input should be >= 1 and <=100, but was %d\n", n)
	}

	if isOdd(n) {
		fmt.Println("Weird")
		return
	}

	switch {
	case n >= 2 && n <= 5:
		fmt.Println("Not Weird")
	case n >= 6 && n <= 20:
		fmt.Println("Weird")
	case n > 20:
		fmt.Println("Not Weird")
	}
}

func isInvalid(n int) bool {
	return n < 1 || n > 100
}

func isOdd(n int) bool {
	return n%2 != 0
}
