package main

import (
	"adventofcode2015/common"
	"fmt"
	"strconv"
)

const DAY = "24"
const TEST = false

func main() {

	var filePath string
	if TEST {
		filePath = "day" + DAY + "/input_test.txt"
	} else {
		filePath = "day" + DAY + "/input.txt"
	}

	data := common.ReadFileByLine(filePath)
	parsed := parseData(data)

	result1 := step1(parsed)
	fmt.Println("Result1 = ", result1)

	result2 := step2(parsed)
	fmt.Println("Result2 = ", result2)
}

func parseData(lines []string) []int {
	var weights []int
	for _, line := range lines {
		val, _ := strconv.Atoi(line)
		weights = append(weights, val)
	}
	return weights
}

func getCombinationofN(set []int, depth int, current []int, accum *[][]int) {
	if depth == 0 {
		*accum = append(*accum, current)
	} else {
		for i := 0; i <= len(set)-depth; i++ {
			getCombinationofN(set[i+1:], depth-1, append(current, set[i]), accum)
		}
	}
}

func sum(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func getQE(slice []int) int {
	qe := 1
	for _, v := range slice {
		qe *= v
	}
	return qe
}

func run(weights []int, groupWeight int) int {
	bestQE := 0
	for packetCount := 1; packetCount < len(weights); packetCount++ {
		possibles := [][]int{}
		getCombinationofN(weights, packetCount, []int{}, &possibles)
		foundMatchingCombo := false
		for _, combo := range possibles {
			if sum(combo) == groupWeight {
				foundMatchingCombo = true
				qe := getQE(combo)
				// fmt.Println("Found potential", groupWeight, packetCount, combo, qe)
				if qe < bestQE || bestQE == 0 {
					bestQE = qe
				}
			}

		}
		if foundMatchingCombo {
			break
		}

	}
	return bestQE
}

func step1(weights []int) int {
	groupWeight := sum(weights) / 3
	result := run(weights, groupWeight)
	return result
}

func step2(weights []int) int {
	groupWeight := sum(weights) / 4
	result := run(weights, groupWeight)
	return result
}
