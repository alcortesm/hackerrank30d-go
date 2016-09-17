package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	squareSize = 6
	minValue   = -9
	maxValue   = 9
)

func main() {
	a, err := readSquare(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	hgs := a.getAllHgs()
	ss := sums(hgs)
	m := max(ss)

	fmt.Println(m)
}

type square [][]int

func readSquare(r io.Reader) (square, error) {
	s := bufio.NewScanner(r)

	n := squareSize
	rows := make([][]int, n)

	for r := 0; r < n; r++ {
		rows[r] = make([]int, n)
		if err := scanRow(s, rows[r]); err != nil {
			return nil, fmt.Errorf(
				"scanning row %d: %s", r+1, err)
		}
	}

	return rows, nil
}

func scanRow(s *bufio.Scanner, row []int) error {
	if !s.Scan() {
		return s.Err()
	}
	words := strings.Split(s.Text(), " ")
	if len(words) != len(row) {
		return fmt.Errorf(
			"unexpected number of elements, %d was expected, but %d were found",
			len(row), len(words))
	}
	for i, w := range words {
		n, err := strconv.Atoi(w)
		if err != nil {
			return fmt.Errorf(
				"element %d: invalid format: %s",
				i+1, err)
		}
		if n < minValue {
			return fmt.Errorf(
				"element %d: invalid value (%d): too small, minimum is -9",
				i+1, n)
		}
		if n > maxValue {
			return fmt.Errorf(
				"element %d: invalid value (%d): too big, maximum is 9",
				i+1, n)
		}
		row[i] = n
	}
	return nil
}

func (s square) getAllHgs() [][]int {
	hgs := make([][]int, 0)
	for i := 0; i < len(s)-2; i++ {
		for j := 0; j < len(s)-2; j++ {
			hgs = append(hgs, s.getHg(i, j))
		}
	}
	return hgs
}

func (s square) getHg(i, j int) []int {
	ret := make([]int, 7)
	ret[0] = s[i][j]
	ret[1] = s[i][j+1]
	ret[2] = s[i][j+2]
	ret[3] = s[i+1][j+1]
	ret[4] = s[i+2][j]
	ret[5] = s[i+2][j+1]
	ret[6] = s[i+2][j+2]
	return ret
}

func sums(ll [][]int) []int {
	ret := make([]int, len(ll))
	for i, l := range ll {
		for _, e := range l {
			ret[i] += e
		}
	}
	return ret
}

func max(l []int) int {
	if len(l) == 0 {
		panic("there is no max element in an empty slice")
	}
	m := l[0]
	for i := 1; i < len(l); i++ {
		if l[i] > m {
			m = l[i]
		}
	}
	return m
}
