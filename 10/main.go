package main

import (
	"fmt"
	"log"
)

const (
	minNumber = 1
	maxNumber = 1000000
)

func main() {
	n, err := readNumber()
	if err != nil {
		log.Fatalf("failed reading number: %s", err)
	}
	if err := checkNumber(n); err != nil {
		log.Fatalf("invalid number value (%d): %s", n, err)
	}

	b := decToBin(n)
	count := maxConsecutive('1', b)
	fmt.Println(count)
}

func readNumber() (int, error) {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return 0, err
	}
	return n, nil
}

func checkNumber(n int) error {
	if n < minNumber {
		return fmt.Errorf("too small, minimum value is %d", minNumber)
	}
	if n > maxNumber {
		return fmt.Errorf("too big, maximum value is %d", maxNumber)
	}
	return nil
}

func decToBin(n int) string {
	return fmt.Sprintf("%b", n)
}

func maxConsecutive(r rune, s string) int {
	if len(s) == 0 {
		return 0
	}
	rr := []rune(s)

	max := 0
	prev := rr[0]
	consecutives := 0
	if prev == r {
		consecutives = 1
	}

	for i := 1; i < len(rr); i++ {
		if rr[i] == r {
			if rr[i] == prev {
				consecutives++
			} else {
				consecutives = 1
			}
		} else {
			consecutives = 0
		}
		prev = rr[i]
		if consecutives > max {
			max = consecutives
		}
	}
	return max
}
