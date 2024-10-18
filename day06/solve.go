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
	answersByGroup [][]string
}

func loadInput(inputLines []string) puzzleInput {
	input := puzzleInput{}
	input.answersByGroup = lib.SplitLines(inputLines, "")
	return input
}

func solvePart1(input puzzleInput) (answer int64) {
	var total int64
	for _, groupAnswers := range input.answersByGroup {
		questionsMap := map[string]bool{}
		for _, personAnswers := range groupAnswers {
			for _, qr := range personAnswers {
				q := string(qr)
				questionsMap[q] = true
			}
		}
		total += int64(len(questionsMap))
	}
	return total
}

func solvePart2(input puzzleInput) (answer int64) {
	var total int64
	for _, groupAnswers := range input.answersByGroup {
		questionsMap := map[string]int{}
		for _, personAnswers := range groupAnswers {
			for _, qr := range personAnswers {
				q := string(qr)
				questionsMap[q] = questionsMap[q] + 1
			}
		}
		for _, value := range questionsMap {
			if value == len(groupAnswers) {
				total++
			}
		}
	}
	return total
}

func main() {
	{ // part 1
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart1,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 11)
		runner.Run("files/input.txt", 6590)
	}

	{ // part 2
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart2,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 6)
		runner.Run("files/input.txt", 3288)
	}
}
