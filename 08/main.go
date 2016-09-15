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
	minMapSize    = 1
	maxMapSize    = 100000
	minNumQueries = 1
	maxNumQueries = 100000
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n, err := readMapSize(scanner)
	if err != nil {
		log.Fatalf("invalid input at line 1: %s", err)
	}

	dict, err := readMap(scanner, n)
	if err != nil {
		log.Fatal("reading address book contents: %s", err)
	}

	err = readQuestionWriteAnswerLoop(scanner, dict)
	if err != nil {
		log.Fatalf("reading queries: %s", err)
	}
}

func readMapSize(s *bufio.Scanner) (int, error) {
	if !s.Scan() {
		return 0, s.Err()
	}
	n, err := strconv.Atoi(s.Text())
	if err != nil {
		return 0, err
	}
	err = checkMapSize(n)
	return n, err
}

func checkMapSize(n int) error {
	if n < minMapSize {
		return fmt.Errorf(
			"map size too small, min is %d, but %d was requested",
			minMapSize, n)
	}
	if n > maxMapSize {
		return fmt.Errorf(
			"map size too big, max is %d, but %d was requested",
			maxMapSize, n)
	}
	return nil
}

func readMap(scanner *bufio.Scanner, n int) (map[string]string, error) {
	dict := make(map[string]string, n)
	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			return nil, fmt.Errorf(
				"error reading data at line %d: %s", i+1, scanner.Err())
		}
		line := scanner.Text()
		words := strings.Split(line, " ")
		if len(words) != 2 {
			return nil, fmt.Errorf(
				"invalid input at line %d: two words were expected, but %d was found",
				i+1, len(words))
		}
		dict[words[0]] = words[1]
	}

	return dict, nil
}

func readQuestionWriteAnswerLoop(s *bufio.Scanner, dict map[string]string) error {
	for s.Scan() {
		query := s.Text()
		answer := answerQuery(query, dict)
		fmt.Println(answer)
	}

	if err := s.Err(); err != nil && err != io.EOF {
		return err
	}
	return nil
}

func answerQuery(query string, dict map[string]string) string {
	answer, ok := dict[query]
	if !ok {
		return "Not found"
	}
	return query + "=" + answer
}
