package dayxx

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type round struct {
	hand []int
	bid  int
}

type rangeOfRounds struct {
	name   string
	rounds []round
}

func (r round) isFlush() bool {
	sh := r.SortHand()

	return sh[0] == sh[4]
}

func (r round) isFourOfAKind() bool {
	sh := r.SortHand()

	// checking [ a, a, a, a, b]
	case1 := sh[0] == sh[1] &&
		sh[1] == sh[2] &&
		sh[2] == sh[3]
	// checking [ a, b, b, b, b]
	case2 := sh[1] == sh[2] &&
		sh[2] == sh[3] &&
		sh[3] == sh[4]
	return case1 || case2
}

func (r round) isFullHouse() bool {
	sh := r.SortHand()
	// checking [ a, a, a, b, b]
	case1 := sh[0] == sh[1] &&
		sh[1] == sh[2] &&
		sh[3] == sh[4]
	// checking [ a, a, b, b, b]
	case2 := sh[0] == sh[1] &&
		sh[2] == sh[3] &&
		sh[3] == sh[4]
	return case1 || case2
}

func (r round) isThreeOfAKind() bool {
	sh := r.SortHand()
	// checking [ x, x, x, a, b]
	case1 := sh[0] == sh[1] &&
		sh[1] == sh[2]
	// checking [a, x, x, x, b]
	case2 := sh[1] == sh[2] &&
		sh[2] == sh[3]
	// checking [a, b, x, x, x]
	case3 := sh[2] == sh[3] &&
		sh[3] == sh[4]
	return case1 || case2 || case3
}

func (r round) is2Pair() bool {
	sh := r.SortHand()
	// checking [ a, a, b, b, c]
	case1 := sh[0] == sh[1] &&
		sh[2] == sh[3]
	// checking [ a, a, b, c, c]
	case2 := sh[0] == sh[1] &&
		sh[3] == sh[4]
	// checking [ a, b, b, c, c]
	case3 := sh[1] == sh[2] &&
		sh[3] == sh[4]
	return case1 || case2 || case3
}

func (r round) isPair() bool {
	sh := r.SortHand()
	// checking [ a, a, x, y, z]
	case1 := sh[0] == sh[1]
	// checking [ x, b, b, y, z]
	case2 := sh[1] == sh[2]
	// checking [ x, y, c, c, z]
	case3 := sh[2] == sh[3]
	// checking [ x, y, z, d, d]
	case4 := sh[3] == sh[4]
	return case1 || case2 || case3 || case4
}

func (r round) SortHand() []int {
	var sortedHand []int
	sortedHand = append(sortedHand, r.hand...)

	sort.SliceStable(sortedHand, func(i, j int) bool {
		return sortedHand[i] < sortedHand[j]
	})

	return sortedHand
}

func sortRounds(input []round) []round {
	var sortedRounds []round
	sortedRounds = append(sortedRounds, input...)

	sort.SliceStable(sortedRounds, func(i, j int) bool {
		for k := 0; k < len(sortedRounds[i].hand); k++ {
			if sortedRounds[i].hand[k] != sortedRounds[j].hand[k] {
				return sortedRounds[i].hand[k] < sortedRounds[j].hand[k]
			}
		}
		return sortedRounds[i].bid < sortedRounds[j].bid
	})

	return sortedRounds
}

func categorizeHands(r round, types []rangeOfRounds) []rangeOfRounds {
	if r.isFlush() {
		types[6].name = "flush"
		types[6].rounds = append(types[6].rounds, r)
		return types
	}
	if r.isFourOfAKind() {
		types[5].name = "four of a kind"
		types[5].rounds = append(types[5].rounds, r)
		return types
	}
	if r.isFullHouse() {
		types[4].name = "full house"
		types[4].rounds = append(types[4].rounds, r)
		return types
	}
	if r.isThreeOfAKind() {
		types[3].name = "three of a kind"
		types[3].rounds = append(types[3].rounds, r)
		return types
	}
	if r.is2Pair() {
		types[2].name = "two pair"
		types[2].rounds = append(types[2].rounds, r)
		return types
	}
	if r.isPair() {
		types[1].name = "pair"
		types[1].rounds = append(types[1].rounds, r)
		return types
	}
	types[0].name = "high card"
	types[0].rounds = append(types[0].rounds, r)
	return types
}

func SolvePart1(input string) string {
	rounds := parseInput(input, false)
	output := ""
	result := 0
	types := make([]rangeOfRounds, 7)
	var allRounds []round

	for _, r := range rounds {
		types = categorizeHands(r, types)
	}

	for _, r := range types {
		r.rounds = sortRounds(r.rounds)
		allRounds = append(allRounds, r.rounds...)
	}

	for i, r := range allRounds {
		result += (i + 1) * r.bid
		output += fmt.Sprintf("<blue>%d</> * <red>%d</>, %v\n", i+1, r.bid, r.hand)
	}

	output += fmt.Sprintf("\n<blue>%d</>", result)

	return output
}

func SolvePart2(input string) string {
	rounds := parseInput(input, false)
	output := ""
	result := 0
	types := make([]rangeOfRounds, 7)
	var allRounds []round

	for _, r := range rounds {
		types = categorizeHands(r, types)
	}

	for _, r := range types {
		r.rounds = sortRounds(r.rounds)
		allRounds = append(allRounds, r.rounds...)
	}

	for i, r := range allRounds {
		result += (i + 1) * r.bid
		output += fmt.Sprintf("<blue>%d</> * <red>%d</>, %v\n", i+1, r.bid, r.hand)
	}

	output += fmt.Sprintf("\n<blue>%d</>", result)

	return output
}

func mapCard(card string, part2 bool) int {
	var cardMap = map[string]int{
		"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10, "9": 9,
		"8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2,
	}
	var cardMapPart2 = map[string]int{
		"A": 14, "K": 13, "Q": 12, "T": 10, "9": 9,
		"8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2, "J": 1,
	}

	if part2 {
		cardMap = cardMapPart2
	}

	if value, ok := cardMap[card]; ok {
		return value
	}
	return 0
}

func parseInput(input string, part2 bool) []round {
	parts := strings.Split(input, "\n")
	var output []round

	for _, part := range parts {
		fields := strings.Fields(part)
		strHand := strings.Split(fields[0], "")
		hand := make([]int, len(strHand))
		for i := range hand {
			hand[i] = mapCard(strHand[i], part2)
		}
		bid, _ := strconv.Atoi(fields[1])
		output = append(output, round{hand, bid})
	}

	return output
}
