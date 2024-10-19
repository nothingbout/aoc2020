package main

import (
	lib "aoc2020/lib"
	v2 "aoc2020/lib/vec/int2"
	"embed"
)

var (
	//go:embed "files/*"
	inputFileSystem embed.FS
)

type puzzleInput struct {
	gridSize v2.Vec
	seats    gridMap
}

type gridMap = map[v2.Vec]bool

func loadInput(inputLines []string) puzzleInput {
	input := puzzleInput{
		seats: gridMap{},
	}
	for y, line := range inputLines {
		input.gridSize.X = int64(len(line))
		input.gridSize.Y++

		for x, rune := range line {
			char := string(rune)
			if char == "L" {
				input.seats[v2.Make(int64(x), int64(y))] = true
			}
		}
	}
	return input
}

func sprintGrid(gridSize v2.Vec, seatGrid, occGrid gridMap) string {
	return v2.SprintGrid(gridSize, func(pos v2.Vec) string {
		if occGrid[pos] {
			return "#"
		}
		if seatGrid[pos] {
			return "L"
		}
		return "."
	})
}

func iterateOcc(gridSize v2.Vec, seatGrid, occGrid gridMap, occThreshold int64, countOccupiedNeighbors func(occGrid gridMap, pos v2.Vec) int64) (bool, gridMap) {
	occChanged := false
	newOccGrid := lib.CloneMap(occGrid)
	for y := int64(0); y < gridSize.Y; y++ {
		for x := int64(0); x < gridSize.X; x++ {
			pos := v2.Make(x, y)
			if !seatGrid[pos] {
				continue
			}
			oc := countOccupiedNeighbors(occGrid, pos)
			if oc == 0 && !occGrid[pos] {
				occChanged = true
				newOccGrid[pos] = true
			}
			if oc >= occThreshold && occGrid[pos] {
				occChanged = true
				delete(newOccGrid, pos)
			}
		}
	}
	return occChanged, newOccGrid
}

func countAdjacent(grid gridMap, pos v2.Vec) int64 {
	var count int64
	for _, off := range v2.AdjOffsets8 {
		if grid[v2.Add(pos, off)] {
			count++
		}
	}
	return count
}

func solvePart1(input puzzleInput) (answer int64) {
	gridSize := input.gridSize
	seatGrid := input.seats
	occGrid := make(gridMap)

	// log.Printf("\n%s\n", sprintGrid(gridSize, seatGrid, occGrid))
	for {
		occChanged, newOccGrid := iterateOcc(gridSize, seatGrid, occGrid, 4, countAdjacent)
		if !occChanged {
			break
		}
		occGrid = newOccGrid
		// log.Printf("\n%s\n", sprintGrid(gridSize, seatGrid, occGrid))
	}

	return int64(len(occGrid))
}

func solvePart2(input puzzleInput) (answer int64) {
	gridSize := input.gridSize
	seatGrid := input.seats
	occGrid := make(gridMap)

	neighborsMap := map[v2.Vec][]v2.Vec{}

	for y := int64(0); y < gridSize.Y; y++ {
		for x := int64(0); x < gridSize.X; x++ {
			origin := v2.Make(x, y)
			if !seatGrid[origin] {
				continue
			}

			for _, off := range v2.AdjOffsets8 {
				offPos := origin
				for {
					offPos = v2.Add(offPos, off)
					if !v2.IsInBounds(offPos, v2.Zero(), gridSize) {
						break
					}
					if seatGrid[offPos] {
						neighborsMap[origin] = append(neighborsMap[origin], offPos)
						break
					}
				}
			}
		}
	}

	countOccupiedNeighbors := func(occ gridMap, pos v2.Vec) int64 {
		var count int64
		for _, offPos := range neighborsMap[pos] {
			if occ[offPos] {
				count++
			}
		}
		return count
	}

	// log.Printf("\n%s\n", sprintGrid(gridSize, seatGrid, occGrid))
	for {
		occChanged, newOccGrid := iterateOcc(gridSize, seatGrid, occGrid, 5, countOccupiedNeighbors)
		if !occChanged {
			break
		}
		occGrid = newOccGrid
		// log.Printf("\n%s\n", sprintGrid(gridSize, seatGrid, occGrid))
	}

	return int64(len(occGrid))
}

func main() {
	{ // part 1
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart1,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 37)
		runner.Run("files/input.txt", 2424)
	}

	{ // part 2
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart2,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 26)
		runner.Run("files/input.txt", 2208)
	}
}
