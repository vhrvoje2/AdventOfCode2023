package main

import (
	"adventofcode2023/utils"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
)

const (
	sampleInput1 = "test_input1.txt"
	sampleInput2 = "test_input2.txt"
	puzzleInput  = "input.txt"
)

type Day struct{}

type Coord struct {
	x int
	y int
}

type Number struct {
	digits      string
	coords      []Coord
	symbolClose bool
}

type Symbol struct {
	symbol       string
	coords       Coord
	adjecentNums []int
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
	solution := 0
	numbers := []*Number{}
	symbols := []Symbol{}

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	for lineIdx, line := range input {
		for i := 0; i < len(line); i++ {
			readChar := string(line[i])
			_, err := strconv.Atoi(string(readChar))
			if readChar != "." && err == nil {
				newNumber := &Number{}
				newNumber.digits = readChar
				newNumber.coords = append(newNumber.coords, Coord{lineIdx, i})
				numbers = append(numbers, newNumber)

				for err == nil {
					i++
					if i > len(line)-1 {
						break
					}
					nextChar := string(line[i])
					_, err := strconv.Atoi(string(nextChar))
					if err == nil {
						newNumber.digits += nextChar
						newNumber.coords = append(newNumber.coords, Coord{lineIdx, i})
					} else {
						break
					}
				}
			}

			if i > len(line)-1 {
				break
			}

			readChar = string(line[i])
			_, err = strconv.Atoi(string(readChar))
			if readChar != "." && err != nil {
				newSymbol := Symbol{readChar, Coord{lineIdx, i}, []int{}}
				symbols = append(symbols, newSymbol)
			}
		}
	}

	for idx, num := range numbers {
		for _, cord := range num.coords {
			for _, sym := range symbols {
				x := float64(sym.coords.x - cord.x)
				y := float64(sym.coords.y - cord.y)
				dist := int(math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2)))
				if dist <= 1 {
					numbers[idx].symbolClose = true
				}
			}
		}
	}

	for _, num := range numbers {
		if num.symbolClose {
			val, _ := strconv.Atoi(num.digits)
			solution += val
		}
	}

	return solution
}

func (d Day) Part2(filename string) int {
	input, err := utils.ReadInput(filename)
	solution := 0
	numbers := []*Number{}
	symbols := []Symbol{}

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	for lineIdx, line := range input {
		for i := 0; i < len(line); i++ {
			readChar := string(line[i])
			_, err := strconv.Atoi(string(readChar))
			if readChar != "." && err == nil {
				newNumber := &Number{}
				newNumber.digits = readChar
				newNumber.coords = append(newNumber.coords, Coord{lineIdx, i})

				numbers = append(numbers, newNumber)
				for err == nil {
					i++
					if i > len(line)-1 {
						break
					}
					nextChar := string(line[i])
					_, err := strconv.Atoi(string(nextChar))
					if err == nil {
						newNumber.digits += nextChar
						newNumber.coords = append(newNumber.coords, Coord{lineIdx, i})
					} else {
						break
					}
				}
			}

			if i > len(line)-1 {
				break
			}

			readChar = string(line[i])
			_, err = strconv.Atoi(string(readChar))
			if readChar != "." && err != nil {
				newSymbol := Symbol{readChar, Coord{lineIdx, i}, []int{}}
				symbols = append(symbols, newSymbol)
			}
		}
	}

	for _, num := range numbers {
		for _, cord := range num.coords {
			for idx, sym := range symbols {
				x := float64(sym.coords.x - cord.x)
				y := float64(sym.coords.y - cord.y)
				dist := int(math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2)))
				if dist <= 1 && sym.symbol == "*" {
					intNum, _ := strconv.Atoi(num.digits)
					doesntHave := !slices.Contains(symbols[idx].adjecentNums, intNum)
					if doesntHave {
						symbols[idx].adjecentNums = append(symbols[idx].adjecentNums, intNum)
					}
				}
			}
		}
	}

	for _, sym := range symbols {
		if len(sym.adjecentNums) == 2 {
			solution += sym.adjecentNums[0] * sym.adjecentNums[1]
		}
	}

	return solution
}
