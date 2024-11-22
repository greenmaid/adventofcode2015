package main

import (
	"adventofcode2015/common"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	filePath := "day04/input.txt"
	fileContent := common.ReadFile(filePath)

	result1 := step1(fileContent)
	fmt.Println("Result1 = ", result1)

	result2 := step2(fileContent)
	fmt.Println("Result2 = ", result2)
}

func step1(data string) int {
	i := 0
	for {
		hash_bytes := md5.Sum([]byte(fmt.Sprintf("%s%d", data, i)))
		hash := hex.EncodeToString(hash_bytes[:])
		if hash[:5] == "00000" {
			return i
		}
		i++
	}
}

func step2(data string) int {
	i := 0
	for {
		hash_bytes := md5.Sum([]byte(fmt.Sprintf("%s%d", data, i)))
		hash := hex.EncodeToString(hash_bytes[:])
		if hash[:6] == "000000" {
			return i
		}
		i++
	}
}
