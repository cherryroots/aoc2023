package day1_test

import (
	"testing"

	day1 "aoc/days/day_01"
)

func TestDay1Parts(t *testing.T) {
	input := day1.Input
	expected := "54418"

	resultPart1 := day1.Calibrate(input)
	if resultPart1 != expected {
		t.Errorf("Day 1, Part 1: expected %s, got %s", expected, resultPart1)
	}
}
