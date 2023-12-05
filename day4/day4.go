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

func (d Day) Part2(filename string) int {
	input, err := utils.ReadInput(filename)
	solution := 0

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	cards := make([]int, len(input))
	for idx := range cards {
		cards[idx] = 1
	}

	for idx, line := range input {
		card := strings.Split(line, ":")
		nums := strings.Split(card[1], "|")
		winningNums := strings.Split(nums[0], " ")
		myNums := strings.Split(nums[1], " ")
		count := 0

		for _, winningNum := range winningNums {
			for _, myNum := range myNums {
				wNum := strings.Trim(winningNum, " ")
				mNum := strings.Trim(myNum, " ")
				if wNum == mNum && wNum != "" && mNum != "" {
					count += 1
				}
			}
		}

		for i := 1; i <= count; i++ {
			cards[idx+i] += cards[idx]
		}
	}

	for _, n := range cards {
		solution += n
	}

	return solution
}
