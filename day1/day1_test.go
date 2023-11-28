package main

import (
	"adventofcode2023/utils"
	"log"
	"testing"
)

func TestSampleInput1(t *testing.T) {
	lines, err := utils.ReadInput()
	if err != nil {
		log.Fatal("Error reading input:", err)
	}
	t.Log(lines)
}

func TestActualInput1(t *testing.T) {
	t.Logf("Ok")
}

func TestSampleInput2(t *testing.T) {
	t.Logf("Ok")
}

func TestActualInput2(t *testing.T) {
	t.Logf("Ok")
}
