package main

import (
	"adventofcode2015/common"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	filePath := "day02/input.txt"
	fileContent := common.ReadFileByLine(filePath)

	result1 := step1(fileContent)
	fmt.Println("Result1 = ", result1)

	result2 := step2(fileContent)
	fmt.Println("Result2 = ", result2)
}

func step1(data []string) int {
	result := 0
	for _, p := range data {
		l, w, h := getDimension(p)
		result = result + 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)
	}
	return result
}

func step2(data []string) int {
	result := 0
	for _, p := range data {
		l, w, h := getDimension(p)
		result = result + l*w*h + min(l+w, w+h, h+l)*2
	}
	return result
}

func getDimension(pack string) (int, int, int) {
	dimensions := strings.Split(pack, "x")
	l, _ := strconv.Atoi(dimensions[0])
	w, _ := strconv.Atoi(dimensions[1])
	h, _ := strconv.Atoi(dimensions[2])
	return l, w, h

}
