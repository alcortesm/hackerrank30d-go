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

	if isWeird(n) {
		fmt.Println("Weird")
	} else {
		fmt.Println("Not Weird")
	}
}

func isInvalid(n int) bool {
	return n < 1 || n > 100
}

func isOdd(n int) bool {
	return n%2 != 0
}

func isWeird(n int) bool {
	return isOdd(n) || (n >= 6 && n <= 20)
}
