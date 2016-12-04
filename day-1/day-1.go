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
	data, err := ioutil.ReadFile("./day-1-input.txt")
	if err != nil {
		fmt.Printf("Err: %v", err)
	}
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

	// visitIndex := 1
	seen[Point{0, 0}] = true
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

		xDiff := x - oldX
		yDiff := y - oldY

		if xDiff != 0 {
			a := int(math.Copysign(1, float64(xDiff)))
			for i := oldX + a; i != x; i += a {
				if seen[Point{i, y}] {
					fmt.Printf("Revisited point: (%v, %v) at a distance of %v\n", i, y, absoluteDistance(i, y))
				}
				seen[Point{i, y}] = true
			}
			if seen[Point{x, y}] {
				fmt.Printf("Revisited point: (%v, %v) at a distance of %v\n", x, y, absoluteDistance(x, y))
			}
			seen[Point{x, y}] = true
		}

		if yDiff != 0 {
			a := int(math.Copysign(1, float64(yDiff)))
			for i := oldY + a; i != y; i += a {
				if seen[Point{x, i}] {
					fmt.Printf("Revisited point: (%v, %v) at a distance of %v\n", x, i, absoluteDistance(x, i))
				}
				seen[Point{x, i}] = true
			}
			if seen[Point{x, y}] {
				fmt.Printf("Revisited point: (%v, %v) at a distance of %v\n", x, y, absoluteDistance(x, y))
			}
			seen[Point{x, y}] = true
		}

	}

	fmt.Println(absoluteDistance(x, y))
}
