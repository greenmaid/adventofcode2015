package main

import (
	"adventofcode2015/common"
	"fmt"
	"strings"
)

const DAY = "08"
const TEST = false

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

func countChars(line string) int {
	return len(line)
}

func countEvaluatedChars(line string) int {
	evaluated := line[1 : len(line)-1]
	evaluated = strings.Replace(evaluated, `\\`, `/`, -1)
	evaluated = strings.Replace(evaluated, `\"`, `/`, -1)
	hex_count := strings.Count(evaluated, `\x`)
	return len(evaluated) - 3*hex_count
}

// For archive, with system call
// func countEvaluateWithsystemCmd(line string) int {
// 	out, _ := exec.Command("/usr/bin/printf", line[1:len(line)-1]).Output()
// 	return len(out)
// }

func countEncoded(line string) int {
	encoded := 2
	for _, c := range line {
		if string(c) == `"` {
			encoded += 2
		} else if string(c) == `\` {
			encoded += 2
		} else {
			encoded++
		}
	}
	return encoded
}

func step1(data []string) int {
	result := 0
	for _, line := range data {
		// result += countChars(line) - countEvaluateWithsystemCmd(line)
		result += countChars(line) - countEvaluatedChars(line)
	}
	return result
}

func step2(data []string) int {
	result := 0
	for _, line := range data {
		result += countEncoded(line) - countChars(line)
	}
	return result
}
