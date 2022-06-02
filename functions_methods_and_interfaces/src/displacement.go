package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	acceleration := readStdIn("Enter acceleration: ")
	velocity := readStdIn("Enter velocity: ")
	displacement := readStdIn("Enter displacement: ")

	fn := GenDisplaceFn(acceleration, velocity, displacement)

	time := readStdIn("Enter time: ")
	fmt.Println("Displacement at time", time, ":", fn(time))

}

func GenDisplaceFn(acceleration, velocity, displacement float64) func(float64) float64 {
	return func(time float64) float64 {
		return (acceleration * time * time) + (velocity * time) + displacement
	}
}

func readStdIn(prompt string) float64 {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	val, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		log.Print("Computer says no: ", err)
		return readStdIn(prompt)
	} else {
		return val
	}
}
