package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
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

type LineData struct {
	symbols []Symbol
	numbers []Number
}

func (l LineData) String() string {
	return fmt.Sprintf("symbols: %s | numbers: %s", l.symbols, l.numbers)
}

type Symbol struct {
	char rune
	idx  int
}

func (s Symbol) String() string {
	return fmt.Sprintf("%c[%d]", s.char, s.idx)
}

type Number struct {
	number   int
	startIdx int
	stopIdx  int
}

func (n Number) String() string {
	return fmt.Sprintf("%d[%d-%d]", n.number, n.startIdx, n.stopIdx)
}

type LineResult = LineData

func parseLine(line string) LineData {
	var data LineData

	var numberString string
	for idx, char := range line {
		if unicode.IsDigit(char) {
			numberString += string(char)
		} else {
			if numberString != "" {
				number, _ := strconv.Atoi(numberString)
				endIdx := idx - 1
				startIdx := idx - len(numberString)
				data.numbers = append(data.numbers, Number{number, startIdx, endIdx})
				numberString = ""
			}
			if char != '.' {
				data.symbols = append(data.symbols, Symbol{char, idx})
			}
		}
		_ = idx
	}
	// handle numbers at line end
	if numberString != "" {
		number, _ := strconv.Atoi(numberString)
		endIdx := len(line) - 1
		startIdx := len(line) - len(numberString)
		data.numbers = append(data.numbers, Number{number, startIdx, endIdx})
	}

	return data
}

// processLineData takes an instance of LineData and processes it according to the day's puzzle requirements.
func processLine(data LineData) LineResult {
	// for now, do nothing because processing needs to be done across multiple lines

	return data
}

func processResults(results []LineResult) (int, int) {
	var part1, part2 int

	previousLine, nextLine := LineResult{}, LineResult{}
	for lineIdx, currentLine := range results {
		if lineIdx >= 1 {
			previousLine = results[lineIdx-1]
		}
		if lineIdx < len(results)-1 {
			nextLine = results[lineIdx+1]
		}

		// append symbols from all three lines
		symbols := append(append(previousLine.symbols, currentLine.symbols...), nextLine.symbols...)
		for _, number := range currentLine.numbers {
			for _, symbol := range symbols {
				if isAdjacent(number, symbol) {
					part1 += number.number
				}
			}
		}

		// append numbers from all three lines
		numbers := append(append(previousLine.numbers, currentLine.numbers...), nextLine.numbers...)
		for _, symbol := range currentLine.symbols {
			if symbol.char == '*' {
				adjacentNumbers := getAdjacentNumbers(symbol, numbers)
				if len(adjacentNumbers) == 2 {
					part2 += adjacentNumbers[0].number * adjacentNumbers[1].number
				}
			}
		}
	}

	return part1, part2
}

func getAdjacentNumbers(symbol Symbol, numbers []Number) []Number {
	adjacentNumbers := []Number{}

	for _, number := range numbers {
		if isAdjacent(number, symbol) {
			adjacentNumbers = append(adjacentNumbers, number)
		}
	}

	return adjacentNumbers
}

func isAdjacent(number Number, symbol Symbol) bool {
	if symbol.idx >= number.startIdx-1 && symbol.idx <= number.stopIdx+1 {
		return true
	}
	return false
}
