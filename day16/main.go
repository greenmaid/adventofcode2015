package main

import (
	"adventofcode2015/common"
	"fmt"
	"regexp"
	"strconv"
)

const DAY = "16"
const TEST = false

func main() {

	var filePath string
	if TEST {
		filePath = "day" + DAY + "/input_test.txt"
	} else {
		filePath = "day" + DAY + "/input.txt"
	}

	mysteryAuntProperties := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	data := common.ReadFileByLine(filePath)
	aunts := parsedata(data)

	result1 := step1(aunts, mysteryAuntProperties)
	fmt.Println("Result1 = ", result1)

	result2 := step2(aunts, mysteryAuntProperties)
	fmt.Println("Result2 = ", result2)
}

type Aunt map[string]int

func parsedata(lines []string) []Aunt {
	aunts := []Aunt{}
	nameRegex := regexp.MustCompile(`^Sue ([\d]+):`)
	propertyRegex := regexp.MustCompile(`([\w]+): ([\d]+)`)

	for _, line := range lines {
		nameMatch := nameRegex.FindStringSubmatch(line)
		name, _ := strconv.Atoi(nameMatch[1])
		aunt := Aunt{}
		aunt["name"] = name
		propertyMatch := propertyRegex.FindAllStringSubmatch(line, -1)
		for _, p := range propertyMatch {
			value, _ := strconv.Atoi(p[2])
			aunt[p[1]] = value
		}
		aunts = append(aunts, aunt)
	}
	return aunts

}

func step1(aunts []Aunt, mysteryAuntProperties map[string]int) int {
	potentialAunts := aunts
	for prop, expected := range mysteryAuntProperties {
		validAunts := []Aunt{}
		for _, aunt := range potentialAunts {
			val, ok := aunt[prop]
			if ok && val != expected {
				continue
			}
			validAunts = append(validAunts, aunt)
		}
		potentialAunts = validAunts
	}
	return potentialAunts[0]["name"]
}

func step2(aunts []Aunt, mysteryAuntProperties map[string]int) int {
	potentialAunts := aunts
	for prop, expected := range mysteryAuntProperties {
		validAunts := []Aunt{}
		for _, aunt := range potentialAunts {
			val, ok := aunt[prop]
			if !ok {
				validAunts = append(validAunts, aunt)
				continue
			}
			switch {
			case prop == "cats" || prop == "trees":
				if val <= expected {
					// fmt.Println(aunt["name"], "ejected because", prop, "=", val, "expected more than", expected)
					continue
				}
			case prop == "pomeranians" || prop == "goldfish":
				if val >= expected {
					// fmt.Println(aunt["name"], "ejected because", prop, "=", val, "expected less than", expected)
					continue
				}
			default:
				if val != expected {
					// fmt.Println(aunt["name"], "ejected because", prop, "=", val, "expected", expected)
					continue
				}
			}
			validAunts = append(validAunts, aunt)
		}
		potentialAunts = validAunts
	}
	fmt.Println(potentialAunts)
	return potentialAunts[0]["name"]
}
