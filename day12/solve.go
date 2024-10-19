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
	instructions []instruction
}

type instruction struct {
	action string
	value  int64
}

func loadInput(inputLines []string) puzzleInput {
	input := puzzleInput{}
	for _, line := range inputLines {
		input.instructions = append(input.instructions, instruction{
			action: line[:1],
			value:  lib.Must(lib.ParseInt64(line[1:])),
		})
	}
	return input
}

func dirToVec(dir string) v2.Vec {
	switch dir {
	case "E":
		return v2.Make(1, 0)
	case "S":
		return v2.Make(0, 1)
	case "W":
		return v2.Make(-1, 0)
	case "N":
		return v2.Make(0, -1)
	}
	panic("invalid direction")
}

func turnDir(dir string, turnDir string) string {
	switch turnDir {
	case "L":
		switch dir {
		case "E":
			return "N"
		case "N":
			return "W"
		case "W":
			return "S"
		case "S":
			return "E"
		}
	case "R":
		switch dir {
		case "E":
			return "S"
		case "S":
			return "W"
		case "W":
			return "N"
		case "N":
			return "E"
		}
	}
	panic("invalid direction or turn direction")
}

func solvePart1(input puzzleInput) (answer int64) {
	shipDir := "E"
	shipPos := v2.Zero()
	for _, instruction := range input.instructions {
		if instruction.action == "L" || instruction.action == "R" {
			for i := int64(0); i < instruction.value/90; i++ {
				shipDir = turnDir(shipDir, instruction.action)
			}
		} else if instruction.action == "F" {
			shipPos = v2.Add(shipPos, v2.Scale(dirToVec(shipDir), instruction.value))
		} else {
			shipPos = v2.Add(shipPos, v2.Scale(dirToVec(instruction.action), instruction.value))
		}
	}
	return lib.Abs(shipPos.X) + lib.Abs(shipPos.Y)
}

func solvePart2(input puzzleInput) (answer int64) {
	shipPos := v2.Zero()
	waypointPos := v2.Make(10, -1)
	for _, instruction := range input.instructions {
		if instruction.action == "L" || instruction.action == "R" {
			for i := int64(0); i < instruction.value/90; i++ {
				if instruction.action == "L" {
					waypointPos = v2.Make(waypointPos.Y, -waypointPos.X)
				} else {
					waypointPos = v2.Make(-waypointPos.Y, waypointPos.X)
				}
			}
		} else if instruction.action == "F" {
			shipPos = v2.Add(shipPos, v2.Scale(waypointPos, instruction.value))
		} else {
			waypointPos = v2.Add(waypointPos, v2.Scale(dirToVec(instruction.action), instruction.value))
		}
	}
	return lib.Abs(shipPos.X) + lib.Abs(shipPos.Y)
}

func main() {
	{ // part 1
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart1,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 25)
		// 181 too low
		runner.Run("files/input.txt", 1133)
	}

	{ // part 2
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart2,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 286)
		runner.Run("files/input.txt", 61053)
	}
}
