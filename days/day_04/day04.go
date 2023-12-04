package day04

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type card struct {
	id             int
	count          int
	myNumbers      []int
	winningNumbers []int
	matches        int
}

func (c card) Id() string {
	return strconv.Itoa(c.id)
}

func (c card) Matches() string {
	return strconv.Itoa(c.matches)
}

func (c card) Count() string {
	return strconv.Itoa(c.count)
}

func SolvePart1(input string) string {
	cards := parseInput(input)
	output := ""
	totalPoints := 0.0
	totalMatches := 0

	cards = checkMatches(cards)

	for _, card := range cards {
		if card.matches > 0 {
			points := 1 * math.Pow(2, float64(card.matches-1))
			totalPoints += points
			totalMatches++
		}
	}

	output += fmt.Sprintf("Total points: <blue>%f</>, Total matches: <blue>%d</>", totalPoints, totalMatches)

	return output
}

func SolvePart2(input string) string {
	cards := parseInput(input)
	output := ""
	totalCards := 0

	cards = checkMatches(cards)

	for i := 0; i < len(cards); i++ {
		card := cards[i]
		for j := 0; j < card.count; j++ {
			if card.matches > 0 {
				cards = addCards(cards, card)
			}
		}
	}

	highestCount, highestCard := 0, 0

	for _, card := range cards {
		totalCards += card.count
		if card.count > highestCount {
			highestCount = card.count
			highestCard = card.id
		}
	}

	for i, card := range cards {
		splitCount := 29
		mod := i % splitCount

		if mod == 0 {
			output += fmt.Sprintf("\n<bg=%s> </>", getColor(highestCount, card.count))
		} else {
			output += fmt.Sprintf("<bg=%s> </>", getColor(highestCount, card.count))
		}

	}

	output += fmt.Sprintf("\nTotal cards: <blue>%d</>, Highest card: <blue>%d</> with <blue>%d</> copies", totalCards, highestCard, highestCount)

	return output
}

func getColor(max, current int) string {
	minimum := 0
	normalized := normalize(float64(minimum), float64(max), float64(current))

	red, green, blue := 0.0, 0.0, 0.0

	if normalized < 0.0015 {
		normalized = normalize(0, 0.001, normalized)
		blue = math.Floor(normalized * 255)
	} else if normalized < 0.3 {
		normalized = normalize(0, 0.3, normalized)
		green = math.Floor(normalized * 255)
	} else {
		normalized = normalize(0, 1, normalized)
		red = math.Floor(normalized * 255)
	}

	result := fmt.Sprintf("%d,%d,%d", int(red), int(green), int(blue))

	return result

}

func normalize(min, max, value float64) float64 {
	return (float64(value) - float64(min)) / (float64(max) - float64(min))
}

func addCards(cards []card, card card) []card {
	matches := card.matches
	for i := 0; i < matches; i++ {
		cards[card.id+i].count++
	}
	return cards
}

func checkMatches(cards []card) []card {
	for i, card := range cards {
		for _, winningNumber := range card.winningNumbers {
			for _, myNumber := range card.myNumbers {
				if myNumber == winningNumber {
					cards[i].matches++
				}
			}
		}
	}

	return cards
}

func parseInput(input string) []card {
	lines := strings.Split(input, "\n")
	var cards []card

	for _, line := range lines {
		cardParts := strings.Split(line, ":")
		cardStringID := strings.Split(cardParts[0], " ")
		cardID, _ := strconv.Atoi(cardStringID[len(cardStringID)-1])

		cardParts = strings.Split(cardParts[1], "|")
		myNumbers := parseNumbers(cardParts[0])
		winningNumbers := parseNumbers(cardParts[1])

		cards = append(cards, card{
			id:             cardID,
			count:          1,
			myNumbers:      myNumbers,
			winningNumbers: winningNumbers,
			matches:        0,
		})

	}
	return cards
}

func parseNumbers(input string) []int {
	var numbers []int

	for _, number := range strings.Split(input, " ") {
		if number == "" {
			continue
		}
		numberInt, _ := strconv.Atoi(number)
		numbers = append(numbers, numberInt)
	}

	return numbers
}
