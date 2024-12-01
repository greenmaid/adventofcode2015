package main

import (
	"adventofcode2015/common"
	"fmt"
	"regexp"
	"strconv"
)

const DAY = "23"
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

func parseData(data []string) [][3]string {
	commands := [][3]string{}
	regex1 := regexp.MustCompile(`^(\w+) ([-+\w\d]+)$`)
	regex2 := regexp.MustCompile(`^(\w+) ([\w\d]+), ([-+\w\d]+)$`)
	for _, line := range data {
		var command [3]string
		match1 := regex1.FindStringSubmatch(line)
		if len(match1) > 0 {
			command = [3]string{match1[1], match1[2], ""}
		} else {
			match2 := regex2.FindStringSubmatch(line)
			command = [3]string{match2[1], match2[2], match2[3]}
		}
		commands = append(commands, command)
	}
	return commands
}

func runProgram(registers map[string]int, commands [][3]string) map[string]int {
	idx := 0
	for idx < len(commands) {
		command := commands[idx]
		switch command[0] {
		case "hlf":
			r := command[1]
			registers[r] = registers[r] / 2
			idx++
		case "tpl":
			r := command[1]
			registers[r] = registers[r] * 3
			idx++
		case "inc":
			r := command[1]
			registers[r] = registers[r] + 1
			idx++
		case "jmp":
			offset, err := strconv.Atoi(command[1])
			if err != nil {
				offset = registers[command[1]]
			}
			idx += offset
		case "jie":
			r := command[1]
			if registers[r]%2 == 0 {
				offset, err := strconv.Atoi(command[2])
				if err != nil {
					offset = registers[command[2]]
				}
				idx += offset
			} else {
				idx++
			}
		case "jio":
			r := command[1]
			if registers[r] == 1 {
				offset, err := strconv.Atoi(command[2])
				if err != nil {
					offset = registers[command[2]]
				}
				idx += offset
			} else {
				idx++
			}
		}

	}

	return registers
}

func step1(commands [][3]string) int {
	registers := map[string]int{
		"a": 0,
		"b": 0,
	}
	registers = runProgram(registers, commands)
	return registers["b"]
}

func step2(commands [][3]string) int {
	registers := map[string]int{
		"a": 1,
		"b": 0,
	}
	registers = runProgram(registers, commands)
	return registers["b"]
}
