package main

import (
	"fmt"
	"sort"
	"strconv"
	"github.com/knappg/advent-of-code-2016/util"
	"github.com/knappg/advent-of-code-2016/day-4/wordbank"
	"regexp"
)

// Guesses
// 185371 is correct

func isValidLine(matches []string) bool {
	wordBank := wordbank.NewWordBank()
	for _, val := range matches[0] {
		if string(val) != "-" {
			wordBank.MarkLetter(string(val))
		}
	}
	sort.Sort(wordBank)
	for idx, val := range matches[2] {
		letterIdx := wordBank.LetterIndex(string(val))
		if letterIdx == -1 || letterIdx != idx {
			return false
		}
	}
	// fmt.Println(wordBank)
	// fmt.Println(matches[2])
	return true
}

func main() {
	input := util.RetrieveInput("./day-4-input.txt")

	// (\w+-)+(\d+)\[(\w+)\]
	lineRegex := regexp.MustCompile(`([\w-]+)-(\d+)\[(\w+)\]`)

	sum := 0
	for _, val := range input {
		matches := lineRegex.FindAllStringSubmatch(val, -1)[0][1:]
		if isValidLine(matches) {
			val, _ := strconv.Atoi(matches[1])
			// fmt.Printf("Valid line, with id: %v\n", val)
			sum += val
		}
	}
	fmt.Println(sum)
}
