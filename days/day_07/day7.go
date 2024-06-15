package dayxx

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type hands struct {
	hand         []int
	originalHand []int
	bid          int
}

type rangeOfRounds struct {
	name   string
	rounds []hands
}

func (r hands) isFlush() bool {
	sh := r.SortHand()

	return sh[0] == sh[4]
}

func (r hands) isFourOfAKind() bool {
	sh := r.SortHand()

	// checking [ a, a, a, a, b]
	case1 := sh[0] == sh[1] && sh[1] == sh[2] && sh[2] == sh[3]
	// checking [ a, b, b, b, b]
	case2 := sh[1] == sh[2] && sh[2] == sh[3] && sh[3] == sh[4]
	return case1 || case2
}

func (r hands) isFullHouse() bool {
	sh := r.SortHand()
	// checking [ a, a, a, b, b]
	case1 := sh[0] == sh[1] && sh[1] == sh[2] && sh[3] == sh[4]
	// checking [ a, a, b, b, b]
	case2 := sh[0] == sh[1] && sh[2] == sh[3] && sh[3] == sh[4]
	return case1 || case2
}

func (r hands) isThreeOfAKind() bool {
	sh := r.SortHand()
	// checking [ x, x, x, a, b]
	case1 := sh[0] == sh[1] && sh[1] == sh[2]
	// checking [a, x, x, x, b]
	case2 := sh[1] == sh[2] && sh[2] == sh[3]
	// checking [a, b, x, x, x]
	case3 := sh[2] == sh[3] && sh[3] == sh[4]
	return case1 || case2 || case3
}

func (r hands) is2Pair() bool {
	sh := r.SortHand()
	// checking [ a, a, b, b, c]
	case1 := sh[0] == sh[1] && sh[2] == sh[3]
	// checking [ a, a, b, c, c]
	case2 := sh[0] == sh[1] && sh[3] == sh[4]
	// checking [ a, b, b, c, c]
	case3 := sh[1] == sh[2] && sh[3] == sh[4]
	return case1 || case2 || case3
}

func (r hands) isPair() bool {
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

func (r hands) SortHand() []int {
	var sortedHand []int
	sortedHand = append(sortedHand, r.hand...)

	sort.SliceStable(sortedHand, func(i, j int) bool {
		return sortedHand[i] < sortedHand[j]
	})

	return sortedHand
}

func sortRounds(input []hands) []hands {
	var sortedRounds []hands
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

func categorizeHands(r hands, types []rangeOfRounds) []rangeOfRounds {
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

func generateAllPossibleCards(r hands) []hands {
	var allPossibleRounds []hands

	var wildcardPositions []int
	for i, card := range r.hand {
		if card == 1 {
			wildcardPositions = append(wildcardPositions, i)
		}
	}

	if len(wildcardPositions) == 0 {
		allPossibleRounds = append(allPossibleRounds, r)
	} else {
		generateWildcardCombinations(r, &allPossibleRounds, wildcardPositions, 0)
	}

	return allPossibleRounds
}

func generateWildcardCombinations(r hands, allPossibleRounds *[]hands, wildcardPositions []int, index int) {
	if index == len(wildcardPositions) {
		// A complete combination is generated, add it to the results
		newRound := hands{hand: make([]int, len(r.hand)), originalHand: make([]int, len(r.hand)), bid: r.bid}
		copy(newRound.hand, r.hand)
		copy(newRound.originalHand, r.originalHand)
		*allPossibleRounds = append(*allPossibleRounds, newRound)
		return
	}

	pos := wildcardPositions[index]
	for cardValue := 1; cardValue <= 14; cardValue++ {
		r.originalHand[pos] = cardValue
		generateWildcardCombinations(r, allPossibleRounds, wildcardPositions, index+1)
	}
}

func SolvePart1(input string) string {
	rounds := parseInput(input, false)
	output := ""
	result := 0
	types := make([]rangeOfRounds, 7)
	var allRounds []hands

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
	rounds := parseInput(input, true)
	output := ""
	result := 0
	types := make([]rangeOfRounds, 7)
	var allRounds []hands

	for _, r := range rounds {
		possibleHands := make([]rangeOfRounds, 7)
		bestHands := rangeOfRounds{}
		uniqueRounds := generateAllPossibleCards(r)
		for _, r := range uniqueRounds {
			possibleHands = categorizeHands(r, possibleHands)
		}
		for i := len(possibleHands) - 1; i >= 0; i-- {
			if len(possibleHands[i].rounds) > 0 {
				bestHands = possibleHands[i]
				break
			}
		}
		bestHands.rounds = sortRounds(bestHands.rounds)
		types = categorizeHands(bestHands.rounds[0], types)
	}

	for _, r := range types {
		r.rounds = sortRounds(r.rounds)
		allRounds = append(allRounds, r.rounds...)
	}

	for i, r := range allRounds {
		result += (i + 1) * r.bid
		output += fmt.Sprintf("<blue>%d</> * <red>%d</>, %v, %v\n", i+1, r.bid, r.hand, r.originalHand)
	}

	output += fmt.Sprintf("\n<blue>%d</>", result)

	return output
}

func mapCard(card string, part2 bool) int {
	cardMap := map[string]int{
		"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10, "9": 9,
		"8": 8, "7": 7, "6": 6, "5": 5, "4": 4, "3": 3, "2": 2,
	}
	cardMapPart2 := map[string]int{
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

func parseInput(input string, part2 bool) []hands {
	parts := strings.Split(input, "\n")
	var output []hands

	for _, part := range parts {
		fields := strings.Fields(part)
		strHand := strings.Split(fields[0], "")
		hand := make([]int, len(strHand))
		originalHand := make([]int, len(strHand))
		for i := range hand {
			hand[i] = mapCard(strHand[i], part2)
		}
		bid, _ := strconv.Atoi(fields[1])
		copy(originalHand, hand)
		output = append(output, hands{hand, originalHand, bid})
	}

	return output
}
