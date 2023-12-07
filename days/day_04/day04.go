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

func SolvePart1(input string) string {
	cards := parseInput(input)
	output := ""
	totalPoints, totalMatches := 0, 0

	cards = checkMatches(cards)

	for _, card := range cards {
		if card.matches > 0 {
			points := 1 * math.Pow(2, float64(card.matches-1))
			totalPoints += int(points)
			totalMatches++
		}
	}

	output += fmt.Sprintf("Total points: <blue>%d</>, Total matches: <blue>%d</>", totalPoints, totalMatches)

	return output
}

func SolvePart2(input string) string {
	cards := parseInput(input)
	output := ""
	totalCards, highestCount, highestCard := 0, 0, 0

	cards = checkMatches(cards)

	for i := 0; i < len(cards); i++ {
		card := cards[i]
		if card.matches > 0 && card.count > 0 {
			for j := 0; j < card.count; j++ {
				cards = addCards(cards, card)
			}
		}
	}

	for _, card := range cards {
		totalCards += card.count
		if card.count > highestCount {
			highestCount = card.count
			highestCard = card.id
		}
	}

	output += fmt.Sprintf("\nTotal cards: <blue>%d</>, Highest card: <blue>%d</> with <blue>%d</> copies", totalCards, highestCard, highestCount)

	return output
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
		if line == "" {
			continue
		}
		cards = append(cards, parseCard(line))
	}
	return cards
}

func parseCard(input string) card {
	cardParts := strings.Split(input, ":")
	cardStringID := strings.Split(cardParts[0], " ")
	cardID, _ := strconv.Atoi(cardStringID[len(cardStringID)-1])

	cardParts = strings.Split(cardParts[1], "|")
	myNumbers := parseNumbers(cardParts[0])
	winningNumbers := parseNumbers(cardParts[1])

	return card{
		id:             cardID,
		count:          1,
		myNumbers:      myNumbers,
		winningNumbers: winningNumbers,
		matches:        0,
	}
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
