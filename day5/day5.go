package main

import (
	"adventofcode2023/utils"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

const (
	sampleInput = "test_input.txt"
	puzzleInput = "input.txt"
)

type Day struct{}

type MapContainer struct {
	name string
	maps []Map
}

func (mc *MapContainer) addMap(m Map) {
	mc.maps = append(mc.maps, m)
}

type Map struct {
	destStart int
	destEnd   int
	srcStart  int
	srcEnd    int
	mapRange  int
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
	solution := math.MaxInt32

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	seeds := []int{}
	allMaps := []*MapContainer{}
	var mapCont *MapContainer
	newMap := Map{}

	for _, line := range input {
		if line == "" {
			continue
		}

		splitLine := strings.Split(line, ":")

		if splitLine[0] == "seeds" {
			seeds = d.parseStrIntArray(strings.Trim(splitLine[1], " "))
			continue
		}

		if strings.Contains(splitLine[0], "map") {
			mapCont = &MapContainer{name: splitLine[0]}
			allMaps = append(allMaps, mapCont)
		} else {
			newMap = Map{}
			ranges := d.parseStrIntArray(line)
			newMap.mapRange = ranges[2]
			newMap.destStart = ranges[0]
			newMap.destEnd = ranges[0] + newMap.mapRange
			newMap.srcStart = ranges[1]
			newMap.srcEnd = ranges[1] + newMap.mapRange
			mapCont.addMap(newMap)
		}
	}

	for i := 0; i < len(seeds); i++ {
		seed := seeds[i]
		for i := 0; i < len(allMaps); i++ {
			mappingFound := false
			for _, mapping := range allMaps[i].maps {
				if seed >= mapping.srcStart && seed < mapping.srcEnd && !mappingFound {
					diff := seed - mapping.srcStart
					seed = mapping.destStart + diff
					mappingFound = true
				}
			}
			if allMaps[i].name == "humidity-to-location map" {
				if seed < solution {
					solution = seed
				}
			}
		}
	}

	return solution
}

func (d Day) Part2(filename string) int {
	input, err := utils.ReadInput(filename)
	solution := math.MaxInt32

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	seeds := []int{}
	allMaps := []*MapContainer{}
	var mapCont *MapContainer
	var newMap Map

	for _, line := range input {
		if line == "" {
			continue
		}

		splitLine := strings.Split(line, ":")

		if splitLine[0] == "seeds" {
			seeds = d.parseSeedRanges(strings.Trim(splitLine[1], " "))
			continue
		}

		if strings.Contains(splitLine[0], "map") {
			mapCont = &MapContainer{name: splitLine[0]}
			allMaps = append(allMaps, mapCont)
		} else {
			newMap = Map{}
			ranges := d.parseStrIntArray(line)
			newMap.mapRange = ranges[2]
			newMap.destStart = ranges[0]
			newMap.destEnd = ranges[0] + newMap.mapRange
			newMap.srcStart = ranges[1]
			newMap.srcEnd = ranges[1] + newMap.mapRange
			mapCont.addMap(newMap)
		}
	}

	for i := 0; i < len(seeds); i++ {
		seed := seeds[i]
		for i := 0; i < len(allMaps); i++ {
			mappingFound := false
			for _, mapping := range allMaps[i].maps {
				if seed >= mapping.srcStart && seed < mapping.srcEnd && !mappingFound {
					diff := seed - mapping.srcStart
					seed = mapping.destStart + diff
					mappingFound = true
				}
			}
			if allMaps[i].name == "humidity-to-location map" {
				if seed < solution {
					solution = seed
				}
			}
		}
	}

	return solution
}

func (d Day) parseStrIntArray(strInts string) []int {
	ints := []int{}
	nums := strings.Split(strInts, " ")
	for _, num := range nums {
		intNum, _ := strconv.Atoi(num)
		ints = append(ints, intNum)
	}
	return ints
}

func (d Day) parseSeedRanges(strInts string) []int {
	actualSeeds := []int{}

	ints := strings.Split(strInts, " ")

	for i := 0; i < len(ints); i += 2 {
		startNum, _ := strconv.Atoi(string(ints[i]))
		rangeNum, _ := strconv.Atoi(string(ints[i+1]))

		for i := startNum; i < startNum+rangeNum; i++ {
			actualSeeds = append(actualSeeds, i)
		}
	}

	return actualSeeds
}
