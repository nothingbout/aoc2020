package main

import (
	lib "aoc2020/lib"
	"embed"
	"slices"
)

var (
	//go:embed "files/*"
	inputFileSystem embed.FS
)

type puzzleInput struct {
	entries []int64
}

func loadInput(inputLines []string) puzzleInput {
	input := puzzleInput{}
	for _, line := range inputLines {
		input.entries = append(input.entries, lib.Must(lib.ParseInt64(line)))
	}
	return input
}

func makeFullListOfAdapters(input puzzleInput) []int64 {
	adapters := lib.CloneSlice(input.entries)
	adapters = append(adapters, 0)
	slices.Sort(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	return adapters
}

func solvePart1(input puzzleInput) (answer int64) {
	adapters := makeFullListOfAdapters(input)
	var dc1, dc3 int64
	for i := int64(1); i < int64(len(adapters)); i++ {
		diff := adapters[i] - adapters[i-1]
		if diff == 1 {
			dc1++
		}
		if diff == 3 {
			dc3++
		}
	}
	return dc1 * dc3
}

func solvePart2(input puzzleInput) (answer int64) {
	adapters := makeFullListOfAdapters(input)
	arrangements := make([]int64, len(adapters))
	arrangements[0] = 1
	for i := int64(1); i < int64(len(adapters)); i++ {
		for j := i - 1; j >= 0; j-- {
			diff := adapters[i] - adapters[j]
			if diff > 3 {
				break
			}
			arrangements[i] += arrangements[j]
		}
	}
	return arrangements[len(arrangements)-1]
}

func main() {
	{ // part 1
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart1,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 220)
		runner.Run("files/input.txt", 2244)
	}

	{ // part 2
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart2,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 19208)
		runner.Run("files/input.txt", 3947645370368)
	}
}
