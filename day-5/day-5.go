package main

import (
	"fmt"
	"strconv"
	"crypto/md5"
)

const ROOM string = "ojvtpuvg"

func isPassword(hash string) bool {
	for i := 0; i < 5; i++ {
		if string(hash[i]) != "0" {
			return false
		}
	}
	return true
}

func main() {

	pass := make([]string, 8)

	foundNums := 0
	for i := 0; ; i++ {
		stringToHash := ROOM + strconv.Itoa(i)
		hashed := fmt.Sprintf("%x", md5.Sum([]byte(stringToHash)))
		if isPassword(hashed) {
			index, err := strconv.Atoi(string(hashed[5]))
			char := string(hashed[6])
			if err == nil && index < 8 && pass[index] == "" {
				fmt.Printf("Found one: %v\n", char)
				foundNums++
				pass[index] = char
			}
		}
		if foundNums == 8 {
			break
		}
	}
	fmt.Printf("Final password: %v\n", pass)
}
