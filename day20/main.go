package main

import (
	"fmt"
)

const DAY = "20"
const TEST = true
const TARGET = 36000000

func main() {

	result1 := run("step1")
	fmt.Println("Result1 = ", result1)

	result2 := run("step2")
	fmt.Println("Result2 = ", result2)
}

func isPrime(n int, primes []int) bool {
	for _, p := range primes {
		if n%p == 0 {
			return false
		}
	}
	return true
}

func getCombinationsOfN(elements []int, n int) map[int]bool {
	return getCombosHelper(elements, n, 0, 1, make(map[int]bool))
}

func getCombosHelper(set []int, depth int, start int, prefix int, accum map[int]bool) map[int]bool {
	if depth == 0 {
		accum[prefix] = true
	} else {
		for i := start; i <= len(set)-depth; i++ {
			accum = getCombosHelper(set, depth-1, i+1, prefix*set[i], accum)
		}
	}
	return accum
}

// Count how many gift are shipped in an house with house as its number decomposed into prime numbers
// example house #12 will be calculated with countHouseGift([]int{2,2,3})
func countHouseGifts(house []int, step string) int {
	allDivisors := make(map[int]bool)
	for n := 0; n <= len(house); n++ {
		for k, v := range getCombinationsOfN(house, n) {
			switch step {
			case "step2":
				if int(getHouseNumber(house)/k) <= 50 {
					allDivisors[k] = v
				}
			default:
				allDivisors[k] = v
			}
		}
	}
	result := 0
	for k := range allDivisors {
		result += k
	}
	var giftsPerHouse int
	switch step {
	case "step2":
		giftsPerHouse = 11
	default:
		giftsPerHouse = 10
	}
	return result * giftsPerHouse
}

func getHouseNumber(houseSlice []int) int {
	result := 1
	for _, v := range houseSlice {
		result *= v
	}
	return result
}

func findBestHouse(current []int, remainingPrimes []int, best int, step string) int {
	for i, p := range remainingPrimes {
		next := []int{}
		next = append(next, current...)
		next = append(next, p)
		if countHouseGifts(next, step) >= TARGET {
			nextNum := getHouseNumber(next)
			if nextNum < best {
				best = nextNum
			}
			continue
		}
		best = findBestHouse(next, remainingPrimes[i:], best, step)
	}
	return best
}

func run(step string) int {
	best := TARGET
	elf := 1
	primes := []int{}
	for elf < 50 {
		elf++
		if isPrime(elf, primes) {
			primes = append(primes, elf)
		}
	}
	best = findBestHouse([]int{}, primes, best, step)
	return best
}
