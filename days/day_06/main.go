package day06

import (
	"strings"
)

func SolvePart1(input string) string {
	return ""
}

func SolvePart2(input string) string {
	return ""
}

func parseInput(input string) string {
	parts := strings.Split(input, "\n")
	output := ""

	for _, part := range parts[1:] {
		output += part
	}

	return output
}
