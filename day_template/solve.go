package main

import (
	lib "aoc2020/lib"
	"embed"
)

var (
	//go:embed "files/*"
	inputFileSystem embed.FS
)

type puzzleInput struct {
	entries []string
}

func loadInput(inputLines []string) puzzleInput {
	input := puzzleInput{}
	for _, line := range inputLines {
		input.entries = append(input.entries, line)
	}
	return input
}

func solvePart1(input puzzleInput) (answer int64) {
	return
}

func solvePart2(input puzzleInput) (answer int64) {
	return
}

func main() {
	{ // part 1
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart1,
			LogInput:        true,
		}

		runner.Run("files/example.txt", 0)
		// runner.Run("files/input.txt", 0)
	}

	{ // part 2
		// runner := &lib.Runner[puzzleInput, int64]{
		// 	InputFileSystem: inputFileSystem,
		// 	InputLoader:     loadInput,
		// 	Solver:          solvePart2,
		// 	LogInput:        false,
		// }

		// runner.Run("files/example.txt", 0)
		// runner.Run("files/input.txt", 0)
	}
}
