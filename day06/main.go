package main

import (
	"adventofcode2015/common"
	"fmt"
	"regexp"
	"strconv"
)

const DAY = "06"
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

type Coord struct {
	x, y int
}

type Command struct {
	action                 string
	upperLeft, bottomRight Coord
}

func getPointsFromSquare(squareUpperLeft Coord, squareBottomRight Coord) []Coord {
	var points []Coord
	for X := squareUpperLeft.x; X <= squareBottomRight.x; X++ {
		for Y := squareUpperLeft.y; Y <= squareBottomRight.y; Y++ {
			points = append(points, Coord{X, Y})
		}
	}
	return points
}

func parseCommands(data []string) []Command {
	regex := regexp.MustCompile(`^([\w\s]+) (\d+),(\d+) through (\d+),(\d+)$`)

	var commands []Command
	for _, command := range data {
		s := regex.FindStringSubmatch(command)
		action := s[1]
		x1, _ := strconv.Atoi(s[2])
		y1, _ := strconv.Atoi(s[3])
		x2, _ := strconv.Atoi(s[4])
		y2, _ := strconv.Atoi(s[5])
		commands = append(commands, Command{action, Coord{x1, y1}, Coord{x2, y2}})
	}
	return commands
}

func step1(data []string) int {
	var board [1000][1000]int
	commands := parseCommands(data)
	for _, c := range commands {
		for _, point := range getPointsFromSquare(c.upperLeft, c.bottomRight) {
			switch c.action {
			case "toggle":
				if board[point.x][point.y] == 0 {
					board[point.x][point.y] = 1
				} else {
					board[point.x][point.y] = 0
				}
			case "turn on":
				board[point.x][point.y] = 1
			case "turn off":
				board[point.x][point.y] = 0
			default:
				panic("Unknow " + c.action)
			}
		}
	}
	count := 0
	for _, row := range board {
		for _, cell := range row {
			if cell == 1 {
				count++
			}
		}
	}
	return count
}

func step2(data []string) int {
	var board [1000][1000]int
	commands := parseCommands(data)
	for _, c := range commands {
		for _, point := range getPointsFromSquare(c.upperLeft, c.bottomRight) {
			switch c.action {
			case "toggle":
				board[point.x][point.y] += 2
			case "turn on":
				board[point.x][point.y]++
			case "turn off":
				if board[point.x][point.y] > 0 {
					board[point.x][point.y]--
				}
			default:
				panic("Unknow " + c.action)
			}
		}
	}
	count := 0
	for _, row := range board {
		for _, cell := range row {
			count += cell
		}
	}
	return count
}
