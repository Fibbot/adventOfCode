package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var words = "./input.txt"
	input := importFile(words)

	counter := 1

	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			counter++
		}
	}
	fmt.Println("Day 1: " + strconv.Itoa(counter))

}

func importFile(input string) []string {
	file, err := os.Open(input)
	var inputList = make([]string, 0)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputList = append(inputList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return inputList
}
