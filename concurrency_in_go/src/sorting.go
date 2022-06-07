package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter numbers separated by space:")
	array := readStdIn()
	slices := split(array)
	fmt.Println("sub arrays, unsorted:", slices)
	channel := make(chan []int)
	for _, slice := range slices {
		go sortMe(slice, channel)
	}

	arr1 := <-channel
	arr2 := <-channel
	arr3 := <-channel
	arr4 := <-channel
	fmt.Println("sub arrays sorted = ", arr1, arr2, arr3, arr4)

	arr12 := merge(arr1, arr2)
	arr34 := merge(arr3, arr4)
	result := merge(arr12, arr34)

	fmt.Println("merged and sorted sub arrays = ", result)
}

func readStdIn() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(scanner.Text(), " ")
	slice := make([]int, len(input))
	for i, v := range input {
		slice[i], _ = strconv.Atoi(v)
	}
	return slice
}

func sortMe(slice []int, channel chan []int) {
	fmt.Println("goroutine sorting:", slice)
	if len(slice) <= 1 {
		channel <- slice
		return
	}
	sort.Ints(slice)
	channel <- slice
}

func split(slice []int) [][]int {
	arr1, arr2, arr3, arr4 := []int{}, []int{}, []int{}, []int{}
	for i := 0; i < len(slice); i += 1 {
		switch i % 4 {
		case 0:
			arr1 = append(arr1, slice[i])
		case 1:
			arr2 = append(arr2, slice[i])
		case 2:
			arr3 = append(arr3, slice[i])
		case 3:
			arr4 = append(arr4, slice[i])
		}
	}
	return [][]int{arr1, arr2, arr3, arr4}
}

func merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}
