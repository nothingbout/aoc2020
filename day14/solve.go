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
	instructions []Instruction
}

type Instruction struct {
	op      InstructionOp
	mask    string
	address int64
	value   int64
}

type InstructionOp int

const (
	InstructionOpSetMask InstructionOp = iota
	InstructionOpSetValue
)

func loadInput(inputLines []string) puzzleInput {
	input := puzzleInput{}
	for _, line := range inputLines {
		// mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
		// mem[8] = 11

		var instruction Instruction

		cols := strings.Split(line, " = ")
		if cols[0] == "mask" {
			instruction.op = InstructionOpSetMask
			instruction.mask = cols[1]
		} else {
			instruction.op = InstructionOpSetValue
			instruction.address = lib.Must(lib.ParseInt64(cols[0][4 : len(cols[0])-1]))
			instruction.value = lib.Must(lib.ParseInt64(cols[1]))
		}

		input.instructions = append(input.instructions, instruction)
	}
	return input
}

func solvePart1(input puzzleInput) (answer int64) {
	mask := ""
	memory := map[int64]int64{}
	for _, instruction := range input.instructions {
		switch instruction.op {
		case InstructionOpSetMask:
			mask = instruction.mask
		case InstructionOpSetValue:
			value := instruction.value
			for idx, theRune := range mask {
				bitPos := len(mask) - 1 - idx
				char := string(theRune)
				switch char {
				case "0":
					value &= ^(1 << bitPos)
				case "1":
					value |= 1 << bitPos
				}
			}
			memory[instruction.address] = value
		}
	}
	var sum int64
	for _, value := range memory {
		sum += value
	}
	return sum
}

type rSetFloatingConsts struct {
	memory map[int64]int64
	mask   string
	value  int64
}

func rSetFloating(c *rSetFloatingConsts, maskIdx int64, address int64) {
	if maskIdx == int64(len(c.mask)) {
		// fmt.Printf("%036b (%d): %d\n", address, address, c.value)
		c.memory[address] = c.value
		return
	}

	bitPos := len(c.mask) - 1 - int(maskIdx)
	switch string(c.mask[maskIdx]) {
	case "0":
		rSetFloating(c, maskIdx+1, address)
	case "1":
		rSetFloating(c, maskIdx+1, address|(1<<bitPos))
	case "X":
		rSetFloating(c, maskIdx+1, address & ^(1<<bitPos))
		rSetFloating(c, maskIdx+1, address|(1<<bitPos))
	}
}

func solvePart2(input puzzleInput) (answer int64) {
	mask := ""
	memory := map[int64]int64{}
	for _, instruction := range input.instructions {
		switch instruction.op {
		case InstructionOpSetMask:
			mask = instruction.mask
		case InstructionOpSetValue:
			rSetFloating(
				&rSetFloatingConsts{memory: memory, mask: mask, value: instruction.value},
				0, instruction.address,
			)
		}
	}
	var sum int64
	for _, value := range memory {
		sum += value
	}
	return sum
}

func main() {
	{ // part 1
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart1,
			LogInput:        false,
		}

		runner.Run("files/example_p1.txt", 165)
		runner.Run("files/input.txt", 17765746710228)
	}

	{ // part 2
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart2,
			LogInput:        false,
		}

		runner.Run("files/example_p2.txt", 208)
		runner.Run("files/input.txt", 4401465949086)
	}
}
