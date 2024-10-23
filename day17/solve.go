package main

import (
	lib "aoc2020/lib"
	v3 "aoc2020/lib/vec/int3"
	"embed"
	"fmt"
)

var (
	//go:embed "files/*"
	inputFileSystem embed.FS
)

type puzzleInput struct {
	gridSize v3.Vec
	grid     mapGrid[bool]
}

type mapGrid[T any] map[v3.Vec]T

func loadInput(inputLines []string) puzzleInput {
	input := puzzleInput{}
	input.grid = mapGrid[bool]{}
	input.gridSize.Z = 1
	for y, line := range inputLines {
		input.gridSize.X = int64(len(line))
		input.gridSize.Y++

		for x, rune := range line {
			char := string(rune)
			if char == "#" {
				input.grid[v3.Make(int64(x), int64(y), 0)] = true
			}
		}
	}
	return input
}

func sprintBoolValue(value bool) string {
	if value {
		return "#"
	}
	return "."
}

func sprintGrid[T comparable](bounds v3.Bounds, grid mapGrid[T], sprintCell func(T) string) string {
	return fmt.Sprintln("\n" + v3.SprintGrid(bounds.Size, func(pos v3.Vec) string { return sprintCell(grid[v3.Add(bounds.Pos, pos)]) }))
}

func simulateCycle(bounds v3.Bounds, grid mapGrid[bool]) (v3.Bounds, mapGrid[bool]) {
	newBoundsMin := v3.MaxValue()
	newBoundsMax := v3.MinValue()
	newGrid := lib.CloneMap(grid)
	nOffsets := v3.AdjOffsets26()

	boundsMin := bounds.MinPos()
	boundsMax := bounds.MaxPos()

	for z := boundsMin.Z - 1; z <= boundsMax.Z+1; z++ {
		for y := boundsMin.Y - 1; y <= boundsMax.Y+1; y++ {
			for x := boundsMin.X - 1; x <= boundsMax.X+1; x++ {
				pos := v3.Make(x, y, z)
				var nActive int64
				for _, noff := range nOffsets {
					if grid[v3.Add(pos, noff)] {
						nActive++
					}
				}

				activeInNew := false

				if grid[pos] {
					if nActive == 2 || nActive == 3 {
						activeInNew = true
					} else {
						delete(newGrid, pos)
					}
				} else {
					if nActive == 3 {
						activeInNew = true
						newGrid[pos] = true
					}
				}

				if activeInNew {
					newBoundsMin = v3.Min(newBoundsMin, pos)
					newBoundsMax = v3.Max(newBoundsMax, pos)
				}
			}
		}
	}
	if len(newGrid) == 0 {
		panic("expected to have a non-empty grid")
	}
	newBounds := v3.MakeBounds(newBoundsMin, v3.Sub(v3.Add(newBoundsMax, v3.Make(1, 1, 1)), newBoundsMin))
	return newBounds, newGrid
}

func solvePart1(input puzzleInput) (answer int64) {
	bounds := v3.MakeBounds(v3.Zero(), input.gridSize)
	grid := lib.CloneMap(input.grid)
	// log.Printf("%s\n", sprintGrid(bounds, grid, sprintBoolValue))
	for cycleIdx := 0; cycleIdx < 6; cycleIdx++ {
		newBounds, newGrid := simulateCycle(bounds, grid)
		bounds = newBounds
		grid = newGrid
		// log.Printf("After %d cycles:\nBounds: %v\nActive: %d\n%s", cycleIdx+1, bounds, len(grid), sprintGrid(bounds, grid, sprintBoolValue))
	}
	return int64(len(grid))
}

func solvePart2(input puzzleInput) (answer int64) {
	return
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
		// runner := &lib.Runner[puzzleInput, int64]{
		// 	InputFileSystem: inputFileSystem,
		// 	InputLoader:     loadInput,
		// 	Solver:          solvePart2,
		// 	LogInput:        false,
		// }

		// runner.Run("files/example.txt", 0)
		// runner.Run("files/input.txt", 0)
	}
}
