package day05

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type conversionResult struct {
	converted bool
	value     int
}

type almanacMap struct {
	destStart   int
	sourceStart int
	rangeLength int
}

func (m almanacMap) String() string {
	return fmt.Sprintf("%d %d %d", m.destStart, m.sourceStart, m.rangeLength)
}

func (m almanacMap) ConvertMap(input int) conversionResult {
	if input >= m.sourceStart && input < m.sourceStart+m.rangeLength {
		return conversionResult{
			converted: true,
			value:     m.destStart + input - m.sourceStart,
		}
	} else {
		return conversionResult{
			converted: false,
			value:     input,
		}
	}
}

type almanacMaps struct {
	name string
	maps []almanacMap
}

func (m almanacMaps) String() string {
	var output []string
	output = append(output, m.name)
	for _, mapItem := range m.maps {
		output = append(output, mapItem.String())
	}
	return strings.Join(output, "\n")
}

func (m almanacMaps) ConvertMaps(input int) int {
	var convertions []conversionResult
	var validConvertions []conversionResult
	for _, mapItem := range m.maps {
		convertions = append(convertions, mapItem.ConvertMap(input))
	}
	for _, destination := range convertions {
		if destination.converted {
			validConvertions = append(validConvertions, destination)
		}
	}
	if len(validConvertions) == 0 {
		return input
	} else {
		sort.Slice(validConvertions, func(i, j int) bool {
			return validConvertions[i].value < validConvertions[j].value
		})
		return validConvertions[0].value
	}
}

type seeds struct {
	seed []int
}

func (s seeds) ConvertToRange() []int {
	var output []int
	var base, rangeLength int
	for i, seed := range s.seed {
		if i%2 == 0 {
			base = seed
		} else {
			rangeLength = seed
			for j := 0; j < rangeLength; j++ {
				output = append(output, base+j)
			}
		}
	}
	return output
}

func (s seeds) String() string {
	var output []string
	for _, seed := range s.seed {
		output = append(output, strconv.Itoa(seed))
	}
	return strings.Join(output, " ")
}

type almanac struct {
	seeds seeds
	maps  []almanacMaps
}

func (a almanac) String() string {
	var output []string
	output = append(output, a.seeds.String())
	for _, mapItem := range a.maps {
		output = append(output, mapItem.String())
	}
	return strings.Join(output, "\n\n")
}

func (a almanac) Run() []int {
	var output []int
	for _, seed := range a.seeds.seed {
		for _, mapItem := range a.maps {
			seed = mapItem.ConvertMaps(seed)
		}
		output = append(output, seed)
	}
	return output
}

func (a almanac) RunPaired() []int {
	var output []int
	for _, seed := range a.seeds.ConvertToRange() {
		for _, mapItem := range a.maps {
			seed = mapItem.ConvertMaps(seed)
		}
		output = append(output, seed)
	}
	return output
}

func SolvePart1(input string) string {
	almanac := parseInput(input)
	var lowestSeed int
	var seeds []int
	output := ""

	output += almanac.String()

	output += "\n\n"

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
	var lowestSeed int
	var seeds []int
	output := ""

	output += "\n\n"

	for _, seed := range almanac.RunPaired() {
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
		seeds.seed = append(seeds.seed, seed)
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
