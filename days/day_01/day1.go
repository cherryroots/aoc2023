package day1

import (
	"strconv"
	"strings"
	"unicode"
)

var numberMap = map[string]int{
	"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4,
	"five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
}

func Calibrate(s string) string {
	lines := strings.Split(s, "\n")
	c := make(chan int, len(lines))
	go func() {
		for _, line := range lines {
			go correctElfMistake(line, c)
		}
	}()

	sum := 0
	for range lines {
		sum += <-c
	}

	return strconv.Itoa(sum)
}

func correctElfMistake(line string, c chan int) {
	var numberList []int
	for count, char := range strings.ToLower(line) {
		if unicode.IsDigit(char) {
			number, _ := strconv.Atoi(string(char))
			numberList = append(numberList, number)
		} else if unicode.IsLetter(char) {
			for word, number := range numberMap {
				if strings.HasPrefix(line[count:], word) {
					numberList = append(numberList, number)
				}
			}
		}
	}

	first := strconv.Itoa(numberList[0])
	last := strconv.Itoa(numberList[len(numberList)-1])
	combined := first + last
	number, _ := strconv.Atoi(combined)

	c <- number
}
