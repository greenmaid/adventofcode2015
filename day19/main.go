package main

import (
	"adventofcode2015/common"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/exp/rand"
)

const DAY = "19"
const TEST = false

func main() {

	var filePath string
	if TEST {
		filePath = "day" + DAY + "/input_test.txt"
	} else {
		filePath = "day" + DAY + "/input.txt"
	}

	data := common.ReadFile(filePath)
	molecule, mutationTable, reverseMutationTable := parseData(data)

	result1, _ := step1(molecule, mutationTable)
	fmt.Println("Result1 = ", result1)

	result2, tries := step2(molecule, reverseMutationTable)
	fmt.Println("Result2 = ", result2, "after", tries, "tries")
}

func parseData(data string) (string, map[string][]string, map[string][]string) {
	splited := strings.Split(data, "\n\n")
	molecule := splited[1]

	mutationTable := make(map[string][]string)
	reverseMutationTable := make(map[string][]string)
	regex := regexp.MustCompile(`^([\w]+) => ([\w]+)$`)
	for _, l := range strings.Split(splited[0], "\n") {
		match := regex.FindStringSubmatch(l)
		mutationTable[match[1]] = append(mutationTable[match[1]], match[2])
		reverseMutationTable[match[2]] = append(reverseMutationTable[match[2]], match[1])
	}
	return molecule, mutationTable, reverseMutationTable
}

func step1(molecule string, mutationTable map[string][]string) (int, map[string]bool) {
	possibleMutations := make(map[string]bool)
	for initial, possibleNews := range mutationTable {
		indexes := common.FindAllSubstringIndexes(molecule, initial)
		for _, idx := range indexes {
			for _, new := range possibleNews {
				mutation := fmt.Sprintf("%s%s%s", molecule[:idx], new[:], molecule[idx+len(initial):])
				possibleMutations[mutation] = true
			}
		}
	}
	// fmt.Println(possibleMutations)
	return len(possibleMutations), possibleMutations
}

func step2(molecule string, mutationTable map[string][]string) (int, int) {

	tries := 1
	for {
		step := 0
		newMolecule := molecule
		for {
			step++
			_, possibleMutations := step1(newMolecule, mutationTable)
			if len(possibleMutations) == 0 {
				break
			}

			// pick a random mutation from possible mutations
			k := rand.Intn(len(possibleMutations))
			for m := range possibleMutations {
				if k == 0 {
					newMolecule = m
					break
				}
				k--
			}

			if newMolecule == "e" {
				return step, tries
			}
		}
		tries++
	}
}
