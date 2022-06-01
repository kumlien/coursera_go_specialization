package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	var data = make(map[string]string)
	var name string
	var address string
	fmt.Print("Please enter your name: ")
	name = read()

	fmt.Print("Please enter your address: ")
	address = read()

	data["name"] = name
	data["address"] = address

	asJson, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Data as json: ", string(asJson))
}

func read() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}
