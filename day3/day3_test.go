package main

import (
	"testing"
)

func TestSampleInput1(t *testing.T) {
	expected := 4361
	day := Day{}
	actual := day.Part1(sampleInput)

	if actual != expected {
		t.Errorf("Returned %d, expected %d", actual, expected)
	}
}

func TestSampleInput2(t *testing.T) {
	expected := 467835

	day := Day{}
	actual := day.Part2(sampleInput)

	if actual != expected {
		t.Errorf("Returned %d, expected %d", actual, expected)
	}
}
