package main

import (
	"testing"
)

func TestSampleInput1(t *testing.T) {
	expected := 142
	day1 := Day1{}
	actual := day1.Part1(sampleInput)

	if actual != expected {
		t.Errorf("Returned %d, expected %d", actual, expected)
	}
}

func TestSampleInput2(t *testing.T) {
	expected := 281

	day1 := Day1{}
	actual := day1.Part2(sampleInput)

	if actual != expected {
		t.Errorf("Returned %d, expected %d", actual, expected)
	}
}
