package common

import (
	"encoding/json"
	"log"
	"reflect"
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

// classical map function
func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

// easily deepcopy struct (only exported Fields !)
// https://stackoverflow.com/questions/50269322/how-to-copy-struct-and-dereference-all-pointers
func DeepCopy(v interface{}) (interface{}, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	vptr := reflect.New(reflect.TypeOf(v))
	err = json.Unmarshal(data, vptr.Interface())
	if err != nil {
		return nil, err
	}
	return vptr.Elem().Interface(), err
}
