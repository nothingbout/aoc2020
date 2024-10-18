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
	rules map[string][]nameAndQuantity
}

type nameAndQuantity struct {
	name     string
	quantity int64
}

func loadInput(inputLines []string) puzzleInput {
	input := puzzleInput{
		rules: map[string][]nameAndQuantity{},
	}
	for _, line := range inputLines {
		// faded blue bags contain no other bags.
		// bright white bags contain 1 shiny gold bag.
		// wavy red bags contain 5 shiny fuchsia bags, 1 dim bronze bag, 4 dim turquoise bags, 3 dotted violet bags.
		mainCols := strings.Split(line[:len(line)-1], " bags contain ")
		contentCols := strings.Split(mainCols[1], ", ")
		var namesAndQuantities []nameAndQuantity
		if contentCols[0] != "no other bags" {
			for _, contentCol := range contentCols {
				sepIdx := strings.Index(contentCol, " ")
				namesAndQuantities = append(namesAndQuantities, nameAndQuantity{
					name:     strings.TrimSuffix(strings.TrimSuffix(contentCol[sepIdx+1:], " bags"), " bag"),
					quantity: lib.Must(lib.ParseInt64(contentCol[:sepIdx])),
				})
			}
		}
		input.rules[mainCols[0]] = namesAndQuantities
	}
	return input
}

func findBagRec(input puzzleInput, curName string, targetName string) bool {
	if curName == targetName {
		return true
	}
	for _, nq := range input.rules[curName] {
		if findBagRec(input, nq.name, targetName) {
			return true
		}
	}
	return false
}

func solvePart1(input puzzleInput) (answer int64) {
	const targetName = "shiny gold"
	var total int64
	for name := range input.rules {
		if name != targetName && findBagRec(input, name, targetName) {
			total++
		}
	}
	return total
}

func countBagsRec(input puzzleInput, curName string) int64 {
	var innerCount int64
	for _, nq := range input.rules[curName] {
		innerCount += countBagsRec(input, nq.name) * nq.quantity
	}
	return 1 + innerCount
}

func solvePart2(input puzzleInput) (answer int64) {
	return countBagsRec(input, "shiny gold") - 1
}

func main() {
	{ // part 1
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart1,
			LogInput:        false,
		}

		runner.Run("files/example_p1.txt", 4)
		runner.Run("files/input.txt", 259)
	}

	{ // part 2
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart2,
			LogInput:        false,
		}

		runner.Run("files/example_p2.txt", 126)
		runner.Run("files/input.txt", 45018)
	}
}
