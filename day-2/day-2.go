package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"math"
)

type Point struct {
	x, y int
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func shouldMakeMove(newX int, newY int, keypad [][]string) bool {
	return newX >= 0 && newX < len(keypad[0]) && newY >= 0 && newY < len(keypad) && keypad[newY][newX] != ""
}

func retrieveInput() []string {
	data, err := ioutil.ReadFile("./day-2-input.txt")
	if err != nil {
		fmt.Printf("Err: %v", err)
	}
	return strings.Split(string(data),"\n")
}

func processMove(move string, curPoint Point, keypad [][]string) Point {
	// fmt.Printf("Processing move: %v\n", move)
	// fmt.Printf("Previous point: (%v, %v)\n", curPoint.x, curPoint.y)
	newX, newY := curPoint.x, curPoint.y
	width, height := len(keypad[0]), len(keypad)
	switch move {
	case "U":
		// fmt.Println("Going up")
		newY = max(curPoint.y - 1, 0)
		// return Point{curPoint.x, max(curPoint.y - 1, 0)}
	case "D":
		// fmt.Println("Going down")
		newY = min(curPoint.y + 1, height - 1)
		// return Point{curPoint.x, min(curPoint.y + 1, HEIGHT - 1)}
	case "L":
		// fmt.Println("Going left")
		newX = max(curPoint.x - 1, 0)
		// return Point{max(curPoint.x - 1, 0), curPoint.y}
	case "R":
		// fmt.Println("Going right")
		newX = min(curPoint.x + 1, width - 1)
		// return Point{min(curPoint.x + 1, WIDTH - 1), curPoint.y}
	}
	if shouldMakeMove(newX, newY, keypad) {
		return Point{newX, newY}
	}
	// fmt.Println("BIG ERRROR IMPOSSIBLE")
	return curPoint
}

func processLine(line string, curPoint Point) Point {
	keypad := [][]string{
		[]string{"",  "",  "1",  "",  ""},
		[]string{"",  "2", "3", "4",  ""},
		[]string{"5", "6", "7", "8", "9"},
		[]string{"",  "A", "B", "C",  ""},
		[]string{"",  "",  "D",  "",  ""},
	}

	for _, c := range line {
		curPoint = processMove(string(c), curPoint, keypad)
		// fmt.Printf("Current position is: (%v, %v)\n", curPoint.x, curPoint.x)
	}
	// fmt.Printf("Final position is: (%v, %v)\n", curPoint.x, curPoint.x)
	fmt.Printf("Letter in the code: %v\n", keypad[curPoint.y][curPoint.x])
	return curPoint
}

func main() {
	input := retrieveInput()
	

	// Starts at 5
	curPoint := Point{1, 1}
	for _, val := range input {
		curPoint = processLine(val, curPoint)
	}
}
