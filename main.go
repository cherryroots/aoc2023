package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gookit/color"

	day1 "aoc/days/day_01"
	day2 "aoc/days/day_02"
	day3 "aoc/days/day_03"
	day4 "aoc/days/day_04"
	day5 "aoc/days/day_05"
)

func main() {
	start := time.Now()

	result, err := RunSolution(5)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	color.Print(result)
	elapsed := time.Since(start)
	fmt.Printf("Program took %s\n", elapsed)
}

func RunSolution(day int) (string, error) {
	switch day {
	case 1:
		// Implement the solution for Day 1 here.
		solution := day1.Calibrate(day1.Input)
		return solution, nil
	case 2:
		// Implement the solution for Day 2 here.
		solutionPart1 := day2.SolvePart1(day2.Input)
		solutionPart2 := day2.SolvePart2(day2.Input)
		solution := fmt.Sprintf("Part 1:\n\n%s\n\nPart 2:\n\n%s\n", solutionPart1, solutionPart2)
		return solution, nil
	case 3:
		// Implement the solution for Day 3 here.
		solutionPart1 := day3.SolvePart1(day3.Input)
		solutionPart2 := day3.SolvePart2(day3.Input)
		solution := fmt.Sprintf("Part 1:\n\n%s\n\nPart 2:\n\n%s\n", solutionPart1, solutionPart2)
		return solution, nil
	case 4:
		// Implement the solution for Day 4 here.
		solutionPart1 := day4.SolvePart1(day4.Input)
		solutionPart2 := day4.SolvePart2(day4.Input)
		solution := fmt.Sprintf("Part 1:\n\n%s\n\nPart 2:\n%s\n", solutionPart1, solutionPart2)
		return solution, nil
	case 5:
		// Implement the solution for Day 5 here.
		solutionPart1 := day5.SolvePart1(day5.Input)
		solutionPart2 := day5.SolvePart2(day5.ExampleInput)
		solution := fmt.Sprintf("Part 1:\n\n%s\n\nPart 2:\n%s\n", solutionPart1, solutionPart2)
		return solution, nil
	// Add cases for other days here...
	default:
		return "", fmt.Errorf("solution for Day %d not found", day)
	}
}
