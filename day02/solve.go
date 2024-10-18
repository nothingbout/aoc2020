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
	entries []PasswordEntry
}

type puzzleAnswer = int64

type PasswordEntry struct {
	minReqCount int64
	maxReqCount int64
	reqLetter   string
	password    string
}

func loadInput(inputLines []string) puzzleInput {
	var input puzzleInput
	for _, line := range inputLines {
		// 1-3 a: abcde
		var entry PasswordEntry
		mainCols := strings.Split(line, " ")
		countCols := strings.Split(mainCols[0], "-")
		entry.minReqCount = lib.Must(lib.ParseInt64(countCols[0]))
		entry.maxReqCount = lib.Must(lib.ParseInt64(countCols[1]))
		entry.reqLetter = strings.TrimSuffix(mainCols[1], ":")
		entry.password = mainCols[2]
		input.entries = append(input.entries, entry)
	}
	return input
}

func solvePart1(input puzzleInput) (answer puzzleAnswer) {
	var validCount int64
	for _, entry := range input.entries {
		remaining := entry.password
		var reqCount int64
		for {
			idx := strings.Index(remaining, entry.reqLetter)
			if idx < 0 {
				break
			}
			reqCount++
			remaining = remaining[idx+1:]
		}
		if reqCount >= entry.minReqCount && reqCount <= entry.maxReqCount {
			validCount++
		}
	}
	return validCount
}

func solvePart2(input puzzleInput) (answer puzzleAnswer) {
	var validCount int64
	for _, entry := range input.entries {
		a := entry.password[entry.minReqCount-1:entry.minReqCount] == entry.reqLetter
		b := entry.password[entry.maxReqCount-1:entry.maxReqCount] == entry.reqLetter
		if (a || b) && !(a && b) {
			validCount++
		}
	}
	return validCount
}

func main() {
	{ // part 1
		runner := &lib.Runner[puzzleInput, puzzleAnswer]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart1,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 2)
		runner.Run("files/input.txt", 548)
	}

	{ // part 2
		runner := &lib.Runner[puzzleInput, puzzleAnswer]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart2,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 1)
		runner.Run("files/input.txt", 502)
	}
}
