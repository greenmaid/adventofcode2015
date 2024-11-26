package main

import (
	"adventofcode2015/common"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

const DAY = "12"
const TEST = false

func main() {

	var filePath string
	if TEST {
		filePath = "day" + DAY + "/input_test.txt"
	} else {
		filePath = "day" + DAY + "/input.txt"
	}

	data := common.ReadFile(filePath)

	result1 := step1(data)
	fmt.Println("Result1 = ", result1)

	result2 := step2(data)
	fmt.Println("Result2 = ", result2)
}

func step1(data string) int {
	result := 0
	digits := map[rune]bool{
		[]rune("-")[0]: true,
		[]rune("0")[0]: true,
		[]rune("1")[0]: true,
		[]rune("2")[0]: true,
		[]rune("3")[0]: true,
		[]rune("4")[0]: true,
		[]rune("5")[0]: true,
		[]rune("6")[0]: true,
		[]rune("7")[0]: true,
		[]rune("8")[0]: true,
		[]rune("9")[0]: true,
	}
	buff := []rune{}
	for _, c := range data {
		if _, ok := digits[c]; ok {
			buff = append(buff, c)
			continue
		}
		if len(buff) > 0 {
			figure, err := strconv.Atoi(string(buff))
			if err != nil {
				panic("unable to convert to int")
			}
			result += figure
			buff = []rune{}
		}
	}
	return result
}

func decodeJsonArray(data string) []interface{} {
	var arbitrary_json_array []interface{}
	err := json.Unmarshal([]byte(data), &arbitrary_json_array)
	if err != nil {
		panic(err)
	}
	return arbitrary_json_array
}

func parseJsonArray(jsonObj []interface{}) []interface{} {
	for i, v := range jsonObj {
		if reflect.ValueOf(v).Kind() == reflect.Map {
			jsonObj[i] = parseJsonMap(v.(map[string]interface{}))
		}
		if reflect.ValueOf(v).Kind() == reflect.Slice {
			jsonObj[i] = parseJsonArray(v.([]interface{}))
		}
	}
	return jsonObj
}

func parseJsonMap(jsonObj map[string]interface{}) map[string]interface{} {
	for k, v := range jsonObj {
		if v == "red" {
			return nil
		}
		if reflect.ValueOf(v).Kind() == reflect.Map {
			jsonObj[k] = parseJsonMap(v.(map[string]interface{}))
		}
		if reflect.ValueOf(v).Kind() == reflect.Slice {
			jsonObj[k] = parseJsonArray(v.([]interface{}))
		}
	}
	return jsonObj
}

func step2(data string) int {
	array := parseJsonArray(decodeJsonArray(data))
	result := step1(fmt.Sprintf("%v", array))
	return result
}
