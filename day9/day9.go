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

type Day struct{}

func main() {
	day := Day{}
	solution1 := day.Part1(puzzleInput)
	fmt.Printf("Part 1 solution %d\n", solution1)

	solution2 := day.Part2(puzzleInput)
	fmt.Printf("Part 2 solution %d\n", solution2)
}

func (d Day) Part1(filename string) int {
	input, err := utils.ReadInput(filename)
	solution := 0

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	for idx, _ := range input {
		history := []int{}
		nums := strings.Split(input[idx], " ")
		sequences := [][]int{}

		for _, num := range nums {
			intNum, _ := strconv.Atoi(num)
			history = append(history, intNum)
		}

		sequences = append(sequences, history)
		idx := 0

		for !d.allZero(sequences[len(sequences)-1]) {
			newSeq := []int{}

			for i := 0; i < len(sequences[idx])-1; i++ {
				newNum := sequences[idx][i+1] - sequences[idx][i]
				newSeq = append(newSeq, newNum)
			}

			sequences = append(sequences, newSeq)
			idx++
		}

		sequences[len(sequences)-1] = append(sequences[len(sequences)-1], 0)

		for i := len(sequences) - 1; i >= 1; i-- {
			lastNum := sequences[i][len(sequences[i])-1]
			previousNum := sequences[i-1][len(sequences[i-1])-1]
			newNum := lastNum + previousNum
			sequences[i-1] = append(sequences[i-1], newNum)
		}

		solution += sequences[0][len(sequences[0])-1]
	}

	return solution
}

func (d Day) Part2(filename string) int {
	input, err := utils.ReadInput(filename)
	solution := 0

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	for idx, _ := range input {
		history := []int{}
		nums := strings.Split(input[idx], " ")
		sequences := [][]int{}

		for _, num := range nums {
			intNum, _ := strconv.Atoi(num)
			history = append(history, intNum)
		}

		sequences = append(sequences, history)
		idx := 0

		for !d.allZero(sequences[len(sequences)-1]) {
			newSeq := []int{}

			for i := 0; i < len(sequences[idx])-1; i++ {
				newNum := sequences[idx][i+1] - sequences[idx][i]
				newSeq = append(newSeq, newNum)
			}

			sequences = append(sequences, newSeq)
			idx++
		}

		sequences[len(sequences)-1] = append([]int{0}, sequences[len(sequences)-1]...)

		for i := len(sequences) - 1; i >= 1; i-- {
			lastNum := sequences[i][0]
			previousNum := sequences[i-1][0]
			newNum := previousNum - lastNum
			sequences[i-1] = append([]int{newNum}, sequences[i-1]...)
		}

		solution += sequences[0][0]
	}

	return solution
}

func (d Day) allZero(seq []int) bool {
	for _, n := range seq {
		if n != 0 {
			return false
		}
	}
	return true
}
