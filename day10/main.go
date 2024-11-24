package main

import (
	"fmt"
)

const DAY = "10"
const START = "1113122113"

func main() {

	result1 := step1()
	fmt.Println("Result1 = ", result1)

	result2 := step2()
	fmt.Println("Result2 = ", result2)
}

func step1() int {
	number := START
	const ROUNDS = 40
	for i := 1; i <= ROUNDS; i++ {
		number = next(number)
	}
	return len(number)
}

func step2() int {
	number := START
	const ROUNDS = 50
	for i := 1; i <= ROUNDS; i++ {
		number = next(number)
	}
	return len(number)
}

func next(number string) string {
	runes := []rune(number)
	result := []rune{}
	index := 0
	for index < len(runes) {
		char := runes[index]
		count := 0
		for index+count < len(runes) && runes[index+count] == char {
			count++
		}
		count_symbol := count + 48 // ASCII 0=48, 1=49, 2=50....
		result = append(result, rune(count_symbol), char)
		index += count
	}
	return string(result)
}

// func next(number string) string {
// 	result := []rune{}
// 	buff := []rune{}
// 	for _, c := range number {
// 		if len(buff) == 0 {
// 			buff = []rune{c}
// 			continue
// 		}
// 		if c == buff[0] {
// 			buff = append(buff, c)
// 			continue
// 		}
// 		result = append(result, rune(len(buff)+48), buff[0])
// 		buff = []rune{c}
// 	}
// 	if len(buff) > 0 {
// 		result = append(result, rune(len(buff)+48), buff[0])
// 	}
// 	return string(result)
// }
