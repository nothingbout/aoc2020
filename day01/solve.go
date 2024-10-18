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
	entries []int64
}

type puzzleAnswer = int64

func loadInput(inputLines []string) puzzleInput {
	var input puzzleInput
	for _, line := range inputLines {
		num := lib.Must(lib.ParseInt64(line))
		input.entries = append(input.entries, num)
	}
	return input
}

func solvePart1(input puzzleInput) (answer puzzleAnswer) {
	for i, x := range input.entries {
		for j, y := range input.entries {
			if i != j && x+y == 2020 {
				return x * y
			}
		}
	}
	return
}

func solvePart2(input puzzleInput) (answer puzzleAnswer) {
	for i, x := range input.entries {
		for j, y := range input.entries {
			for k, z := range input.entries {
				if i != j && j != k && k != i && x+y+z == 2020 {
					return x * y * z
				}
			}
		}
	}
	return
}

func main() {
	{ // part 1
		runner := &lib.Runner[puzzleInput, puzzleAnswer]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart1,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 514579)
		runner.Run("files/input.txt", 0)
	}

	{ // part 2
		runner := &lib.Runner[puzzleInput, puzzleAnswer]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart2,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 241861950)
		runner.Run("files/input.txt", 0)
	}
}
