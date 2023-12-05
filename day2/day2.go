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

	gameIdx := 1
	gameValid := true
	availableCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, line := range input {
		spltLine := strings.Split(line, ":")
		subsets := strings.Split(spltLine[1], ";")

		for _, sub := range subsets {
			allCubes := strings.Split(sub, ",")
			for _, cubes := range allCubes {
				cubes = strings.Trim(cubes, " ")
				values := strings.Split(cubes, " ")
				numCubes, _ := strconv.Atoi(values[0])
				if numCubes > availableCubes[values[1]] {
					gameValid = false
				}
			}
		}

		if gameValid {
			solution += gameIdx
		}
		gameIdx++
		gameValid = true
	}

	return solution
}

func (d Day) Part2(filename string) int {
	input, err := utils.ReadInput(filename)
	solution := 0

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	for _, line := range input {

		spltLine := strings.Split(line, ":")
		subsets := strings.Split(spltLine[1], ";")
		gameCubes := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}

		for _, sub := range subsets {
			allCubes := strings.Split(sub, ",")
			for _, cubes := range allCubes {
				cubes = strings.Trim(cubes, " ")
				values := strings.Split(cubes, " ")
				cubeColor := values[1]
				numCubes, _ := strconv.Atoi(values[0])

				if numCubes > gameCubes[cubeColor] {
					gameCubes[cubeColor] = numCubes
				}
			}
		}

		power := gameCubes["red"] * gameCubes["blue"] * gameCubes["green"]
		solution += power
	}

	return solution
}
