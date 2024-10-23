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
	gridSize Vec
	grid     map[Vec]bool
}

type Vec struct {
	X int64
	Y int64
	Z int64
	W int64
}

func MakeVec(X, Y, Z, W int64) Vec { return Vec{X, Y, Z, W} }
func AddVec(a, b Vec) Vec          { return MakeVec(a.X+b.X, a.Y+b.Y, a.Z+b.Z, a.W+b.W) }

func loadInput(inputLines []string) puzzleInput {
	input := puzzleInput{}
	input.grid = map[Vec]bool{}
	input.gridSize.Z = 1
	for y, line := range inputLines {
		input.gridSize.X = int64(len(line))
		input.gridSize.Y++

		for x, rune := range line {
			char := string(rune)
			if char == "#" {
				input.grid[MakeVec(int64(x), int64(y), 0, 0)] = true
			}
		}
	}
	return input
}

func simulateCycle(grid map[Vec]bool, offsets []Vec) map[Vec]bool {
	counts := map[Vec]int64{}
	for pos, _ := range grid {
		for _, off := range offsets {
			npos := AddVec(pos, off)
			counts[npos]++
		}
	}
	newGrid := map[Vec]bool{}
	for pos, count := range counts {
		if count == 3 {
			newGrid[pos] = true
		} else if count == 2 && grid[pos] {
			newGrid[pos] = true
		}
	}
	return newGrid
}

func makeNeighborOffsets(dimensions int64) (offsets []Vec) {
	wRange := 0
	if dimensions == 4 {
		wRange = 1
	}
	for w := -wRange; w <= wRange; w++ {
		for z := -1; z <= 1; z++ {
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					if w == 0 && z == 0 && y == 0 && x == 0 {
						continue
					}
					offsets = append(offsets, MakeVec(int64(x), int64(y), int64(z), int64(w)))
				}
			}
		}
	}
	return
}

func solve(input puzzleInput, dimensions int64) int64 {
	grid := lib.CloneMap(input.grid)
	offsets := makeNeighborOffsets(dimensions)
	for cycleIdx := 0; cycleIdx < 6; cycleIdx++ {
		newGrid := simulateCycle(grid, offsets)
		grid = newGrid
	}
	return int64(len(grid))
}

func solvePart1(input puzzleInput) (answer int64) {
	return solve(input, 3)
}

func solvePart2(input puzzleInput) (answer int64) {
	return solve(input, 4)
}

func main() {
	{ // part 1
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart1,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 112)
		runner.Run("files/input.txt", 348)
	}

	{ // part 2
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart2,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 848)
		runner.Run("files/input.txt", 2236)
	}
}
