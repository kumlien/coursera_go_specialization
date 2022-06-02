package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	animals := map[string]Animal{}

	for {
		fmt.Println("Enter 'newanimal' or 'query'")
		args := readStdIn()
		switch args[0] {
		case "newanimal":
			animal, err := createAnimal(args[2])
			if err != nil {
				fmt.Println(err)
				continue
			}
			animals[args[1]] = animal
			fmt.Println("Created it!")
		case "query":
			a, exists := animals[args[1]]
			if !exists {
				fmt.Println("No animal with name", args[1])
				continue
			}
			switch args[2] {
			case "eat":
				a.Eat()
			case "move":
				a.Move()
			case "speak":
				a.Speak()
			default:
				fmt.Println("Unsupported action:", args[2], ". Supported actions are eat, move and speak")
			}
		default:
			fmt.Println("Unsupported action, use newanimal or query")
		}
	}
}

func createAnimal(typeOfAnimal string) (Animal, error) {
	var animal Animal
	var err error = nil
	switch typeOfAnimal {
	case "bird":
		animal = &Bird{}
	case "snake":
		animal = &Snake{}
	case "cow":
		animal = &Cow{}
	default:
		err = errors.New("Unsupported type: " + typeOfAnimal + ". Supported types: bird, snake, cow")
	}
	return animal, err
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

func (a *Bird) Eat() {
	fmt.Println("worms")
}
func (a *Bird) Move() {
	fmt.Println("fly")
}
func (a *Bird) Speak() {
	fmt.Println("peep")
}

func (a *Snake) Eat() {
	fmt.Println("mice")
}
func (a *Snake) Move() {
	fmt.Println("slither")
}
func (a *Snake) Speak() {
	fmt.Println("hzzz")
}

func (a *Cow) Eat() {
	fmt.Println("grass")
}
func (a *Cow) Move() {
	fmt.Println("walk")
}
func (a *Cow) Speak() {
	fmt.Println("moo")
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

type Animal interface {
	Eat()
	Move()
	Speak()
}
