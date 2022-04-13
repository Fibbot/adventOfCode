package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var words = "./input.txt"
	input := importFile(words)

	totals := make(map[string]int, 0)
	totals["forward"] = 0
	totals["up"] = 0
	totals["down"] = 0

	for i := 0; i < len(input); i++ {
		if strings.Contains(input[i], "forward") {
			split := strings.Split(input[i], " ")
			total, _ := strconv.Atoi(split[1])
			total += totals["forward"]
			totals["forward"] = total

		}
		if strings.Contains(input[i], "up") {
			split := strings.Split(input[i], " ")
			total, _ := strconv.Atoi(split[1])
			total += totals["up"]
			totals["up"] = total
		}
		if strings.Contains(input[i], "down") {
			split := strings.Split(input[i], " ")
			total, _ := strconv.Atoi(split[1])
			total += totals["down"]
			totals["down"] = total
		}

	}
	fmt.Println(totals)
	fmt.Printf("Total depth = %d\n", totals["down"]-totals["up"])
	fmt.Printf("Total horz position = %d\n", (totals["down"]-totals["up"])*totals["forward"])
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
	var words = "./input.txt"
	input := importFile(words)

	totals := make(map[string]int, 0)
	totals["forward"] = 0
	totals["up"] = 0
	totals["down"] = 0
	totals["aim"] = 0
	totals["depth"] = 0

	for i := 0; i < len(input); i++ {
		if strings.Contains(input[i], "forward") {
			split := strings.Split(input[i], " ")
			total, _ := strconv.Atoi(split[1])
			total2, _ := strconv.Atoi(split[1])
			total += totals["forward"]
			totals["forward"] = total
			if totals["aim"] != 0 {
				totals["depth"] = totals["depth"] + (totals["aim"] * total2)
			}

		}
		if strings.Contains(input[i], "up") {
			split := strings.Split(input[i], " ")
			total, _ := strconv.Atoi(split[1])
			totals["aim"] = totals["aim"] - total
		}
		if strings.Contains(input[i], "down") {
			split := strings.Split(input[i], " ")
			total, _ := strconv.Atoi(split[1])
			totals["aim"] = totals["aim"] + total
		}

	}
	fmt.Printf("PART2: Total depth = %d\n", totals["depth"])
	fmt.Printf("PART2: Total aim = %d\n", totals["aim"])
	fmt.Printf("PART2:Total horz position = %d\n", totals["forward"])
	fmt.Printf("Answer? : %d", (totals["depth"] * totals["forward"]))
}
