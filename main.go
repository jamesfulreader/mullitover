package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/jamesfulreader/mullitover"
)

func main() {
	fmt.Println("Mull it Over")

	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("error opening input.txt: %v", err)
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	finalTotal := 0
	for scanner.Scan() {
		corruptedMemoryString := scanner.Text()

		getMultipliers := mullitover.PatternMatch(corruptedMemoryString)
		extractMultiply, err := mullitover.ExtractNumbers(getMultipliers)
		if err != nil {
			fmt.Println("err ", err)
		}
		finalTotal += extractMultiply
	}
	fmt.Println("total is ", finalTotal)
}
