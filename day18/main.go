package main

import (
	"adventofcode2015/common"
	"fmt"
)

const DAY = "18"
const TEST = false
const STEPS = 100

func main() {

	var filePath string
	if TEST {
		filePath = "day" + DAY + "/input_test.txt"
	} else {
		filePath = "day" + DAY + "/input.txt"
	}

	grid := common.ReadFileToGrid(filePath)

	result1 := run(grid, "step1")
	fmt.Println("Result1 = ", result1)

	result2 := run(grid, "step2")
	fmt.Println("Result2 = ", result2)
}

func deepCopyGrid(grid [][]rune) [][]rune {
	var GridCopy [][]rune
	for _, row := range grid {
		rowCopy := make([]rune, len(row))
		_ = copy(rowCopy, row)
		GridCopy = append(GridCopy, rowCopy)
	}
	return GridCopy
}

func countLightsOn(grid [][]rune) int {
	dimension := len(grid) // assuming a square grid
	count := 0
	for y := 0; y < dimension; y++ {
		for x := 0; x < dimension; x++ {
			if grid[x][y] == rune('#') {
				count++
			}
		}
	}
	return count
}

func run(grid [][]rune, step string) int {
	dimension := len(grid) // assuming a square grid
	currentGrid := deepCopyGrid(grid)

	if step == "step2" {
		currentGrid[0][0] = rune('#')
		currentGrid[0][dimension-1] = rune('#')
		currentGrid[dimension-1][dimension-1] = rune('#')
		currentGrid[dimension-1][0] = rune('#')
	}

	for i := 0; i < STEPS; i++ {
		newGrid := deepCopyGrid(currentGrid)
		for y := 0; y < dimension; y++ {
			for x := 0; x < dimension; x++ {
				count := 0
				if x > 0 && y > 0 && string(currentGrid[x-1][y-1]) == "#" {
					count++
				}
				if y > 0 && string(currentGrid[x][y-1]) == "#" {
					count++
				}
				if x < dimension-1 && y > 0 && string(currentGrid[x+1][y-1]) == "#" {
					count++
				}
				if x > 0 && string(currentGrid[x-1][y]) == "#" {
					count++
				}
				if x < dimension-1 && string(currentGrid[x+1][y]) == "#" {
					count++
				}
				if x > 0 && y < dimension-1 && string(currentGrid[x-1][y+1]) == "#" {
					count++
				}
				if y < dimension-1 && string(currentGrid[x][y+1]) == "#" {
					count++
				}
				if x < dimension-1 && y < dimension-1 && string(currentGrid[x+1][y+1]) == "#" {
					count++
				}

				switch currentGrid[x][y] {
				case rune('#'):
					if count != 2 && count != 3 {
						newGrid[x][y] = rune('.')
					}
				case rune('.'):
					if count == 3 {
						newGrid[x][y] = rune('#')
					}
				}

				if step == "step2" {
					newGrid[0][0] = rune('#')
					newGrid[0][dimension-1] = rune('#')
					newGrid[dimension-1][dimension-1] = rune('#')
					newGrid[dimension-1][0] = rune('#')
				}
			}
		}

		currentGrid = newGrid
	}
	// common.DisplayGrid(currentGrid)
	result := countLightsOn(currentGrid)
	return result
}
