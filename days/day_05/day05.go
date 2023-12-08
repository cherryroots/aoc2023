package day05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type almanacMap struct {
	destStart   int
	sourceStart int
	rangeLength int
}

func (m almanacMap) ConvertMap(input int, reverse bool) (bool, int) {
	var value, start, end int

	if reverse {
		value = m.sourceStart + input - m.destStart
		start = m.destStart
		end = m.destStart + m.rangeLength
	} else {
		value = m.destStart + input - m.sourceStart
		start = m.sourceStart
		end = m.sourceStart + m.rangeLength
	}

	if input >= start && input < end {
		return true, value
	} else {
		return false, value
	}
}

type almanacMaps struct {
	name string
	maps []almanacMap
}

func (m almanacMaps) ConvertMaps(input int, reverse bool) int {
	var lowest int = 0
	var validConvertion bool = false

	for _, mapItem := range m.maps {
		valid, value := mapItem.ConvertMap(input, reverse)
		if valid {
			if lowest == 0 {
				validConvertion = true
				lowest = value
			} else if value < lowest {
				validConvertion = true
				lowest = value
			}
		} else {
			continue
		}
	}

	if !validConvertion {
		return input
	} else {
		return lowest
	}
}

type seeds struct {
	seeds []int
}

type seedRange struct {
	lowest  int
	highest int
}

func (s seeds) ConvertToRange() []seedRange {
	var base, rangeLength int
	var seedRanges []seedRange

	for i, seed := range s.seeds {
		if i%2 == 0 {
			base = seed
		} else {
			rangeLength = seed - 1
			seedRanges = append(seedRanges, seedRange{base, base + rangeLength})
		}
	}

	return seedRanges
}

type almanac struct {
	seeds
	maps []almanacMaps
}

func (a almanac) RunSeed(seed int) int {
	for _, mapItem := range a.maps {
		seed = mapItem.ConvertMaps(seed, false)
	}

	return seed
}

// slow brute force method iterating over all the seeds, enough for part 1
func (a almanac) Run() []int {
	var output []int

	for _, seed := range a.seeds.seeds {
		output = append(output, a.RunSeed(seed))
	}

	return output
}

// reverse searching, going over each location until one matches a seed range, for part 2
func (a almanac) RunReverse() int {
	maxLocation := int(^uint(0) >> 1)
	seedRanges := a.seeds.ConvertToRange()
	var reverseMaps []almanacMaps
	reverseMaps = append(reverseMaps, a.maps...)
	slices.Reverse(reverseMaps)

	for i := 0; i < maxLocation; i++ {
		var potentialSeed int = i
		for _, mapItem := range reverseMaps {
			potentialSeed = mapItem.ConvertMaps(potentialSeed, true)
		}
		for _, sr := range seedRanges {
			if potentialSeed >= sr.lowest && potentialSeed <= sr.highest {
				return potentialSeed
			}
		}
	}

	return 0
}

func SolvePart1(input string) string {
	almanac := parseInput(input)
	var lowestSeed int
	var seeds []int
	output := ""

	for _, seed := range almanac.Run() {
		if seed < lowestSeed || lowestSeed == 0 {
			lowestSeed = seed
		}
		seeds = append(seeds, seed)
	}

	for _, seed := range seeds {
		if seed != lowestSeed {
			output += fmt.Sprintf("<blue>%d</> ", seed)
		} else {
			output += fmt.Sprintf("<red>%d</> ", seed)
		}
	}

	return output
}

func SolvePart2(input string) string {
	almanac := parseInput(input)
	lowestSeed := almanac.RunReverse()
	seedOutput := almanac.RunSeed(lowestSeed)
	output := fmt.Sprintf("Seed: <blue>%d</>, lowest location: <red>%d</> ", lowestSeed, seedOutput)

	return output
}

func parseInput(input string) almanac {
	parts := strings.Split(input, "\n\n")
	var almanac almanac

	almanac.seeds = parseSeeds(parts[0])
	for _, part := range parts[1:] {
		almanac.maps = append(almanac.maps, parseMaps(part))
	}

	return almanac
}

func parseSeeds(input string) seeds {
	var seeds seeds
	seedStrings := strings.Split(input, " ")

	for _, seedString := range seedStrings[1:] {
		seed, _ := strconv.Atoi(seedString)
		seeds.seeds = append(seeds.seeds, seed)
	}

	return seeds
}

func parseMaps(input string) almanacMaps {
	var almanacMaps almanacMaps
	almanacMaps.name = input[:strings.Index(input, " ")]
	lines := strings.Split(input, "\n")

	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		var mapItem almanacMap
		split := strings.Split(line, " ")
		mapItem.destStart, _ = strconv.Atoi(split[0])
		mapItem.sourceStart, _ = strconv.Atoi(split[1])
		mapItem.rangeLength, _ = strconv.Atoi(split[2])
		almanacMaps.maps = append(almanacMaps.maps, mapItem)
	}

	return almanacMaps
}
