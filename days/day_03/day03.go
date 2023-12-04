package day03

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type symbolinfo struct {
	sign   string
	row    int
	column int
}
type partinfo struct {
	number string
	valid  bool
	symbol symbolinfo
}

func (p partinfo) NumberInt() int {
	number, _ := strconv.Atoi(p.number)
	return number
}

func SolvePart1(s string) string {
	array := initArray(s)

	output := ""
	result := 0

	for row := 0; row < len(array[0]); row++ {
		inDots := true
		part := partinfo{
			number: "",
			valid:  false,
			symbol: symbolinfo{},
		}
		for column := 0; column < len(array); column++ {
			cell := array[row][column]
			char := []rune(cell)[0]
			if unicode.IsDigit(char) {
				inDots = false
				if !part.valid {
					part.valid, part.symbol = checkAround(array, row, column)
				}
				part.number += cell
			} else {
				if !inDots {
					if part.valid {
						output += fmt.Sprintf("<blue>%s</>", part.number)
						result += part.NumberInt()
					} else {
						output += fmt.Sprintf("<red>%s</>", part.number)
					}
					part = partinfo{}
				}
				if cell != "." {
					output += fmt.Sprintf("<green>%s</>", cell)
					continue
				} else {
					output += cell
				}
				inDots = true
			}

			if column == len(array)-1 {
				if part.valid {
					output += fmt.Sprintf("<blue>%s</>", part.number)
					result += part.NumberInt()
				} else {
					output += fmt.Sprintf("<red>%s</>", part.number)
				}
			}
		}
		output += "\n"
	}
	output += strconv.Itoa(result)

	return output
}

func SolvePart2(s string) string {
	array := initArray(s)

	output := ""
	result := 0
	var parts []partinfo

	for row := 0; row < len(array[0]); row++ {
		inDots := true
		part := partinfo{
			number: "",
			valid:  false,
			symbol: symbolinfo{},
		}
		for column := 0; column < len(array); column++ {
			cell := array[row][column]
			char := []rune(cell)[0]
			if unicode.IsDigit(char) {
				inDots = false
				if !part.valid || part.number == "" {
					part.valid, part.symbol = checkAround(array, row, column)
				}
				part.number += cell
			} else {
				if !inDots {
					if part.symbol.sign == "*" {
						output += fmt.Sprintf("<blue>%s</>", part.number)
					} else {
						output += fmt.Sprintf("<red>%s</>", part.number)
					}
					if part.symbol.sign == "*" {
						parts = append(parts, part)
					}
					part = partinfo{}
				}
				if cell == "*" {
					output += fmt.Sprintf("<green>%s</>", cell)
					continue
				} else {
					output += cell
				}
				inDots = true
			}

			if column == len(array)-1 {
				if part.symbol.sign == "*" {
					output += fmt.Sprintf("<blue>%s</>", part.number)
				} else {
					output += fmt.Sprintf("<red>%s</>", part.number)
				}
				if part.symbol.sign == "*" {
					parts = append(parts, part)
				}
			}
		}
		output += "\n"
	}

	sort.Slice(parts, func(i, j int) bool {
		if parts[i].symbol.row == parts[j].symbol.row {
			return parts[i].symbol.column < parts[j].symbol.column
		}
		return parts[i].symbol.row < parts[j].symbol.row
	})

	for i := 0; i < len(parts)-1; i++ {
		if parts[i].symbol.column == parts[i+1].symbol.column && parts[i].symbol.row == parts[i+1].symbol.row {
			firstNumber, _ := strconv.Atoi(parts[i].number)
			lastNumber, _ := strconv.Atoi(parts[i+1].number)
			result += firstNumber * lastNumber
		}
	}

	output += strconv.Itoa(result)

	return output
}

func initArray(s string) [][]string {
	array := [][]string{}

	lines := strings.Split(s, "\n")

	for _, line := range lines {
		array = append(array, strings.Split(line, ""))
	}

	return array
}

// Checks if theres a symbol around the given number and returns the info about it
func checkAround(array [][]string, row, column int) (bool, symbolinfo) {
	for i := row - 1; i <= row+1; i++ {
		for j := column - 1; j <= column+1; j++ {
			if i < 0 || j < 0 || i >= len(array) || j >= len(array[0]) {
				continue
			}

			cell := array[i][j]
			char := []rune(cell)[0]

			if !unicode.IsDigit(char) && cell != "." {

				return true, symbolinfo{
					sign:   cell,
					row:    i,
					column: j,
				}
			}
		}
	}
	return false, symbolinfo{}
}
