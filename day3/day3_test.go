package main

import (
	"testing"
)

func TestSampleInput1(t *testing.T) {
	expected := 4361
	day := Day{}
	actual := day.Part1(sampleInput1)

	if actual != expected {
		t.Errorf("Returned %d, expected %d", actual, expected)
	}
}

func TestSampleInput2(t *testing.T) {
	expected := 467835

	day := Day{}
	actual := day.Part2(sampleInput2)

	if actual != expected {
		t.Errorf("Returned %d, expected %d", actual, expected)
	}
}
