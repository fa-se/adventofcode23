package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open(fmt.Sprint("input.txt"))
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)

	var lineResults []LineResult
	for scanner.Scan() {
		line := scanner.Text()
		lineData := parseLine(line)
		lineResults = append(lineResults, processLine(lineData))
	}
	resultPart1, resultPart2 := processResults(lineResults)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", resultPart1, resultPart2)
}

type Game struct {
	id       int
	cubeSets []CubeSet
}

func (g Game) String() string {
	return fmt.Sprintf("Game %d: %s", g.id, g.cubeSets)
}

type CubeSet struct {
	red   int
	green int
	blue  int
}

func (cs CubeSet) String() string {
	return fmt.Sprintf("(%dr,%dg,%db)", cs.red, cs.green, cs.blue)
}

type LineResult struct {
	gameID     int
	isPossible bool
	power      int
}

func (lr LineResult) String() string {
	return fmt.Sprintf("Game %d: isPossible: %t power: %d", lr.gameID, lr.isPossible, lr.power)
}

func parseLine(line string) Game {
	game := Game{}

	// Extract the game id
	re := regexp.MustCompile(`Game (\d+): (.*)`)
	matches := re.FindStringSubmatch(line)

	if len(matches) < 3 {
		fmt.Println("Invalid line format")
	}

	game.id, _ = strconv.Atoi(matches[1])

	// Split the rest of the string into cube sets
	cubeSets := strings.Split(matches[2], ";")

	for _, set := range cubeSets {
		cubeSet := CubeSet{}
		// Split the cube set into red, blue, and green cubes
		cubes := strings.Split(strings.TrimSpace(set), ",")

		for _, cube := range cubes {
			parts := strings.Fields(strings.TrimSpace(cube))
			count, _ := strconv.Atoi(parts[0])
			color := parts[1]

			switch color {
			case "red":
				cubeSet.red = count
			case "green":
				cubeSet.green = count
			case "blue":
				cubeSet.blue = count
			}
		}
		game.cubeSets = append(game.cubeSets, cubeSet)
	}

	return game
}

func processLine(data Game) LineResult {
	result := LineResult{data.id, true, 0}

	redLimit, greenLimit, blueLimit := 12, 13, 14
	redMax, greenMax, blueMax := 0, 0, 0

	for _, cubeSet := range data.cubeSets {
		// a game is possible, if for all sets the number of cubes does not exceed the limits for each color
		result.isPossible = result.isPossible && (cubeSet.red <= redLimit && cubeSet.blue <= blueLimit && cubeSet.green <= greenLimit)

		// part 2 requires finding for each color the maximum of cubes of that color
		redMax = max(redMax, cubeSet.red)
		greenMax = max(greenMax, cubeSet.green)
		blueMax = max(blueMax, cubeSet.blue)
	}

	// power is defined as the product of number of red, green, and blue cubes
	result.power = redMax * greenMax * blueMax
	return result
}

func processResults(results []LineResult) (int, int) {
	var part1, part2 int

	for _, result := range results {
		if result.isPossible {
			part1 += result.gameID
		}
	}

	for _, result := range results {
		part2 += result.power
	}

	return part1, part2
}
