package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func main() {

	var numbers = make([]int, 0, 3)
	for {
		fmt.Print("Write a number or X to quit: ")
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Fatal(err)
		}
		if input == "X" {
			break
		}
		i, err := strconv.Atoi(input)
		if err != nil {
			// handle error
			fmt.Println(input, "is not a number!")
			continue
		}
		numbers = append(numbers, i)
		sort.Ints(numbers)
		fmt.Println("Sorted slice:", numbers)
	}
	fmt.Println("Bye...")
}
