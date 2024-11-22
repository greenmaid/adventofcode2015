package main

import (
	"adventofcode2015/common"
	"fmt"
)

func main() {
	filePath := "day03/input.txt"
	fileContent := common.ReadFile(filePath)

	result1 := step1(fileContent)
	fmt.Println("Result1 = ", result1)

	result2 := step2(fileContent)
	fmt.Println("Result2 = ", result2)
}

type Coord struct {
	x, y int
}

func step1(data string) int {

	currentLocation := Coord{0, 0}
	houses := map[Coord]bool{currentLocation: true}

	for _, d := range data {
		switch string(d) {
		case "^":
			currentLocation.y += 1
		case "v":
			currentLocation.y -= 1
		case "<":
			currentLocation.x -= 1
		case ">":
			currentLocation.x += 1
		}
		houses[currentLocation] = true
	}
	result := len(houses)
	return result
}

func step2(data string) int {
	santaCurrentLocation := Coord{0, 0}
	robotCurrentLocation := Coord{0, 0}
	houses := map[Coord]bool{santaCurrentLocation: true}

	for index, d := range data {

		currentLocation := &robotCurrentLocation
		if index%2 == 0 {
			currentLocation = &santaCurrentLocation
		}

		switch string(d) {
		case "^":
			currentLocation.y += 1
		case "v":
			currentLocation.y -= 1
		case "<":
			currentLocation.x -= 1
		case ">":
			currentLocation.x += 1
		}
		houses[*currentLocation] = true
	}
	result := len(houses)
	return result
}
