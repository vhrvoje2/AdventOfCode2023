package main

import (
	"adventofcode2023/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	sampleInput1 = "test_input1.txt"
	sampleInput2 = "test_input2.txt"
	puzzleInput  = "input.txt"
)

type Day1 struct{}

func main() {
	day1 := Day1{}
	solution1 := day1.Part1(puzzleInput)
	fmt.Printf("Part 1 solution %d\n", solution1)

	solution2 := day1.Part2(puzzleInput)
	fmt.Printf("Part 2 solution %d\n", solution2)

}

func (d Day1) Part1(filename string) int {
	input, err := utils.ReadInput(filename)
	solution := 0

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	numbers := map[string]int{}

	for _, line := range input {
		card := strings.Split(line, ":")
		gameNumbers := strings.Replace(card[1], "|", "", -1)
		allNumbers := strings.Split(gameNumbers, " ")
		for _, num := range allNumbers {
			num = strings.Trim(num, " ")
			_, exists := numbers[num]
			if exists {
				numbers[num] += 1
			} else {
				numbers[num] = 1
			}
		}

		worth := 0

		for key, val := range numbers {
			_, err := strconv.Atoi(key)
			if val > 1 && err == nil {
				if worth == 0 {
					worth = 1
				} else {
					worth *= 2
				}
			}
		}

		solution += worth
		numbers = make(map[string]int)
	}

	return solution
}

func (d Day1) Part2(filename string) int {
	_, err := utils.ReadInput(filename)
	solution := 0

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	return solution
}
