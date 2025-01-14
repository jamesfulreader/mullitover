package mullitover

import (
	"fmt"
	"regexp"
	"strconv"
)

func PatternMatch(input string) [][]string {
	// disable, _ := regexp.Compile(`don't()`)
	// idxOfDisables := disable.FindAllStringSubmatch(input, -1)
	// enable, _ := regexp.Compile(`do()`)
	// idxOfEnables := enable.FindAllStringSubmatch(input, -1)
	mul, _ := regexp.Compile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	idxOfMultipliers := mul.FindAllStringSubmatch(input, -1)
	return idxOfMultipliers
}

func ExtractNumbers(multiplierIndexes [][]string) (int, error) {
	total := 0
	for i, numbers := range multiplierIndexes {
		num1, err := strconv.Atoi(numbers[1])
		if err != nil {
			return 0, fmt.Errorf("error parsing first number in pair %d: %w", i, err)
		}

		num2, err := strconv.Atoi(numbers[2])
		if err != nil {
			return 0, fmt.Errorf("error parsing first number in pair %d: %w", i, err)
		}

		total += (num1 * num2)
	}
	return total, nil
}
