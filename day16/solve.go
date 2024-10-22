package main

import (
	lib "aoc2020/lib"
	"embed"
	"slices"
	"strings"
)

var (
	//go:embed "files/*"
	inputFileSystem embed.FS
)

type puzzleInput struct {
	fieldRules    []fieldRule
	myTicket      ticket
	nearbyTickets []ticket
}

type ticket struct {
	values []int64
}

type fieldRule struct {
	name   string
	ranges []fieldRange
}

type fieldRange struct {
	min int64
	max int64
}

func loadInput(inputLines []string) puzzleInput {
	input := puzzleInput{}

	lineGroups := lib.SplitLines(inputLines, "")
	for _, line := range lineGroups[0] {
		// class: 1-3 or 5-7
		var fieldRule fieldRule
		cols := strings.Split(line, ": ")
		fieldRule.name = cols[0]

		rangeStrs := strings.Split(cols[1], " or ")
		for _, rangeStr := range rangeStrs {
			rangeCols := strings.Split(rangeStr, "-")
			var fieldRange fieldRange
			fieldRange.min = lib.Must(lib.ParseInt64(rangeCols[0]))
			fieldRange.max = lib.Must(lib.ParseInt64(rangeCols[1]))
			fieldRule.ranges = append(fieldRule.ranges, fieldRange)
		}
		input.fieldRules = append(input.fieldRules, fieldRule)
	}

	{
		//your ticket:
		//7,1,14
		line := lineGroups[1][1]
		input.myTicket.values = lib.MapSlice(strings.Split(line, ","), func(str string) int64 {
			return lib.Must(lib.ParseInt64(str))
		})
	}

	for _, line := range lineGroups[2][1:] {
		//nearby tickets:
		//7,3,47
		var ticket ticket
		ticket.values = lib.MapSlice(strings.Split(line, ","), func(str string) int64 {
			return lib.Must(lib.ParseInt64(str))
		})
		input.nearbyTickets = append(input.nearbyTickets, ticket)
	}
	return input
}

func isValidValueByRule(value int64, fieldRule fieldRule) bool {
	for _, fieldRange := range fieldRule.ranges {
		if value >= fieldRange.min && value <= fieldRange.max {
			return true
		}
	}
	return false
}

func isValidValueByRules(value int64, fieldRules []fieldRule) bool {
	for _, fieldRule := range fieldRules {
		if isValidValueByRule(value, fieldRule) {
			return true
		}
	}
	return false
}

func sumTicketInvalidValues(ticket ticket, fieldRules []fieldRule) (int64, bool) {
	foundInvalidValue := false
	var invalidSum int64
	for _, value := range ticket.values {
		if !isValidValueByRules(value, fieldRules) {
			foundInvalidValue = true
			invalidSum += value
		}
	}
	return invalidSum, foundInvalidValue
}

func solvePart1(input puzzleInput) (answer int64) {
	var totalSum int64
	for _, ticket := range input.nearbyTickets {
		ticketInvalidSum, _ := sumTicketInvalidValues(ticket, input.fieldRules)
		totalSum += ticketInvalidSum
	}
	return totalSum
}

func solvePart2(input puzzleInput) (answer int64) {
	var validTickets []ticket
	validTickets = append(validTickets, input.myTicket)
	for _, ticket := range input.nearbyTickets {
		_, invalid := sumTicketInvalidValues(ticket, input.fieldRules)
		if !invalid {
			validTickets = append(validTickets, ticket)
		}
	}

	remainingRules := lib.CloneSlice(input.fieldRules)
	var remainingCols []int64
	for i := 0; i < len(input.fieldRules); i++ {
		remainingCols = append(remainingCols, int64(i))
	}
	ruleToColMap := map[string]int64{}

	for len(remainingRules) > 0 {
		foundExactlyOneValidCol := false

		for remRuleIdx, fieldRule := range remainingRules {
			var validCols []int64
			for _, colIdx := range remainingCols {
				isValidCol := true
				for _, ticket := range validTickets {
					if !isValidValueByRule(ticket.values[colIdx], fieldRule) {
						isValidCol = false
						break
					}
				}
				if isValidCol {
					validCols = append(validCols, colIdx)
				}
			}

			// not a valid assumption in general, but exactly one such column is always found for some rule in the provided input
			if len(validCols) == 1 {
				foundExactlyOneValidCol = true
				ruleToColMap[fieldRule.name] = validCols[0]
				remainingRules = lib.SliceRemoveAt(remainingRules, remRuleIdx)
				remainingCols = lib.SliceRemoveAt(remainingCols, slices.Index(remainingCols, validCols[0]))
				break
			}
		}

		if !foundExactlyOneValidCol {
			panic("expected to find exactly one valid col for some rule")
		}
	}

	departureProduct := int64(1)
	for ruleName, colIdx := range ruleToColMap {
		if strings.HasPrefix(ruleName, "departure") {
			departureProduct *= input.myTicket.values[colIdx]
		}
	}
	return departureProduct
}

func main() {
	{ // part 1
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart1,
			LogInput:        false,
		}

		runner.Run("files/example_p1.txt", 71)
		runner.Run("files/input.txt", 19093)
	}

	{ // part 2
		runner := &lib.Runner[puzzleInput, int64]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart2,
			LogInput:        false,
		}

		// runner.Run("files/example_p1.txt", 0)
		runner.Run("files/example_p2.txt", 1)
		runner.Run("files/input.txt", 5311123569883)
	}
}
