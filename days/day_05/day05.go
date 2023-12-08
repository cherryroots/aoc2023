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

type conversionResult struct {
	valid bool
	value int
}

func (m almanacMap) String() string {
	return fmt.Sprintf("%d %d %d", m.destStart, m.sourceStart, m.rangeLength)
}

func (m almanacMap) ConvertMap(input int, reverse bool) conversionResult {
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
		return conversionResult{
			valid: true,
			value: value,
		}
	} else {
		return conversionResult{
			valid: false,
			value: input,
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

func (m almanacMaps) ConvertMaps(input int, reverse bool) int {
	var lowest int = 0
	var validConvertion bool = false
	for _, mapItem := range m.maps {
		var convertion conversionResult
		if reverse {
			convertion = mapItem.ConvertMap(input, true)
		} else {
			convertion = mapItem.ConvertMap(input, false)
		}
		if convertion.valid {
			if lowest == 0 {
				validConvertion = true
				lowest = convertion.value
			} else if convertion.value < lowest {
				validConvertion = true
				lowest = convertion.value
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
	seed []int
}

type seedRange struct {
	lowest  int
	highest int
}

func (s seeds) ConvertToRange(split int) []seedRange {
	var base, rangeLength int
	var seedRanges []seedRange
	for i, seed := range s.seed {
		if i%2 == 0 {
			base = seed
		} else {
			rangeLength = seed
			increment := rangeLength / split
			for j := 0; j < split; j++ {
				newLowest := base + (j * increment)
				newHighest := newLowest + increment - 1

				if j == split-1 {
					newHighest = base + rangeLength - 1
				}

				seedRanges = append(seedRanges, seedRange{
					lowest:  newLowest,
					highest: newHighest,
				})
			}
		}
	}

	return seedRanges

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

func (a almanac) RunSeed(seed int) int {
	for _, mapItem := range a.maps {
		seed = mapItem.ConvertMaps(seed, false)
	}

	return seed
}

// slow brute force method iterating over all the seeds
func (a almanac) Run() []int {
	var output []int
	results := make(chan int, len(a.seeds.seed))

	for _, seed := range a.seeds.seed {
		go func(s int) {
			for _, mapItem := range a.maps {
				s = mapItem.ConvertMaps(s, false)
			}
			results <- s
		}(seed)
	}

	for range a.seeds.seed {
		output = append(output, <-results)
	}

	close(results)

	return output
}

// reverse searching, going over each location until one matches a seed range
func (a almanac) RunReverse() int {
	maxLocation := int(^uint(0) >> 1)
	seedRanges := a.seeds.ConvertToRange(1)
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

// slow brute force method iterating over all the seeds in the ranges
func (a almanac) RunPaired() []int {
	var output []int
	ranges := a.seeds.ConvertToRange(5)
	channel := make(chan int, len(ranges))
	fmt.Printf("Ranges: %d\n", len(ranges))

	for _, sr := range ranges {
		go func(sr seedRange) {
			lowChan := make(chan int, sr.highest-sr.lowest+1)
			for i := sr.lowest; i <= sr.highest; i++ {
				seed := i
				go func(seed int) {
					for _, mapItem := range a.maps {
						seed = mapItem.ConvertMaps(seed, false)
					}
					lowChan <- seed

				}(seed)
			}
			lowest := 0
			for i := 0; i < sr.highest-sr.lowest+1; i++ {
				if i == 0 {
					lowest = <-lowChan
				} else {
					seed := <-lowChan
					if seed < lowest {
						lowest = seed
					}
				}
			}
			channel <- lowest
			fmt.Printf("Finished range: %d\n", sr.highest-sr.lowest+1)
		}(sr)
	}
	for range ranges {
		output = append(output, <-channel)
	}

	close(channel)

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
	output := ""

	output += "\n\n"

	lowestSeed = almanac.RunReverse()
	seedOutput := almanac.RunSeed(lowestSeed)

	output += fmt.Sprintf("Lowest seed: <blue>%d</>, location: <red>%d</> ", lowestSeed, seedOutput)

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
