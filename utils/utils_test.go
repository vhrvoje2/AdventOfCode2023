package utils_test

import (
	"adventofcode2023/utils"
	"testing"
)

func areEqualSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestReadFileLineByLine(t *testing.T) {
	expected := []string{"this", "is", "test", "input"}
	lines, err := utils.ReadInput()

	if err != nil {
		t.Errorf("Error reading input file.")
	}

	if !areEqualSlices(lines, expected) {
		t.Errorf("ReadInput = %v, want %v", lines, expected)
	}
}
