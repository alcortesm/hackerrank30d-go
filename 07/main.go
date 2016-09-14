package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	n, err := readArraySize()
	if err != nil {
		log.Fatalf("invalid input at line 1: %s", err)
	}

	a, err := readArray(n)
	if err != nil {
		log.Fatal(err)
	}

	printReverse(a)
}

func readArraySize() (int, error) {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return 0, err
	}
	err := checkArraySize(n)
	return n, err
}

func checkArraySize(n int) error {
	if n < 1 || n > 1000 {
		return fmt.Errorf("invalid array size: 1 <= T <= 1000")
	}
	return nil
}

func readArray(n int) ([]int, error) {
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			return nil, fmt.Errorf("error reading data: %s", scanner.Err())
		}
		word := scanner.Text()
		if a[i], err = strconv.Atoi(word); err != nil {
			return nil, fmt.Errorf("error reading word %d as an integer: %s", i+1, err)
		}
		if err := checkElement(a[i]); err != nil {
			return nil, fmt.Errorf("invalid value for element %d (%d): %s", i+1, a[i], err)
		}
	}

	return a, nil
}

func checkElement(e int) error {
	if e < 1 {
		return fmt.Errorf("to small")
	}
	if e > 10000 {
		return fmt.Errorf("to big")
	}
	return nil
}

func printReverse(a []int) {
	var buf bytes.Buffer
	sep := ""
	for i := len(a) - 1; i >= 0; i-- {
		buf.WriteString(sep)
		buf.WriteString(strconv.Itoa(a[i]))
		sep = " "
	}
	fmt.Println(buf.String())
}
