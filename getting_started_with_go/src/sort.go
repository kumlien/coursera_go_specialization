package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var slice = make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		fmt.Print("Enter int", i+1, ": ")
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Fatal(err)
		}
		integer, err := strconv.Atoi(input)
		if err != nil {
			// handle error
			fmt.Println(input, "is not a number, try again")
			i--
			continue
		}
		slice = append(slice, integer)
	}
	BubbleSort(slice)
	fmt.Println(slice)

}

func BubbleSort(sortMe []int) {
	for i := 0; i < len(sortMe); i++ {
		for j := 0; j < len(sortMe)-1; j++ {
			if sortMe[j] > sortMe[j+1] {
				Swap(sortMe, j)
			}
		}
	}
}

func Swap(slice []int, idx int) {
	slice[idx], slice[idx+1] = slice[idx+1], slice[idx]
}
