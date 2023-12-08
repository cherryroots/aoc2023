package day05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
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

	if start <= input && input < end {
		return true, value
	} else {
		return false, value
	}
}

type almanacMaps struct {
	maps []almanacMap
}

func (m almanacMaps) ConvertMaps(input int, reverse bool) int {
	for _, mapItem := range m.maps {
		valid, value := mapItem.ConvertMap(input, reverse)
		if valid {
			return value
		}
	}

	return input
}

type seeds struct {
	seeds []int
}

type seedRange struct {
	lowest  int
	highest int
}

func (s seeds) ConvertToRanges() []seedRange {
	var base, rangeLength int
	var seedRanges []seedRange

	for i, seed := range s.seeds {
		if i%2 == 0 {
			base = seed
		} else {
			rangeLength = seed
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

func (a almanac) RunIntervals() int {
	seeds := a.seeds.ConvertToRanges()
	for _, mapItem := range a.maps {
		var newSeedRanges []seedRange
		for len(seeds) > 0 {
			start, end := seeds[0].lowest, seeds[0].highest
			seeds = seeds[1:]
			found := false
			for _, eachMap := range mapItem.maps {
				dst, src, length := eachMap.destStart, eachMap.sourceStart, eachMap.rangeLength
				overlapStart := max(start, src)
				overlapEnd := min(end, src+length)
				if overlapStart < overlapEnd {
					found = true
					newSeedRanges = append(newSeedRanges, seedRange{overlapStart - src + dst, overlapEnd - src + dst})
					if overlapStart > start {
						seeds = append(seeds, seedRange{start, overlapStart})
					}
					if end > overlapEnd {
						seeds = append(seeds, seedRange{overlapEnd, end})
					}
					break
				}
			}
			if !found {
				newSeedRanges = append(newSeedRanges, seedRange{start, end})
			}
		}
		seeds = newSeedRanges
	}
	var lowest int
	for _, seed := range seeds {
		if seed.lowest < lowest || lowest == 0 {
			lowest = seed.lowest
		}
	}

	return lowest
}

// reverse searching. Takes about 0.92 seconds to run
func (a almanac) RunReverse() int {
	maxLocation := int(^uint(0) >> 1)
	seedRanges := a.seeds.ConvertToRanges()
	reverseMaps := append([]almanacMaps{}, a.maps...)
	slices.Reverse(reverseMaps)

	for i := 0; i < maxLocation; i++ {
		var potentialSeed int = i
		for _, mapItem := range reverseMaps {
			potentialSeed = mapItem.ConvertMaps(potentialSeed, true)
		}
		for _, sr := range seedRanges {
			if sr.lowest <= potentialSeed && potentialSeed < sr.highest {
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
	startInterval := time.Now()
	lowestLocation := almanac.RunIntervals()
	elapsedInterval := time.Since(startInterval)
	startReverse := time.Now()
	_ = almanac.RunReverse()
	elapsedReverse := time.Since(startReverse)
	seedOutput := almanac.maps[0].ConvertMaps(lowestLocation, true)
	output := fmt.Sprintf("Seed: <blue>%d</>, lowest location: <red>%d</>\nElapsed interval: %s\nElapsed reverse: %s", seedOutput, lowestLocation, elapsedInterval, elapsedReverse)

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
