package common

import (
	"log"
	"time"
)

// handy error checker
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

// get each line as a list of integer
func ParseLineAsBits(line string) []int {
	var bits []int
	for _, bitStr := range line {
		bits = append(bits, convertRuneToInt(bitStr))
	}
	return bits
}

// https://stackoverflow.com/questions/21322173/convert-rune-to-int
func convertRuneToInt(rune rune) int {
	return int(rune - '0')
}
