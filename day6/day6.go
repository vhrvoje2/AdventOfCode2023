package main

import (
	"adventofcode2023/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	sampleInput = "test_input.txt"
	puzzleInput = "input.txt"
)

type Day struct{}

type Race struct {
	time int
	dist int
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
	solution := 1

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	races := []*Race{}

	times := strings.Fields(strings.TrimSpace(strings.Split(input[0], ":")[1]))
	dists := strings.Fields(strings.TrimSpace(strings.Split(input[1], ":")[1]))

	for i := 0; i < len(times); i++ {
		timeInt, _ := strconv.Atoi(string(times[i]))
		race := Race{time: timeInt}
		races = append(races, &race)
	}

	for i := 0; i < len(dists); i++ {
		distInt, _ := strconv.Atoi(string(dists[i]))
		races[i].dist = distInt
	}

	for _, race := range races {
		wins := d.parseRace(*race)
		solution *= wins
	}

	return solution
}

func (d Day) Part2(filename string) int {
	input, err := utils.ReadInput(filename)

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	time, _ := strconv.Atoi(strings.Replace(strings.Split(input[0], ":")[1], " ", "", -1))
	dist, _ := strconv.Atoi(strings.Replace(strings.Split(input[1], ":")[1], " ", "", -1))

	race := Race{time: time, dist: dist}

	return d.parseRace(race)
}

func (d Day) parseRace(r Race) int {
	wins := 0

	for i := 1; i < r.time; i++ {
		hold := i
		travelTime := r.time - hold
		distTraveled := hold * travelTime
		if distTraveled > r.dist {
			wins++
		}
	}

	return wins
}
