package main

import (
	"fmt"
)

const DAY = "11"
const INTIAL_PASS = "vzbxkghb"

func main() {

	result1 := run(INTIAL_PASS)
	fmt.Println("Result1 = ", result1)

	result2 := run(result1)
	fmt.Println("Result2 = ", result2)
}

func run(password string) string {
	for {
		password = next(password)
		if isValid(password) {
			return password
		}
	}
}

func isValid(password string) bool {
	rule1 := false
	rule2 := 0
	previous := [2]rune{}
	for i, c := range password {
		if string(c) == "i" || string(c) == "o" || string(c) == "l" {
			return false
		}

		if i >= 3 && c == previous[1]+1 && previous[1] == previous[0]+1 {
			rule1 = true
		}
		if i >= 2 && c == previous[1] && previous[0] != previous[1] {
			rule2++
		}
		previous[0] = previous[1]
		previous[1] = c
	}
	if rule1 && rule2 > 1 {
		return true
	}
	return false
}

func next(password string) string {
	runes := []rune(password)
	index := len(runes) - 1
	for {
		char := runes[index] + 1
		if string(char) == "i" || string(char) == "o" || string(char) == "l" {
			char++
		}
		if char != 123 { // after "z"
			runes[index] = char
			return string(runes)
		}
		runes[index] = 97 // "a"
		index--
	}

}
