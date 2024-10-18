package main

import (
	lib "aoc2020/lib"
	"embed"
	"fmt"
	"slices"
)

var (
	//go:embed "files/*"
	inputFileSystem embed.FS
)

type puzzleInput struct {
	entries []string
}

type puzzleAnswer = int64

func loadInput(inputLines []string) puzzleInput {
	input := puzzleInput{}
	input.entries = append(input.entries, inputLines...)
	return input
}

func calcBinary(str string, lowChar, highChar string, low, high int64) int64 {
	for _, rune := range str {
		char := string(rune)
		switch char {
		case lowChar:
			high = (low + high) / 2
		case highChar:
			low = (low+high)/2 + 1
		default:
			panic(fmt.Sprintf("invalid char %s, expected %s or %s", char, lowChar, highChar))
		}
		// log.Printf("char %s, low %d, high %d\n", char, low, high)
	}
	if low != high {
		panic(fmt.Sprintf("should be equal: low %d, high: %d", low, high))
	}
	return low
}

func calcRowCol(seatStr string) (int64, int64) {
	return calcBinary(seatStr[0:7], "F", "B", 0, 127), calcBinary(seatStr[7:10], "L", "R", 0, 7)
}

func solvePart1(input puzzleInput) (answer puzzleAnswer) {
	var highestId int64
	for _, seatStr := range input.entries {
		row, col := calcRowCol(seatStr)
		seatId := row*8 + col
		if seatId > highestId {
			highestId = seatId
		}
	}
	return highestId
}

func solvePart2(input puzzleInput) (answer puzzleAnswer) {
	var seatIds []int64
	for _, seatStr := range input.entries {
		row, col := calcRowCol(seatStr)
		seatId := row*8 + col
		seatIds = append(seatIds, seatId)
	}
	slices.Sort(seatIds)
	for i, id := range seatIds {
		if seatIds[i+1]-id > 1 {
			return id + 1
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

		runner.Run("files/example.txt", 357)
		runner.Run("files/example2.txt", 567)
		runner.Run("files/example3.txt", 119)
		runner.Run("files/example4.txt", 820)
		runner.Run("files/input.txt", 994)
	}

	{ // part 2
		runner := &lib.Runner[puzzleInput, puzzleAnswer]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart2,
			LogInput:        false,
		}

		// runner.Run("files/example.txt", 0)
		runner.Run("files/input.txt", 741)
	}
}
