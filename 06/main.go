package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	n, err := readNumber()
	if err != nil {
		log.Fatalf("invalid input at line 1: %s", err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		line, err := readLine(scanner)
		if err != nil {
			log.Fatalf("line %d: %s", i+1, err)
		}

		printEvenOdd(line)
	}

}

func readNumber() (int, error) {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return 0, err
	}
	err := checkNumber(n)
	return n, err
}

func checkNumber(n int) error {
	if n < 1 || n > 10 {
		return fmt.Errorf("invalid T: 1 <= T <= 10")
	}
	return nil
}

// Default scanner buffer size (bufio.MaxScanTokenSize) is big
// enough for the problem constraints (len(S) <= 10000)
func readLine(s *bufio.Scanner) (string, error) {
	if !s.Scan() {
		if err := s.Err(); err != nil {
			return "", fmt.Errorf("error reading line: %s", err)
		} else {
			return "", fmt.Errorf("unexpected EOF")
		}
	}
	line := s.Text()
	if err := checkLine(line); err != nil {
		return "", fmt.Errorf("invalid line: %s", err)
	}
	return line, nil
}

func checkLine(s string) error {
	if len(s) < 2 {
		return fmt.Errorf("too short")
	}
	if len(s) > 10000 {
		return fmt.Errorf("too long")
	}
	return nil
}

func printEvenOdd(s string) {
	var even, odd []rune
	nextIsEven := true
	for _, r := range s {
		if nextIsEven {
			even = append(even, r)
			nextIsEven = false
		} else {
			odd = append(odd, r)
			nextIsEven = true
		}
	}
	fmt.Println(string(even), string(odd))
}
