package main

import (
	"adventofcode2015/common"
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

const DAY = "14"
const TEST = false
const RACE_DURATION = 2503

func main() {

	var filePath string
	if TEST {
		filePath = "day" + DAY + "/input_test.txt"
	} else {
		filePath = "day" + DAY + "/input.txt"
	}

	data := common.ReadFileByLine(filePath)
	deers := parseData(data)
	run(deers)

	result1 := getWinner(deers, "distance")
	fmt.Println("Result1 = ", result1)

	result2 := getWinner(deers, "score")
	fmt.Println("Result2 = ", result2)
}

type Deer struct {
	name      string
	speed     int
	endurance int
	rest      int
	state     string
	remaining int
	distance  int
	score     int
}

func parseData(data []string) []Deer {
	deers := []Deer{}
	regex := regexp.MustCompile(`^(?P<name>\w+) can fly (?P<speed>\d+) km/s for (?P<endurance>\d+) seconds, but then must rest for (?P<rest>\d+) seconds.$`)
	for _, d := range data {
		match := regex.FindStringSubmatch(d)
		name := match[slices.Index(regex.SubexpNames(), "name")]
		speed, _ := strconv.Atoi(match[slices.Index(regex.SubexpNames(), "speed")])
		endurance, _ := strconv.Atoi(match[slices.Index(regex.SubexpNames(), "endurance")])
		rest, _ := strconv.Atoi(match[slices.Index(regex.SubexpNames(), "rest")])
		deers = append(deers, Deer{name: name, speed: speed, endurance: endurance, rest: rest, state: "flying", remaining: endurance})
	}
	return deers
}

func run(deers []Deer) {
	for i := 1; i <= RACE_DURATION; i++ {
		for k := range deers {
			deer := &deers[k]
			deer.remaining--
			switch deer.state {
			case "flying":
				deer.distance += deer.speed
				if deer.remaining == 0 {
					deer.state = "resting"
					deer.remaining = deer.rest
				}
			case "resting":
				if deer.remaining == 0 {
					deer.state = "flying"
					deer.remaining = deer.endurance
				}
			}
		}
		cuurent_best := getWinner(deers, "distance")
		for k := range deers {
			deer := &deers[k]
			if deer.distance == cuurent_best {
				deer.score++
			}
		}
	}
}

func getWinner(deers []Deer, criteria string) int {
	best := 0
	for _, d := range deers {
		value := 0
		switch criteria {
		case "distance":
			value = d.distance
		case "score":
			value = d.score
		}
		if value > best {
			best = value
		}
	}
	return best
}
