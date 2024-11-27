package main

import (
	"adventofcode2015/common"
	"fmt"
	"strconv"
)

const DAY = "17"
const TEST = false
const TARGET_CAPACITY = 150

func main() {

	var filePath string
	if TEST {
		filePath = "day" + DAY + "/input_test.txt"
	} else {
		filePath = "day" + DAY + "/input.txt"
	}

	data := common.ReadFileByLine(filePath)
	parsed := parseData(data)

	possibleConbinations, result1 := step1(parsed)
	fmt.Println("Result1 = ", result1)

	result2 := step2(possibleConbinations)
	fmt.Println("Result2 = ", result2)
}

func parseData(data []string) []int {
	result := []int{}
	for _, v := range data {
		int_val, _ := strconv.Atoi(v)
		result = append(result, int_val)

	}
	return result
}

func getTotalCapacity(containers []int) int {
	result := 0
	for _, v := range containers {
		result += v
	}
	return result
}

func addContainer(containerList []int, remainingContainers []int, validContainerCombination *[][]int) {
	for i, k := range remainingContainers {
		nextContainers := containerList
		nextContainers = append(nextContainers, k)
		capacity := getTotalCapacity(nextContainers)
		if capacity == TARGET_CAPACITY {
			*validContainerCombination = append(*validContainerCombination, nextContainers)
		}
		if capacity < TARGET_CAPACITY && i < len(remainingContainers)-1 {
			addContainer(nextContainers, remainingContainers[i+1:], validContainerCombination)
		}
	}

}

func step1(containers []int) ([][]int, int) {
	combinations := [][]int{}
	addContainer([]int{}, containers, &combinations)
	return combinations, len(combinations)
}

func step2(possibleConbinations [][]int) int {
	best := len(possibleConbinations[0])
	count := 0
	for _, k := range possibleConbinations {
		switch {
		case len(k) == best:
			count += 1
		case len(k) < best:
			best = len(k)
			count = 1
		}

	}
	return count
}
