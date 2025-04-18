package main

import (
	"fmt"
	"os"

	"github.com/jamesfulreader/solver"
	"github.com/jamesfulreader/utils"
)

func main() {
	fmt.Println("Mull it Over")

	inputFilePath := "./input.txt"

	fullCorruptMem, err := utils.ReadLines(inputFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file:%v\n", err)
		os.Exit(1)
	}

	answer := solver.Solver(fullCorruptMem)

	fmt.Println("answer ", answer)
}
