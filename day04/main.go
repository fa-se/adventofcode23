package main

import (
	"bufio"
	"fmt"
	"math"
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

type Card struct {
	idx            int
	winningNumbers []int
	ownNumbers     []int
}

func (c Card) String() string {
	return fmt.Sprintf("Card %d: Winning Numbers: %v Own Numbers: %v",
		c.idx, c.winningNumbers, c.ownNumbers)
}

type LineResult struct {
	value int
}

func parseLine(line string) Card {
	card := Card{}

	re := regexp.MustCompile(`Card +(\d+): ([\d\s]+)\| ([\d\s]+)`)
	matches := re.FindStringSubmatch(line)

	if len(matches) < 4 {
		fmt.Println("Invalid line format: ", line)
	}

	card.idx, _ = strconv.Atoi(matches[1])
	card.winningNumbers = extractNumbers(matches[2])
	card.ownNumbers = extractNumbers(matches[3])

	return card
}

func extractNumbers(s string) []int {
	var numbers []int
	for _, str := range strings.Fields(s) {
		num, err := strconv.Atoi(str)
		if err == nil {
			numbers = append(numbers, num)
		}
	}
	return numbers
}

// processLineData takes an instance of LineData and processes it according to the day's puzzle requirements.
func processLine(data Card) LineResult {
	result := LineResult{value: 0}

	// map for efficient lookup
	winningNumbers := make(map[int]bool)
	for _, winningNumber := range data.winningNumbers {
		winningNumbers[winningNumber] = true
	}

	var matches []int
	for _, number := range data.ownNumbers {
		if winningNumbers[number] {
			matches = append(matches, number)
		}
	}

	if len(matches) > 0 {
		result.value = int(math.Pow(2, float64(len(matches)-1)))
	}

	return result
}

func processResults(results []LineResult) (int, int) {
	var part1, part2 int

	for _, result := range results {
		part1 += result.value
	}

	return part1, part2
}
