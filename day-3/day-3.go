package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"sort"
)

func retrieveInput() []string {
	data, err := ioutil.ReadFile("./day-3-input.txt")
	if err != nil {
		fmt.Printf("Err: %v", err)
	}
	return strings.Split(string(data),"\n")
}

func convertStringToInt(c string) int {
	val, _ := strconv.Atoi(c)
	return val
}

func groupIsValid(line1, line2, line3 string) int {
	lineInts1 := strings.Fields(line1)
	lineInts2 := strings.Fields(line2)
	lineInts3 := strings.Fields(line3)

	totalInvalidLines := 0
	for i := 0; i < 3; i++ {
		triangleValid := intsAreValid([]int{ convertStringToInt(lineInts1[i]), convertStringToInt(lineInts2[i]), convertStringToInt(lineInts3[i])})
		if !triangleValid {
			totalInvalidLines++
		}
	}
	return totalInvalidLines
}

func lineIsValid(line string) bool {
	line = strings.TrimSpace(line)
	numStrings := strings.Fields(line)
	if len(numStrings) != 3 {
		fmt.Printf("Error while parsing line: %v\n", line)
	}
	ints := make([]int, 0)
	for _, val := range numStrings {
		intVal, _ := strconv.Atoi(string(val))
		ints = append(ints, intVal)
	}
	return intsAreValid(ints)
}

func intsAreValid(ints []int) bool {
	sort.Ints(ints)
	// fmt.Printf("Here are our ints: %v\n", ints)
	if ints[0] + ints[1] <= ints[2] {
		fmt.Printf("Impossible line: %v\n", ints)
		return false
	}
	return true
}

func main() {
	input := retrieveInput()

	invalidLines := 0
	for i := 0; i < len(input) / 3; i++ {
		invalidLines += groupIsValid(input[3 * i], input[3 * i + 1], input[3 * i + 2])
	}

	fmt.Printf("Total number of lines: %v\n", len(input))
	fmt.Printf("Found this many impossible lines: %v\n", invalidLines)
}
