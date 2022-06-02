package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	animalsMap := map[string]*Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	for {
		fmt.Print("> ")
		input := readStdIn()

		val, hasKey := animalsMap[input[0]]
		if len(input) != 2 || !hasKey || (input[1] != "eat" && input[1] != "move" && input[1] != "speak") {
			fmt.Println("I don't know what you're talking about")
			continue
		}
		switch input[1] {
		case "eat":
			val.Eat()
		case "move":
			val.Move()
		case "speak":
			val.Speak()
		}
	}

}

func readStdIn() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(scanner.Text(), " ")
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.sound)
}

type Animal struct {
	food, locomotion, sound string
}
