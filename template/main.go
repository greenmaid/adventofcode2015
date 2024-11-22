package main

import (
	"adventofcode2015/common"
	"fmt"
)

func main() {
	filePath := "dayxx/input_test.txt"
	fileContent := common.ReadFileByLine(filePath)

	result1 := step1(fileContent)
	fmt.Println("Result1 = ", result1)

	result2 := step2(fileContent)
	fmt.Println("Result2 = ", result2)
}

func step1(data []string) int {
	fmt.Println(data)
	result := 0
	return result
}

func step2(data []string) int {
	result := 0
	return result
}
