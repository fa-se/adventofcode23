package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	re := regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine|\\d)")

	numberMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	// map spelled-out numbers to the corresponding digit
	replaceFunc := func(match string) string {
		value, exists := numberMap[match]
		if !exists {
			// if the input is not in the map, assume it is already a digit and return it unchanged
			return match
		}
		return value
	}

	calibrationSum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var numbers []string

		// because the input may contain overlapping matches, we need to iteratively process each line
		index := 0
		for {
			match := re.FindStringIndex(line[index:])
			if match == nil {
				break
			}
			// Adjust the index relative to the start of the line
			match[0] += index
			match[1] += index

			numbers = append(numbers, line[match[0]:match[1]])

			index = match[0] + 1
		}

		// get first and last numbers, replacing spelled-out numbers with digits
		first := replaceFunc(numbers[0])
		last := replaceFunc(numbers[len(numbers)-1])
		// concatenate digits and convert to int
		calibrationValue, _ := strconv.Atoi(first + last)

		calibrationSum += calibrationValue
	}
	print(calibrationSum)
}
