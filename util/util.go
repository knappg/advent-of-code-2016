package util

import (
	"fmt"
	"strings"
	"math"
	"io/ioutil"
)

func RetrieveInput(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Err: %v", err)
	}
	return strings.Split(string(data),"\n")
}

func Max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func Min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func Round(input float64) int {
	if input < 0 {
		return int(math.Ceil(input - 0.5))
	}
	return int(math.Floor(input + 0.5))
}