package solver

import (
	"fmt"
	"regexp"
	"strconv"
)

func Solver(corruptMemory string) int {
	enabled := true
	finalTotal := 0

	re := regexp.MustCompile(`(do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\))`)

	matches := re.FindAllStringSubmatch(corruptMemory, -1)

	for _, match := range matches {
		instruction := match[1]

		switch instruction {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if match[2] != "" && match[3] != "" {
				if enabled {
					num1, err1 := strconv.Atoi(match[2])
					num2, err2 := strconv.Atoi(match[3])

					if err1 != nil || err2 != nil {
						// This should not happen with the current regex, but good practice
						fmt.Printf(
							"Warning: Error converting numbers in matched mul: %s (%v, %v)\n",
							instruction,
							err1,
							err2,
						)
						continue // Skip this malformed match
					}
					result := num1 * num2
					finalTotal += result

				} else {
					// fmt.Printf("  Action: Mul instruction %s ignored (disabled)\n", instruction) // Debugging
				}
			} else {
				// This case should technically not be reached if the regex is correct
				fmt.Printf(
					"Warning: Regex matched '%s' but number groups are missing.\n",
					instruction,
				)
			}
		}
	}
	return finalTotal
}
