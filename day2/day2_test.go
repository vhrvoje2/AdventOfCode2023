package main

import (
	"testing"
)

func TestSampleInput1(t *testing.T) {
	expected := 8
	day1 := Day1{}
	actual := day1.Part1(sampleInput1)

	if actual != expected {
		t.Errorf("Returned %d, expected %d", actual, expected)
	}
}

func TestSampleInput2(t *testing.T) {
	expected := 2286

	day1 := Day1{}
	actual := day1.Part2(sampleInput2)

	if actual != expected {
		t.Errorf("Returned %d, expected %d", actual, expected)
	}
}
