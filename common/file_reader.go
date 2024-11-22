package common

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFileToInt(path string) []int {
	file, err := os.Open(path)
	Check(err)
	defer file.Close()
	var content []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scannedText, _ := strconv.Atoi(scanner.Text())
		content = append(content, scannedText)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return content
}

func ReadFile(path string) string {
	dat, err := os.ReadFile(path)
	Check(err)
	return string(dat)
}

func ReadFileByLine(path string) []string {
	file, err := os.Open(path)
	Check(err)
	scanner := bufio.NewScanner(file)
	var dat []string
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		dat = append(dat, scanner.Text())
	}
	return dat
}
