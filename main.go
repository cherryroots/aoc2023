package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"

	day1 "aoc/days/day_01"
	day2 "aoc/days/day_02"
	day3 "aoc/days/day_03"
	day4 "aoc/days/day_04"
	day5 "aoc/days/day_05"
	day6 "aoc/days/day_06"
	day7 "aoc/days/day_07"
	day8 "aoc/days/day_08"
)

type day struct {
	input        string
	exampleInput string
	part1        func(string) string
	part2        func(string) string
}

var days = map[int]day{
	1: {day1.Input, day1.ExampleInput, day1.SolvePart1, day1.SolvePart2},
	2: {day2.Input, day2.ExampleInput, day2.SolvePart1, day2.SolvePart2},
	3: {day3.Input, day3.ExampleInput, day3.SolvePart1, day3.SolvePart2},
	4: {day4.Input, day4.ExampleInput, day4.SolvePart1, day4.SolvePart2},
	5: {day5.Input, day5.ExampleInput, day5.SolvePart1, day5.SolvePart2},
	6: {day6.Input, day6.ExampleInput, day6.SolvePart1, day6.SolvePart2},
	7: {day7.Input, day7.ExampleInput, day7.SolvePart1, day7.SolvePart2},
	8: {day8.Input, day8.ExampleInput, day8.SolvePart1, day8.SolvePart2},
}

func (d day) Solve(input string) string {
	return fmt.Sprintf("Part 1:\n%s\n\nPart 2:\n%s\n\n", d.part1(input), d.part2(input))
}

func main() {
	var day int
	var example bool
	app := &cli.App{
		Name:  "Advent of Code",
		Usage: "Solve the puzzle for the given day",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "day",
				Usage:       "Which day to solve",
				Aliases:     []string{"d"},
				Value:       1,
				Destination: &day,
			},
			&cli.BoolFlag{
				Name:        "example",
				Usage:       "Use example input",
				Aliases:     []string{"e"},
				Value:       false,
				Destination: &example,
			},
		},
		Action: func(_ *cli.Context) error {
			runDay(day, example)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func runDay(day int, example bool) {
	start := time.Now()

	if day > len(days) {
		color.Red.Println("Invalid day")
		return
	}

	var result string
	if example {
		result = days[day].Solve(days[day].exampleInput)
	} else {
		result = days[day].Solve(days[day].input)
	}
	color.Print(result)
	elapsed := time.Since(start)
	fmt.Printf("Program took %s\n", elapsed)
}
