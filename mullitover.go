package mullitover

import (
	"regexp"
)

func PatternMatch(input string) [][]string {
	r, _ := regexp.Compile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	idxOfMultipliers := r.FindAllStringSubmatch(input, -1)
	return idxOfMultipliers
}
