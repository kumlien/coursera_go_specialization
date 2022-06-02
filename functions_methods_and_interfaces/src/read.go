package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var fileName string
	fmt.Print("Please enter the filename: ")
	fileName = readStdIn()
	lines := readFile(fileName)
	fullNames := make([]FullName, 0, len(lines))
	for i := 0; i < len(lines); i++ {
		line := strings.Fields(lines[i])
		fullName := FullName{fname: line[0], lname: line[1]}
		fullNames = append(fullNames, fullName)
	}
	for _, fullName := range fullNames {
		fmt.Printf("fname: [%-20s]; lname: [%-20s]\n", fullName.fname, fullName.lname)
	}

}

func readFile(filePath string) []string {
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	defer readFile.Close()

	return fileLines
}

func readStdIn() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}

type FullName struct {
	fname string
	lname string
}
