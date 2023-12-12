package day08

import (
	"fmt"
	"strings"
	"unicode"
)

type node struct {
	instructions []string
	elements     map[string]element
	curElement   string
	step         int
	steps        int
}

type element struct {
	left  string
	right string
}

func (n *node) Step() {
	element := n.elements[n.curElement]
	if n.step >= len(n.instructions) {
		n.step = 0
	}
	instruction := n.instructions[n.step]

	if instruction == "R" {
		n.curElement = element.right
	} else {
		n.curElement = element.left
	}

	n.step++
	n.steps++
}

func SolvePart1(input string) string {
	node := parseInput(input)

	for node.curElement != "ZZZ" {
		node.Step()
	}

	output := fmt.Sprintf("<red>%d</>", node.steps)

	return output
}

func SolvePart2(input string) string {
	return ""
}

func parseInput(input string) node {
	parts := strings.Split(input, "\n\n")
	var output node

	output.instructions = strings.Split(parts[0], "")
	output.curElement = "AAA"
	output.elements = make(map[string]element)

	lines := strings.Split(parts[1], "\n")

	for _, line := range lines {
		fields := strings.FieldsFunc(line, func(c rune) bool {
			return !unicode.IsLetter(c)
		})

		output.elements[fields[0]] = element{
			left:  fields[1],
			right: fields[2],
		}
	}

	return output
}
