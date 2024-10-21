package main

import (
	lib "aoc2020/lib"
	"embed"
	"strings"
)

var (
	//go:embed "files/*"
	inputFileSystem embed.FS
)

type puzzleInput struct {
	startingNumbers []int64
}

func loadInput(inputLines []string) puzzleInput {
	input := puzzleInput{}
	for _, col := range strings.Split(inputLines[0], ",") {
		input.startingNumbers = append(input.startingNumbers, lib.Must(lib.ParseInt64(col)))
	}
	return input
}

func solveBruteforce(startingNumbers []int64, endTurnCount int64) int64 {
	spokenNumbers := lib.CloneSlice(startingNumbers)
	for turnIdx := int64(len(spokenNumbers) - 1); ; turnIdx++ {
		lastSpoken := spokenNumbers[turnIdx]
		interval := int64(0)
		for i := turnIdx - 1; i >= 0; i-- {
			if spokenNumbers[i] == lastSpoken {
				interval = turnIdx - i
				break
			}
		}
		spokenNumbers = append(spokenNumbers, interval)
		if int64(len(spokenNumbers)) == endTurnCount {
			break
		}
	}
	return spokenNumbers[len(spokenNumbers)-1]
}

func solveOptimized(startingNumbers []int64, endTurnCount int64) int64 {
	spokenNumbers := lib.CloneSlice(startingNumbers)
	lastSpokenTurnIndices := make([]int64, endTurnCount)
	for i := 0; i < len(lastSpokenTurnIndices); i++ {
		lastSpokenTurnIndices[i] = -1
	}
	for i, num := range startingNumbers {
		lastSpokenTurnIndices[num] = int64(i)
	}

	for turnIdx := int64(len(spokenNumbers) - 1); ; turnIdx++ {
		lastSpokenNum := spokenNumbers[turnIdx]
		interval := int64(0)
		if lastSpokenTurnIdx := lastSpokenTurnIndices[lastSpokenNum]; lastSpokenTurnIdx >= 0 {
			interval = turnIdx - lastSpokenTurnIdx
		}
		lastSpokenTurnIndices[lastSpokenNum] = turnIdx
		spokenNumbers = append(spokenNumbers, interval)
		if int64(len(spokenNumbers)) == endTurnCount {
			break
		}
	}
	return spokenNumbers[len(spokenNumbers)-1]
}

func solvePart1(input puzzleInput) (answer int64) {
	return solveBruteforce(input.startingNumbers, 2020)
}

func solvePart2(input puzzleInput) (answer int64) {
	return solveOptimized(input.startingNumbers, 30000000)
}

func main() {
	{ // part 1
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart1,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 436)
		runner.Run("files/input.txt", 371)
	}

	{ // part 2
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart2,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 175594)
		runner.Run("files/input.txt", 352)
	}
}
