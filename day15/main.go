package main

import (
	"adventofcode2015/common"
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

const DAY = "15"
const TEST = false

func main() {

	var filePath string
	if TEST {
		filePath = "day" + DAY + "/input_test.txt"
	} else {
		filePath = "day" + DAY + "/input.txt"
	}

	data := common.ReadFileByLine(filePath)
	ingredients := parseData(data)

	recipe1, result1 := step1(ingredients)
	fmt.Println("Result1 = ", result1)

	_, result2 := step2(recipe1, ingredients)
	fmt.Println("Result2 = ", result2)
}

type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

type Recipe map[Ingredient]int

func (R Recipe) getScore() int {
	total := Ingredient{"total", 0, 0, 0, 0, 0}
	for ingredient, count := range R {
		if count <= 0 && ingredient.name != "calories" {
			return 0
		}
		total.capacity += ingredient.capacity * count
		total.durability += ingredient.durability * count
		total.flavor += ingredient.flavor * count
		total.texture += ingredient.texture * count
		total.calories += ingredient.calories * count
	}
	return total.capacity * total.durability * total.flavor * total.texture
}

func (R Recipe) count() int {
	sum := 0
	for _, v := range R {
		sum += v
	}
	return sum
}

func (R Recipe) getCalories() int {
	sum := 0
	for ing, v := range R {
		sum += ing.calories * v
	}
	return sum
}

func (R Recipe) replace(i1 Ingredient, i2 Ingredient) {
	if R[i1] > 0 {
		R[i1]--
		R[i2]++
	}
}

func parseData(data []string) map[string]Ingredient {
	ingredients := make(map[string]Ingredient)
	regex := regexp.MustCompile(`(?P<name>\w+): capacity (?P<capacity>[-\d]+), durability (?P<durability>[-\d]+), flavor (?P<flavor>[-\d]+), texture (?P<texture>[-\d]+), calories (?P<calories>[-\d]+)$`)
	for _, d := range data {
		match := regex.FindStringSubmatch(d)
		name := match[slices.Index(regex.SubexpNames(), "name")]
		capacity, _ := strconv.Atoi(match[slices.Index(regex.SubexpNames(), "capacity")])
		durability, _ := strconv.Atoi(match[slices.Index(regex.SubexpNames(), "durability")])
		flavor, _ := strconv.Atoi(match[slices.Index(regex.SubexpNames(), "flavor")])
		texture, _ := strconv.Atoi(match[slices.Index(regex.SubexpNames(), "texture")])
		calories, _ := strconv.Atoi(match[slices.Index(regex.SubexpNames(), "calories")])
		ingredients[name] = Ingredient{name: name, capacity: capacity, durability: durability, flavor: flavor, texture: texture, calories: calories}
	}
	return ingredients
}

func step1(ingredients map[string]Ingredient) (Recipe, int) {
	recipe := Recipe{}
	for _, ingredient := range ingredients {
		recipe[ingredient] = 1
	}
	for recipe.count() < 100 {
		best := 0
		best_recipe := Recipe{}
		for _, ing := range ingredients {
			try_recipe := Recipe{}
			for _, ingredient := range ingredients {
				try_recipe[ingredient] = recipe[ingredient]
			}
			try_recipe[ing]++
			try_score := try_recipe.getScore()
			if try_score > best {
				best = try_score
				best_recipe = try_recipe
			}
		}
		recipe = best_recipe
	}
	return recipe, recipe.getScore()
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return i * -1
}

func step2(current_recipe Recipe, ingredients map[string]Ingredient) (Recipe, int) {
	for {
		current_cal := current_recipe.getCalories()
		current_score := current_recipe.getScore()
		next_best := 0
		best_recipe := Recipe{}
		for _, i1 := range ingredients {
			for _, i2 := range ingredients {
				if i1.name != i2.name {
					try_recipe := Recipe{}
					for _, ingredient := range ingredients {
						try_recipe[ingredient] = current_recipe[ingredient]
					}
					try_recipe.replace(i1, i2)
					if abs(try_recipe.getCalories()-500) > abs(current_cal-500) {
						continue
					}
					if abs(try_recipe.getCalories()-500) == abs(current_cal-500) && try_recipe.getScore() <= current_score {
						continue
					}
					try_score := try_recipe.getScore()
					if try_score > next_best {
						next_best = try_score
						best_recipe = try_recipe
					}
				}
			}
		}
		if next_best > 0 {
			current_recipe = best_recipe
			continue
		}
		break
	}
	return current_recipe, current_recipe.getScore()
}
