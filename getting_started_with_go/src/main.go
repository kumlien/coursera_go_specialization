package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Printf("Hello, world")
}

func read2() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}
