package main

import (
	"adventofcode2023/utils"
	"fmt"
	"log"
)

func main() {
	part1()
	part2()
}

func part1() {
	solution, err := utils.ReadInput()

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	fmt.Println("Part 1 solution")
	fmt.Println(solution)
}

func part2() {
	solution, err := utils.ReadInput()

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	fmt.Println("Part 2 solution")
	fmt.Println(solution)
}
