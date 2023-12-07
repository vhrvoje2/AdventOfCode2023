package main

import (
	"adventofcode2023/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

const (
	sampleInput1 = "test_input1.txt"
	sampleInput2 = "test_input2.txt"
	puzzleInput  = "input.txt"
)

type Day struct{}

type HandType int8

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeKind
	FullHouse
	FourKind
	FiveKind
)

type Hand struct {
	cards    string
	bid      int
	handType HandType
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
	hands := []Hand{}

	cardMap := map[string]int{
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"J": 11,
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	for _, line := range input {
		parts := strings.Split(line, " ")
		bidInt, _ := strconv.Atoi(parts[1])
		newHand := Hand{cards: parts[0], bid: bidInt, handType: 0}
		newHand.handType = d.parseHand(newHand)
		hands = append(hands, newHand)
	}

	sort.Slice(hands, func(i, j int) bool {
		// 32T3K
		// KTJJT
		// KK677
		// T55J5
		// QQQJA
		h1 := hands[i]
		h2 := hands[j]

		if h1.handType == h2.handType {
			for i := 0; i < len(h1.cards); i++ {
				c1 := h1.cards[i]
				c2 := h2.cards[i]
				if c1 != c2 {
					return cardMap[string(c1)] < cardMap[string(c2)]
				} else {
					continue
				}
			}
		}

		return h1.handType < h2.handType
	})

	for i := 1; i <= len(hands); i++ {
		solution += i * hands[i-1].bid
	}

	return solution
}

func (d Day) Part2(filename string) int {
	input, err := utils.ReadInput(filename)
	solution := 0
	hands := []Hand{}

	cardMap := map[string]int{
		"J": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	for _, line := range input {
		parts := strings.Split(line, " ")
		bidInt, _ := strconv.Atoi(parts[1])
		newHand := Hand{cards: parts[0], bid: bidInt, handType: 0}
		newHand.handType = d.parseHand2(newHand)
		hands = append(hands, newHand)
	}

	sort.Slice(hands, func(i, j int) bool {
		// 32T3K
		// KK677
		// T55J5
		// QQQJA
		// KTJJT
		h1 := hands[i]
		h2 := hands[j]

		if h1.handType == h2.handType {
			for i := 0; i < len(h1.cards); i++ {
				c1 := h1.cards[i]
				c2 := h2.cards[i]
				if c1 != c2 {
					return cardMap[string(c1)] < cardMap[string(c2)]
				} else {
					continue
				}
			}
		}

		return h1.handType < h2.handType
	})

	for i := 1; i <= len(hands); i++ {
		solution += i * hands[i-1].bid
	}

	return solution
}

func (d Day) parseHand(h Hand) HandType {
	cards := strings.Split(h.cards, "")
	cardDict := map[string]int{}

	for _, c := range cards {
		cardDict[c]++
	}

	maxN := 1
	hasTwo := 0

	for _, v := range cardDict {
		if v > maxN {
			maxN = v
		}
		if v == 2 {
			hasTwo++
		}
	}

	switch maxN {
	case 5:
		return FiveKind
	case 4:
		return FourKind
	case 3:
		if hasTwo == 1 {
			return FullHouse
		}
		return ThreeKind
	case 2:
		if hasTwo == 2 {
			return TwoPair
		}
		return OnePair
	default:
		return HighCard
	}
}

func (d Day) parseHand2(h Hand) HandType {
	newHands := []Hand{}

	cardMap := map[string]int{
		"J": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"T": 10,
		"Q": 12,
		"K": 13,
		"A": 14,
	}

	valuesMap := []string{
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"T",
		"J",
		"Q",
		"K",
		"A",
	}

	for _, card := range valuesMap {
		hand := strings.Replace(h.cards, "J", card, -1)
		newHand := Hand{cards: hand, bid: h.bid, handType: 0}
		newHand.handType = d.parseHand(newHand)
		newHands = append(newHands, newHand)
	}

	sort.Slice(newHands, func(i, j int) bool {
		// 32T3K
		// KTJJT
		// KK677
		// T55J5
		// QQQJA
		h1 := newHands[i]
		h2 := newHands[j]

		if h1.handType == h2.handType {
			for i := 0; i < len(h1.cards); i++ {
				c1 := h1.cards[i]
				c2 := h2.cards[i]
				if c1 != c2 {
					return cardMap[string(c1)] < cardMap[string(c2)]
				} else {
					continue
				}
			}
		}

		return h1.handType < h2.handType
	})

	return newHands[len(newHands)-1].handType
}
