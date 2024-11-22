package main

import (
	"adventofcode2015/common"
	"fmt"
)

const DAY = "xx"
const TEST = true

func main() {

	var filePath string
	if TEST {
		filePath = "day" + DAY + "/input_test.txt"
	} else {
		filePath = "day" + DAY + "/input.txt"
	}

	data := common.ReadFileByLine(filePath)

	result1 := step1(data)
	fmt.Println("Result1 = ", result1)

	result2 := step2(data)
	fmt.Println("Result2 = ", result2)
}

func step1(data []string) int {
	result := 0
	return result
}

func step2(data []string) int {
	result := 0
	return result
}
