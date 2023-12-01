package main

import (
	"adventofcode2023/utils"
	"fmt"
	"log"
	"strconv"
	"unicode"
)

const (
	sampleInput = "test_input.txt"
	puzzleInput = "input.txt"
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

	values := make([]string, 0)

	for _, line := range input {
		var result string
		for _, char := range line {
			if !unicode.IsLetter(char) {
				result += string(char)
			}
		}
		trimmed := result[:1] + result[len(result)-1:]
		values = append(values, trimmed)
	}

	for _, val := range values {
		intVal, _ := strconv.Atoi(val)
		solution += intVal
	}

	return solution
}

func (d Day1) Part2(filename string) int {
	solution, err := utils.ReadInput(filename)

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	fmt.Println("Part 2 solution")
	fmt.Println(solution)
}
