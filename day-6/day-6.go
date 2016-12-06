package main

import (
	"fmt"
	"github.com/knappg/advent-of-code-2016/util"
)

func processLine(line string, indexes []map[string]int) {
	for i, val := range line {
		indexes[i][string(val)] += 1
	}
}

func mostCommonChar(index map[string]int) string {
	best := -1
	bestChar := ""
	for key, val := range index {
		if best == -1 || val < best {
			best = val
			bestChar = key
		}
	}
	return bestChar
}

func main() {
	indexes := make([]map[string]int, 8)
	for i := 0; i < 8; i++ {
		indexes[i] = make(map[string]int) 
	}
	input := util.RetrieveInput("./day-6-input.txt")

	for _, val := range input {
		processLine(val, indexes)
	}
	for i := 0; i < 8; i++ {
		fmt.Print(mostCommonChar(indexes[i]))
	}
}
