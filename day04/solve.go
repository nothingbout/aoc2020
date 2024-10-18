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
	passports []passport
}

type puzzleAnswer = int64

type passport struct {
	fields map[string]string
}

func loadInput(inputLines []string) puzzleInput {
	input := puzzleInput{}

	lineGroups := lib.SplitLines(inputLines, "")
	for _, lines := range lineGroups {
		passport := passport{
			fields: map[string]string{},
		}

		line := strings.Join(lines, " ")
		mainCols := strings.Split(line, " ")
		for _, col := range mainCols {
			fieldCols := strings.Split(col, ":")
			passport.fields[fieldCols[0]] = fieldCols[1]
		}
		input.passports = append(input.passports, passport)
	}
	return input
}

func solvePart1(input puzzleInput) (answer puzzleAnswer) {
	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		// "cid",
	}
	var validCount int64
	for _, passport := range input.passports {
		isValid := true
		for _, reqField := range requiredFields {
			_, hasField := passport.fields[reqField]
			if !hasField {
				isValid = false
				break
			}
		}
		if isValid {
			validCount++
		}
	}
	return validCount
}

func parseIntAndValidateRange(str string, min, max int64) bool {
	x, err := lib.ParseInt64(str)
	if err != nil {
		return false
	}
	return x >= min && x <= max
}

var (
	validHairColorChars = []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"a", "b", "c", "d", "e", "f",
	}
	validEyeColors = []string{
		"amb", "blu", "brn", "gry", "grn", "hzl", "oth",
	}
)

func validatePassport(passport passport) bool {
	if !parseIntAndValidateRange(passport.fields["byr"], 1920, 2002) {
		return false
	}

	if !parseIntAndValidateRange(passport.fields["iyr"], 2010, 2020) {
		return false
	}

	if !parseIntAndValidateRange(passport.fields["eyr"], 2020, 2030) {
		return false
	}

	height := passport.fields["hgt"]
	if strings.HasSuffix(height, "cm") {
		if !parseIntAndValidateRange(strings.TrimSuffix(height, "cm"), 150, 193) {
			return false
		}
	} else if strings.HasSuffix(height, "in") {
		if !parseIntAndValidateRange(strings.TrimSuffix(height, "in"), 59, 76) {
			return false
		}
	} else {
		return false
	}

	hairColor := passport.fields["hcl"]
	if strings.HasPrefix(hairColor, "#") && len(hairColor) == 7 {
		for _, c := range hairColor[1:] {
			if !slices.Contains(validHairColorChars, string(c)) {
				return false
			}
		}
	} else {
		return false
	}

	if !slices.Contains(validEyeColors, passport.fields["ecl"]) {
		return false
	}

	passportId := passport.fields["pid"]
	if len(passportId) == 9 {
		_, err := lib.ParseInt64(passportId)
		if err != nil {
			return false
		}
	} else {
		return false
	}

	return true
}

func solvePart2(input puzzleInput) (answer puzzleAnswer) {
	var validCount int64
	for _, passport := range input.passports {
		if validatePassport(passport) {
			validCount++
		}
	}
	return validCount
}

func main() {
	{ // part 1
		runner := &lib.Runner[puzzleInput, puzzleAnswer]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart1,
			LogInput:        false,
		}

		runner.Run("files/example.txt", 2)
		runner.Run("files/input.txt", 239)
	}

	{ // part 2
		runner := &lib.Runner[puzzleInput, puzzleAnswer]{
			InputFileSystem: inputFileSystem,
			InputLoader:     loadInput,
			Solver:          solvePart2,
			LogInput:        false,
		}

		runner.Run("files/example2.txt", 4)
		runner.Run("files/input.txt", 188)
	}
}
