package main

import (
	lib "aoc2020/lib"
	"embed"
	"math"
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

func findInvalidNumberIdx(input puzzleInput, preambleLength int64) int64 {
	for curIdx := preambleLength; curIdx < int64(len(input.entries)); curIdx++ {
		curValue := input.entries[curIdx]
		valid := false
		for i := curIdx - preambleLength; i < curIdx; i++ {
			for j := curIdx - preambleLength; j < curIdx; j++ {
				if i != j && input.entries[i]+input.entries[j] == curValue {
					valid = true
				}
			}
		}
		if !valid {
			return curIdx
		}
	}
	panic("expected to find an invalid number")
}

func solvePart1(input puzzleInput, preambleLength int64) (answer int64) {
	return input.entries[findInvalidNumberIdx(input, preambleLength)]
}

func solvePart2(input puzzleInput, preambleLength int64) (answer int64) {
	invalidIdx := findInvalidNumberIdx(input, preambleLength)
	invalidValue := input.entries[invalidIdx]
	for i := int64(0); i < invalidIdx; i++ {
		var sum int64
		var smallest int64 = math.MaxInt64
		var largest int64 = 0
		for j := int64(i); j < invalidIdx; j++ {
			value := input.entries[j]
			smallest = lib.Min(smallest, value)
			largest = lib.Max(largest, value)
			sum += value
			if sum > invalidValue {
				break
			}
			if sum == invalidValue {
				return smallest + largest
			}
		}
	}
	panic("expected to find an answer")
}

func main() {
	{ // part 1
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          nil,
			LogInput:        false,
		}

		runner.Solver = func(input puzzleInput) int64 { return solvePart1(input, 5) }
		runner.Run("files/example.txt", 127)
		runner.Solver = func(input puzzleInput) int64 { return solvePart1(input, 25) }
		runner.Run("files/input.txt", 90433990)
	}

	{ // part 2
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          nil,
			LogInput:        false,
		}

		runner.Solver = func(input puzzleInput) int64 { return solvePart2(input, 5) }
		runner.Run("files/example.txt", 62)
		runner.Solver = func(input puzzleInput) int64 { return solvePart2(input, 25) }
		runner.Run("files/input.txt", 11691646)
	}
}
