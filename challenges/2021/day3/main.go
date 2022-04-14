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

	gammaBin := ""
	epsilonBin := ""

	for i := 0; i < len(input[0]); i++ {
		numZeros := 0
		numOnes := 0
		for j := 0; j < len(input); j++ {
			if input[j][i:i+1] == "0" {
				numZeros++
			}
			if input[j][i:i+1] == "1" {
				numOnes++
			}
		}
		if numZeros > numOnes {
			gammaBin += "0"
			epsilonBin += "1"
		} else {
			gammaBin += "1"
			epsilonBin += "0"
		}

	}

	gammaDec, _ := strconv.ParseInt(gammaBin, 2, 64)
	epsilonDec, _ := strconv.ParseInt(epsilonBin, 2, 64)
	fmt.Printf("%d\n", gammaDec)
	fmt.Printf("%d\n", epsilonDec)
	fmt.Printf("Part 1 answer: %d", gammaDec*epsilonDec)

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
	var words = "./input2.txt"
	input := importFile(words)

	oxyBin := ""
	c02Bin := ""

	for i := 0; i < len(input[0]); i++ {
		numZeros := 0
		numOnes := 0
		for j := 0; j < len(input); j++ {
			if input[j][i:i+1] == "0" {
				numZeros++
			}
			if input[j][i:i+1] == "1" {
				numOnes++
			}
		}

	}

	gammaDec, _ := strconv.ParseInt(oxyBin, 2, 64)
	epsilonDec, _ := strconv.ParseInt(c02Bin, 2, 64)
	fmt.Printf("%d\n", gammaDec)
	fmt.Printf("%d\n", epsilonDec)
	fmt.Printf("Part 1 answer: %d", gammaDec*epsilonDec)
}
