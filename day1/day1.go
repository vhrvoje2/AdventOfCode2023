package main

import (
	"adventofcode2023/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
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

	/* solution2 := day1.Part2(puzzleInput) */
	solution2 := day1.Part2(sampleInput2)
	fmt.Printf("Part 2 solution %d\n", solution2)

}

func (d Day1) Part1(filename string) int {
	input, err := utils.ReadInput(filename)

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	return d.adderHelper(input)
}

func (d Day1) Part2(filename string) int {
	input, err := utils.ReadInput(filename)
	strNums := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	var newInput []string
	var newLine string

	for _, line := range input {
		newLine = line
		for strNum, strNumVal := range strNums {
			if strings.Contains(newLine, strNum) {
				newLine = strings.Replace(newLine, strNum, strNumVal, -1)
			}
		}

		newInput = append(newInput, newLine)
	}

	return d.adderHelper(newInput)
}

func (d Day1) adderHelper(input []string) int {
	values := make([]string, 0)
	solution := 0

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
