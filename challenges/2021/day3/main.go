package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var words = "./input2.txt"
	input := importFile(words)

	// need 2 arrays for epsilon and gamma
	// input array already made
	// a map to count total #'s of 1's and 0's  countThis["0s"] .. countThis["1s"]
	// one last array of the number totals
	// convert binary to decimal (probably just import something for this ;D)
	var gammaBin = make([]string, 0)

	for i := 0; i < len(input); i++ {
		// input[i][i : i+1]
		numZeros := 0
		numOnes := 0
		for j := 0; j < len(input[i]); j++ {

			if input[i][j:j+1] == "0" {

				numZeros++
			}
			if input[i][j:j+1] == "1" {

				numOnes++
			}
		}
		if numZeros > numOnes {
			gammaBin = append(gammaBin, "0")
		} else {
			gammaBin = append(gammaBin, "1")
		}
	}

	fmt.Printf("%s", gammaBin)
	day2()
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

func day2() {

}
