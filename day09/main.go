package main

import (
	"adventofcode2015/common"
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

const DAY = "09"
const TEST = false

func main() {

	var filePath string
	if TEST {
		filePath = "day" + DAY + "/input_test.txt"
	} else {
		filePath = "day" + DAY + "/input.txt"
	}

	data := common.ReadFileByLine(filePath)
	cities, distances := parseDistance(data)

	result1, result2 := step(cities, distances)
	fmt.Println("Result1 = ", result1)
	fmt.Println("Result2 = ", result2)
}

func parseDistance(data []string) ([]string, map[string]map[string]int) {
	regex := regexp.MustCompile(`^(?P<city1>\w+) to (?P<city2>\w+) = (?P<distance>\d+)$`)

	distances := make(map[string]map[string]int)
	var cities []string
	for _, d := range data {
		match := regex.FindStringSubmatch(d)
		city1 := match[slices.Index(regex.SubexpNames(), "city1")]
		city2 := match[slices.Index(regex.SubexpNames(), "city2")]
		dist := match[slices.Index(regex.SubexpNames(), "distance")]
		distance, _ := strconv.Atoi(dist)
		if distances[city1] == nil {
			distances[city1] = make(map[string]int)
		}
		distances[city1][city2] = distance

		if !slices.Contains(cities, city1) {
			cities = append(cities, city1)
		}
		if !slices.Contains(cities, city2) {
			cities = append(cities, city2)
		}
	}
	return cities, distances
}

func getDistance(city1, city2 string, distances map[string]map[string]int) int {
	if c1, ok := distances[city1]; ok {
		if val, ok := c1[city2]; ok {
			return val
		}
	}
	return distances[city2][city1] // assume exists because reverse not found
}

func trip(current_city string, longest, traveled int, cities, visited []string, distances map[string]map[string]int, possibles *[]int) {

	if len(visited) == len(cities) {
		*possibles = append(*possibles, traveled)
	}

	for _, next_city := range cities {
		if slices.Contains(visited, next_city) {
			continue
		}
		next_traveled := 0
		if current_city != "" {
			next_traveled = traveled + getDistance(current_city, next_city, distances)
		}
		next_visited, _ := common.DeepCopy(visited)
		next_visited = append(next_visited.([]string), next_city)
		trip(next_city, longest, next_traveled, cities, next_visited.([]string), distances, possibles)
	}
}

func step(cities []string, distances map[string]map[string]int) (int, int) {
	visited := []string{}
	possibles := []int{}

	trip("", 0, 0, cities, visited, distances, &possibles)
	return slices.Min(possibles), slices.Max(possibles)
}
