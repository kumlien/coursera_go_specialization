package main

import "fmt"

func main() {
	fmt.Printf("Print a number: ")
	var deci float64
	fmt.Scanln(&deci)
	var noDeci = int(deci)
	fmt.Println("As int:", noDeci)
}
