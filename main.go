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

	counter2 := 0
	var inputsEveryThree = make([]int, 0)
	outerCounter := 0
	threeCounter := 1

	for i := 0; i < len(input); i++ {

		if threeCounter == 3 {
			inputsEveryThree[outerCounter] = 0
			sum, _ := strconv.Atoi(input[i])
			sum2, _ := strconv.Atoi(input[i-1])
			sum3, _ := strconv.Atoi(input[i-2])

			inputsEveryThree[outerCounter] += (sum + sum2 + sum3)

			outerCounter++
			threeCounter = 0
		}
		threeCounter++

	}
	fmt.Println(inputsEveryThree)
	fmt.Println("Day 2: " + strconv.Itoa(counter2))

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
