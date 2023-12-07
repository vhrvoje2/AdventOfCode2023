package main

import (
	"adventofcode2023/utils"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"unicode"
)

const (
	sampleInput1 = "test_input1.txt"
	sampleInput2 = "test_input2.txt"
	puzzleInput  = "input.txt"
)

type Day struct{}

func (d Day) strNums() map[string]string {
	m := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}

	return m
}

func main() {
	day := Day{}
	solution1 := day.Part1(puzzleInput)
	fmt.Printf("Part 1 solution %d\n", solution1)

	solution2 := day.Part2(puzzleInput)
	fmt.Printf("Part 2 solution %d\n", solution2)
}

func (d Day) Part1(filename string) int {
	input, err := utils.ReadInput(filename)

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	values := make([]string, 0)
	solution := 0

	for _, line := range input {
		result := ""
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

func (d Day) Part2(filename string) int {
	input, err := utils.ReadInput(filename)

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	solution := 0
	leftIdx := math.MaxInt32
	rightIdx := -1
	leftNum := ""
	rightNum := ""

	for _, line := range input {
		for key, val := range d.strNums() {
			leftKeyIdx := strings.Index(line, key)
			rightKeyIdx := strings.LastIndex(line, key)
			if leftKeyIdx > -1 && leftKeyIdx < leftIdx {
				leftIdx = leftKeyIdx
				leftNum = val
			}
			if rightKeyIdx > -1 && rightKeyIdx > rightIdx {
				rightIdx = rightKeyIdx
				rightNum = val
			}
		}

		res := leftNum + rightNum
		intVal, _ := strconv.Atoi(res)
		solution += intVal

		leftIdx = math.MaxInt32
		rightIdx = -1
	}

	return solution
}
