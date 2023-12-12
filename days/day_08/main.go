package day08

import (
	"fmt"
	"strings"
	"unicode"
)

type node struct {
	instructions []string
	insIndex     int
	elements     map[string]element
	curElement   string
	steps        int
}

type element struct {
	left  string
	right string
}

func (n *node) Step() {
	element := n.elements[n.curElement]
	if n.insIndex >= len(n.instructions) {
		n.insIndex = 0
	}

	instruction := n.instructions[n.insIndex]
	if instruction == "R" {
		n.curElement = element.right
	} else {
		n.curElement = element.left
	}

	n.insIndex++
	n.steps++
}

func (n *node) findZ() {
	for n.curElement[2] != 'Z' {
		n.Step()
	}
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
	baseNode := parseInput(input)
	var ghost []node

	for k := range baseNode.elements {
		if k[2] == 'A' {
			newNode := baseNode
			newNode.curElement = k
			ghost = append(ghost, newNode)
		}
	}

	output := ""
	var multiples []int
	for i := 0; i < len(ghost); i++ {
		ghost[i].findZ()
		output += fmt.Sprintf("<blue>%d</> ", ghost[i].steps)
		multiples = append(multiples, ghost[i].steps)
	}

	output += fmt.Sprintf("\n<red>%d</>", LCM(multiples[0], multiples[1], multiples[2:]...))

	return output
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
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
			return !(unicode.IsLetter(c) || unicode.IsDigit(c))
		})
		output.elements[fields[0]] = element{left: fields[1], right: fields[2]}
	}

	return output
}
