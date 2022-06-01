package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var str string
	fmt.Println("Write a string:")
	_, err := fmt.Scan(&str)
	if err != nil {
		log.Fatal(err)
	}
	str = strings.ToLower(str)
	var answer string
	if strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str, "a") {
		answer = "Found!"
	} else {
		answer = "Not Found!"
	}
	fmt.Println(answer)

	s := strings.Replace("ianianian", "ni", "in", 2)
	fmt.Println(s)
}
