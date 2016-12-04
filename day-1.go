package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"math"
	"strconv"
)

func round(input float64) int {
	if input < 0 {
		return int(math.Ceil(input - 0.5))
	}
	return int(math.Floor(input + 0.5))
}

func retrieveInput() []string {
	data, _ := ioutil.ReadFile("./day-1-input.txt")
	return strings.Split(strings.Replace(string(data), " ", "", -1), ",")
}

func absoluteDistance(x, y int) int {
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

type Point struct {
	x, y int
}

func main() {
	input := retrieveInput()

	seen := make(map[Point]bool)
	currentDirection := math.Pi / 2
	x, y := 0, 0

	for _, val := range input {
		direction := val[0]
		steps, _ := strconv.Atoi(val[1:])

		oldX, oldY := x, y

		directionDiff := math.Pi / 2
		if string(direction) == "R" {
			directionDiff = -directionDiff
		}

		currentDirection = math.Mod(currentDirection + directionDiff, math.Pi * 2)

		x += round(float64(steps) * math.Cos(currentDirection))
		y += round(float64(steps) * math.Sin(currentDirection))
		a := math.Copysign(1, x)

		for i = oldX + 1; i <= int(math.Abs(float64(x))); i + 1 {

		}

		if seen[Point{x, y}] {
			fmt.Printf("Revisited point: (%v, %v) at a distance of %v\n", x, y, absoluteDistance(x, y))
		}
		seen[Point{x, y}] = true
	}

	fmt.Println(absoluteDistance(x, y))
}
