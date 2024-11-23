package main

import (
	"adventofcode2015/common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const DAY = "07"
const TEST = false

func main() {

	var filePath string
	if TEST {
		filePath = "day" + DAY + "/input_test.txt"
	} else {
		filePath = "day" + DAY + "/input.txt"
	}

	data := common.ReadFileByLine(filePath)
	instructions := parseData(data)

	result1 := step1("a", instructions)
	fmt.Println("Result1 = ", result1)

	result2 := step2("a", instructions, result1)
	fmt.Println("Result2 = ", result2)
}

type Instruction struct {
	operation      string
	input1, input2 string
}

func parseData(data []string) map[string]Instruction {

	regex := regexp.MustCompile(`^([\w\d]+) ([ANDORLSHIFT]+) ([\w\d]+)$`)
	regexNOT := regexp.MustCompile(`^NOT ([\w\d]+)$`)
	regexDIRECT := regexp.MustCompile(`([\w\d]+)$`)

	instructions := make(map[string]Instruction)

	for _, line := range data {
		splittedLine := strings.Split(line, " -> ")
		wire := splittedLine[1]
		var instruction Instruction
		if parsed := regex.FindStringSubmatch(splittedLine[0]); len(parsed) != 0 {
			instruction = Instruction{parsed[2], parsed[1], parsed[3]}
		} else if parsed := regexNOT.FindStringSubmatch(splittedLine[0]); len(parsed) != 0 {
			instruction = Instruction{"NOT", parsed[1], ""}
		} else if parsed := regexDIRECT.FindStringSubmatch(splittedLine[0]); len(parsed) != 0 {
			instruction = Instruction{"DIRECT", parsed[1], ""}
		} else {
			panic("unparsed " + splittedLine[0])
		}
		instructions[wire] = instruction
	}
	return instructions
}

func evaluate(wire string, instructions map[string]Instruction, solved_wires map[string]uint16) uint16 {

	if val, ok := solved_wires[wire]; ok {
		return val
	}
	if val, err := strconv.ParseInt(wire, 10, 16); err == nil {
		return uint16(val)
	}
	var result uint16
	instruction := instructions[wire]
	switch instruction.operation {
	case "DIRECT":
		if val, err := strconv.ParseInt(instruction.input1, 10, 16); err == nil {
			return uint16(val)
		} else {
			result = evaluate(instruction.input1, instructions, solved_wires)
		}
	case "NOT":
		result = ^evaluate(instruction.input1, instructions, solved_wires)
	case "AND":
		result = evaluate(instruction.input1, instructions, solved_wires) & evaluate(instruction.input2, instructions, solved_wires)
	case "OR":
		result = evaluate(instruction.input1, instructions, solved_wires) | evaluate(instruction.input2, instructions, solved_wires)
	case "LSHIFT":
		return evaluate(instruction.input1, instructions, solved_wires) << evaluate(instruction.input2, instructions, solved_wires)
	case "RSHIFT":
		result = evaluate(instruction.input1, instructions, solved_wires) >> evaluate(instruction.input2, instructions, solved_wires)
	default:
		panic("unhandled " + instruction.operation)
	}
	solved_wires[wire] = result
	return result
}

func step1(wire string, instructions map[string]Instruction) uint16 {
	solved_wires := make(map[string]uint16)
	return evaluate(wire, instructions, solved_wires)
}

func step2(wire string, instructions map[string]Instruction, result1 uint16) uint16 {
	instructions["b"] = Instruction{"DIRECT", fmt.Sprint(result1), ""}
	solved_wires := make(map[string]uint16)
	return evaluate(wire, instructions, solved_wires)
}
