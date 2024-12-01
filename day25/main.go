package main

import (
	"fmt"
)

const DAY = "25"
const START = 20151125
const TARGET_ROW = 2947
const TARGET_COLUMN = 3029

func main() {
	result1 := step1()
	fmt.Println("Result1 = ", result1)
}

func next(value int) int {
	return (value * 252533) % 33554393
}

func step1() int {

	column := 1
	row := 1
	rank := 1

	value := START
	for {
		row--
		column++
		if row == 0 {
			rank++
			row = rank
			column = 1
		}
		value = next(value)
		if column == TARGET_COLUMN && row == TARGET_ROW {
			return value
		}
	}
}
