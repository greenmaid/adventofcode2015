package main

import (
	"adventofcode2015/common"
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

const DAY = "13"
const TEST = false

func main() {

	var filePath string
	if TEST {
		filePath = "day" + DAY + "/input_test.txt"
	} else {
		filePath = "day" + DAY + "/input.txt"
	}

	data := common.ReadFileByLine(filePath)
	people := parseData(data)

	result1 := step1(people)
	fmt.Println("Result1 = ", result1)

	result2 := step2(people)
	fmt.Println("Result2 = ", result2)
}

func parseData(data []string) map[string]map[string]int {
	people := make(map[string]map[string]int)
	regex := regexp.MustCompile(`^(?P<person1>\w+) would (?P<gainOrLose>\w+) (?P<quantity>\d+) happiness units by sitting next to (?P<person2>\w+).$`)
	for _, d := range data {
		match := regex.FindStringSubmatch(d)
		person1 := match[slices.Index(regex.SubexpNames(), "person1")]
		person2 := match[slices.Index(regex.SubexpNames(), "person2")]
		gainOrLose := match[slices.Index(regex.SubexpNames(), "gainOrLose")]
		quantity_str := match[slices.Index(regex.SubexpNames(), "quantity")]
		quantity, _ := strconv.Atoi(quantity_str)
		if _, ok := people[person1]; !ok {
			people[person1] = make(map[string]int)
		}
		if gainOrLose == "gain" {
			people[person1][person2] = quantity
		}
		if gainOrLose == "lose" {
			people[person1][person2] = quantity * -1
		}
	}
	return people
}

func getHappiness(placement []interface{}, people map[string]map[string]int) int {
	result := 0
	for i, p := range placement {
		if i == len(placement)-1 {
			result += people[p.(string)][placement[0].(string)]
			result += people[placement[0].(string)][p.(string)]
		} else {
			result += people[p.(string)][placement[i+1].(string)]
			result += people[placement[i+1].(string)][p.(string)]
		}
	}
	return result
}

func step1(people map[string]map[string]int) int {
	var persons []interface{}
	for k := range people {
		persons = append(persons, k)
	}
	result := 0
	for _, p := range common.Permutations(persons) {
		totalHappiness := getHappiness(p, people)
		if totalHappiness > result {
			result = totalHappiness
		}
	}
	return result
}

func step2(people map[string]map[string]int) int {
	var persons []interface{}
	people["me"] = make(map[string]int)
	for k := range people {
		people[k]["me"] = 0
		people["me"][k] = 0
		persons = append(persons, k)
	}
	result := 0
	for _, p := range common.Permutations(persons) {
		totalHappiness := getHappiness(p, people)
		if totalHappiness > result {
			result = totalHappiness
		}
	}
	return result
}
