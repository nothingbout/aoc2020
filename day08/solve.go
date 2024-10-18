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
	instructions []instruction
}

type instruction struct {
	op  string
	arg int64
}

func loadInput(inputLines []string) puzzleInput {
	input := puzzleInput{}
	for _, line := range inputLines {
		cols := strings.Split(line, " ")
		input.instructions = append(input.instructions, instruction{
			op:  cols[0],
			arg: lib.Must(lib.ParseInt64(cols[1])),
		})
	}
	return input
}

type execResult int

const (
	execResultUnknown execResult = iota
	execResultSuccess
	execResultInfiniteLoop
	execResultFail
)

func executeProgram(instructions []instruction) (execResult, int64) {
	var curLine int64 = 0
	var accValue int64 = 0
	isLineVisited := make([]bool, len(instructions))
	execResult := execResultUnknown
	for {
		if curLine == int64(len(instructions)) {
			execResult = execResultSuccess
			break
		}
		if curLine > int64(len(instructions)) {
			execResult = execResultFail
			break
		}
		if isLineVisited[curLine] {
			execResult = execResultInfiniteLoop
			break
		}
		isLineVisited[curLine] = true
		instruction := instructions[curLine]
		switch instruction.op {
		case "nop":
			curLine++
		case "acc":
			accValue += instruction.arg
			curLine++
		case "jmp":
			curLine += instruction.arg
		}
	}
	return execResult, accValue
}

func solvePart1(input puzzleInput) (answer int64) {
	result, accValue := executeProgram(input.instructions)
	if result != execResultInfiniteLoop {
		panic("expected infinite loop")
	}
	return accValue
}

func solvePart2(input puzzleInput) (answer int64) {
	for lineIdx, instruction := range input.instructions {
		newOp := ""
		switch instruction.op {
		case "nop":
			newOp = "jmp"
		case "jmp":
			newOp = "nop"
		default:
			continue
		}

		newInstructions := lib.CloneSlice(input.instructions)
		newInstructions[lineIdx].op = newOp

		result, accValue := executeProgram(newInstructions)
		if result == execResultSuccess {
			return accValue
		}
	}
	panic("expected a success")
}

func main() {
	{ // part 1
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart1,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 5)
		runner.Run("files/input.txt", 1744)
	}

	{ // part 2
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart2,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 8)
		runner.Run("files/input.txt", 1174)
	}
}
