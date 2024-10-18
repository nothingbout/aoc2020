package main

import (
	lib "aoc2020/lib"
	v2 "aoc2020/lib/vec2int"
	"embed"
)

var (
	//go:embed "files/*"
	inputFileSystem embed.FS
)

type puzzleInput struct {
	gridSize v2.Vec
	trees    map[v2.Vec]bool
}

type puzzleAnswer = int64

func loadInput(inputLines []string) puzzleInput {
	input := puzzleInput{
		trees: map[v2.Vec]bool{},
	}
	for y, line := range inputLines {
		input.gridSize.X = int64(len(line))
		input.gridSize.Y++
		for x, c := range line {
			if string(c) == "#" {
				input.trees[v2.Make(int64(x), int64(y))] = true
			}
		}
	}
	return input
}

func traverseSlope(input puzzleInput, slope v2.Vec) int64 {
	curPos := v2.Make(0, 0)
	var treeCount int64
	for {
		curPos = v2.Add(curPos, slope)
		if curPos.Y >= input.gridSize.Y {
			break
		}
		if input.trees[v2.Make(curPos.X%input.gridSize.X, curPos.Y)] {
			treeCount++
		}
	}
	return treeCount
}

func solvePart1(input puzzleInput) (answer puzzleAnswer) {
	return traverseSlope(input, v2.Make(3, 1))
}

func solvePart2(input puzzleInput) (answer puzzleAnswer) {
	slopes := []v2.Vec{
		{X: 1, Y: 1},
		{X: 3, Y: 1},
		{X: 5, Y: 1},
		{X: 7, Y: 1},
		{X: 1, Y: 2},
	}
	var treeCountProd int64 = 1
	for _, slope := range slopes {
		treeCountProd *= traverseSlope(input, slope)
	}
	return treeCountProd
}

func main() {
	{ // part 1
		runner := &lib.Runner[puzzleInput, puzzleAnswer]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart1,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 7)
		runner.Run("files/input.txt", 274)
	}

	{ // part 2
		runner := &lib.Runner[puzzleInput, puzzleAnswer]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart2,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 336)
		runner.Run("files/input.txt", 6050183040)
	}
}
