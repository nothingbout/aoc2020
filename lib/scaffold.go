package lib

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"reflect"
	"time"
)

type Benchmark struct {
	name      string
	startTime time.Time
}

func NewBenchmark(name string) *Benchmark {
	return &Benchmark{
		name:      name,
		startTime: time.Now(),
	}
}

func (b *Benchmark) Finish() {
	endTime := time.Now()
	log.Printf("%s took %.3fms", b.name, endTime.Sub(b.startTime).Seconds()*1000)
}

func ReadInputLines(fs fs.FS, inputFilePath string) []string {
	inputFile, err := fs.Open(inputFilePath)
	if err != nil {
		log.Panicf("Input file not found: %s", inputFilePath)
	}
	scanner := bufio.NewScanner(inputFile)
	var inputLines []string
	for scanner.Scan() {
		line := scanner.Text()
		inputLines = append(inputLines, line)
	}
	return inputLines
}

type Runner[InputType, OutputType any] struct {
	InputFileSystem fs.FS
	InputLoader     func(inputLines []string) InputType
	Solver          func(InputType) OutputType
	LogInput        bool
}

func (r *Runner[InputType, OutputType]) Run(inputFilePath string, expectedOutput OutputType) {
	log.Printf("### running %s\n", inputFilePath)
	inputLines := ReadInputLines(r.InputFileSystem, inputFilePath)
	input := r.InputLoader(inputLines)
	if r.LogInput {
		log.Printf("input: %+v\n", input)
	}
	benchmark := NewBenchmark("solver")
	answer := r.Solver(input)
	benchmark.Finish()
	answerOk := "OK"
	if !reflect.DeepEqual(answer, expectedOutput) {
		answerOk = fmt.Sprintf("FAIL, expected %v", expectedOutput)
	}
	log.Printf("answer: %v [%s]\n\n", answer, answerOk)
}
