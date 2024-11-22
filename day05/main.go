package main

import (
	"adventofcode2015/common"
	"fmt"
	"strings"
)

const DAY = "05"
const TEST = false

func main() {

	var filePath string
	if TEST {
		filePath = "day" + DAY + "/input_test.txt"
	} else {
		filePath = "day" + DAY + "/input.txt"
	}

	words := common.ReadFileByLine(filePath)

	result1 := step1(words)
	fmt.Println("Result1 = ", result1)

	result2 := step2(words)
	fmt.Println("Result2 = ", result2)
}

func isNice(word string) bool {
	vowels := 0
	double := false
	previous := ""
	if strings.Contains(word, "ab") || strings.Contains(word, "cd") || strings.Contains(word, "pq") || strings.Contains(word, "xy") {
		return false
	}
	for _, c := range word {
		if strings.Contains("aeiou", string(c)) {
			vowels++
		}
		if string(c) == previous {
			double = true
		}
		previous = string(c)
	}
	if double && vowels > 2 {
		return true
	}
	return false
}

func step1(words []string) int {
	niceWords := 0
	for _, w := range words {
		if isNice(w) {
			niceWords++
		}
	}
	return niceWords
}

func isNice2(word string) bool {
	rule1 := false
	rule2 := false
	previous := ""
	previousPrevious := ""

	for i, c := range word {

		if previous != "" {
			searched := previous + string(c)
			if strings.Contains(word[i+1:], searched) {
				rule1 = true
			}
		}
		if previousPrevious == string(c) {
			rule2 = true
		}
		previousPrevious = previous
		previous = string(c)
	}
	if rule1 && rule2 {
		return true
	}
	return false
}

func step2(words []string) int {
	niceWords := 0
	for _, w := range words {
		if isNice2(w) {
			niceWords++
		}
	}
	return niceWords
}
