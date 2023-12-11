package day06

import (
	"fmt"
	"strconv"
	"strings"
)

type race struct {
	time     int
	distance int
}

func (r race) CheckWins() int {
	wins := 0
	for i := 0; i <= r.time; i++ {
		traveledDistance := i * (r.time - i)
		if traveledDistance > r.distance {
			wins++
		}
	}

	return wins
}

func SolvePart1(input string) string {
	output := ""
	var marginOfError int = 1
	races := parseInputPart1(input)

	for _, race := range races {
		marginOfError *= race.CheckWins()
	}

	output += fmt.Sprintf("Margin of error: <blue>%d</>", marginOfError)

	return output
}

func SolvePart2(input string) string {
	race := parseInputPart2(input)
	marginOfError := race.CheckWins()
	output := fmt.Sprintf("Margin of error: <blue>%d</>", marginOfError)

	return output
}

func parseInputPart1(input string) []race {
	parts := strings.Split(input, "\n")
	var output []race

	times := strings.Fields(parts[0])
	distances := strings.Fields(parts[1])

	for i := 1; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		output = append(output, race{time, distance})
	}

	return output
}

func parseInputPart2(input string) race {
	parts := strings.Split(input, "\n")
	times := strings.Fields(parts[0])
	distances := strings.Fields(parts[1])

	time, _ := strconv.Atoi(strings.Join(times[1:], ""))
	distance, _ := strconv.Atoi(strings.Join(distances[1:], ""))

	output := race{time, distance}

	return output
}
