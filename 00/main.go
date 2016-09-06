package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	fmt.Println("Hello, World.")
	fmt.Print(line)
}
