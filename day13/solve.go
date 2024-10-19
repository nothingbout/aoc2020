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
	earliestDepartTime int64
	busIds             []int64
}

func loadInput(inputLines []string) puzzleInput {
	input := puzzleInput{}
	input.earliestDepartTime = lib.Must(lib.ParseInt64(inputLines[0]))
	for _, col := range strings.Split(inputLines[1], ",") {
		id := int64(0)
		if col != "x" {
			id = lib.Must(lib.ParseInt64(col))
		}
		input.busIds = append(input.busIds, id)
	}
	return input
}

func solvePart1(input puzzleInput) (answer int64) {
	for time := input.earliestDepartTime; ; time++ {
		for _, busId := range input.busIds {
			if busId == 0 {
				continue
			}
			if time%busId == 0 {
				return (time - input.earliestDepartTime) * busId
			}
		}
	}
}

func solvePart2(input puzzleInput) (answer int64) {
	var increment int64 = 1
	foundTimeFor := make([]bool, len(input.busIds))

	for time := int64(1); ; time += increment {
		isValidTime := true
		for idx, id := range input.busIds {
			if id == 0 || foundTimeFor[idx] {
				continue
			}
			if (time+int64(idx))%id == 0 {
				foundTimeFor[idx] = true
				increment = lib.LCD(increment, id)
			} else {
				isValidTime = false
			}
		}
		if isValidTime {
			return time
		}
	}
}

func main() {
	{ // part 1
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart1,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 295)
		runner.Run("files/input.txt", 246)
	}

	{ // part 2
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart2,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 1068781)
		runner.Run("files/input.txt", 939490236001473)
	}
}
