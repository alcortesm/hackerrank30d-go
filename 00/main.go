package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello, World.")
	fmt.Print(line)
}
