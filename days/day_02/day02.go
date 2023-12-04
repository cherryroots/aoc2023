package day2

import (
	"strconv"
	"strings"
)

func SolvePart1(input string) string {
	lines := strings.Split(input, "\n")
	//lineLength := len(lines)
	maxRed := 12
	maxGreen := 13
	maxBlue := 14
	result := 0
	for _, line := range lines {
		validGame := false
		id, rounds := splitGame(line)

		for _, round := range rounds {
			red, green, blue := checkBagDraw(round)
			if red > maxRed || green > maxGreen || blue > maxBlue {
				validGame = false
				break
			} else {
				validGame = true
			}
		}

		if validGame {
			result += id
		}
	}

	return strconv.Itoa(result)
}

func SolvePart2(input string) string {
	lines := strings.Split(input, "\n")
	//lineLength := len(lines)
	power := 0
	for _, line := range lines {
		_, rounds := splitGame(line)
		var minimumReds, minimumGreens, minimumBlues = 0, 0, 0

		for _, round := range rounds {
			red, green, blue := checkBagDraw(round)
			if red > minimumReds {
				minimumReds = red
			}
			if green > minimumGreens {
				minimumGreens = green
			}
			if blue > minimumBlues {
				minimumBlues = blue
			}
		}

		power += minimumReds * minimumGreens * minimumBlues
	}

	return strconv.Itoa(power)
}

func splitGame(input string) (int, []string) {
	game := strings.Split(input, ":")[0]
	game = strings.Split(game, " ")[1]
	gameID, _ := strconv.Atoi(game)
	allRounds := strings.Split(input, ":")[1]
	rounds := strings.Split(allRounds, ";")
	for i := range rounds {
		rounds[i] = strings.TrimSpace(rounds[i])
	}

	return gameID, rounds
}

func checkBagDraw(round string) (int, int, int) {
	var redCubes int
	var greenCubes int
	var blueCubes int
	cubes := strings.Split(round, ",")
	for _, cube := range cubes {
		cube = strings.TrimSpace(cube)
		numCubes, _ := strconv.Atoi(strings.Split(cube, " ")[0])
		cubeType := strings.Split(cube, " ")[1]
		switch cubeType {
		case "red":
			redCubes += numCubes
		case "green":
			greenCubes += numCubes
		case "blue":
			blueCubes += numCubes
		}
	}
	return redCubes, greenCubes, blueCubes
}
