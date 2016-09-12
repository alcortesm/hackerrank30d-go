package main

import (
	"fmt"
	"log"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	if err := check(n); err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}

func check(n int) error {
	if n < 2 || n > 20 {
		return fmt.Errorf("invalid n: 2 <= n <= 20")
	}
	return nil
}
