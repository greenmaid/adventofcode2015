package main

import (
	"fmt"
)

const DAY = "21"

type Gear struct {
	name                string
	cost, damage, armor int
}

type Character struct {
	life, attack, defense int
}

func main() {

	weapons := []Gear{
		{"Dagger", 8, 4, 0},
		{"Shortsword", 10, 5, 0},
		{"Warhammer", 25, 6, 0},
		{"Longsword", 40, 7, 0},
		{"Greataxe", 74, 8, 0},
	}

	armors := []Gear{
		{"Leather", 13, 0, 1},
		{"Chainmail", 31, 0, 2},
		{"Splintmail", 53, 0, 3},
		{"Bandedmail", 75, 0, 4},
		{"Platemail", 102, 0, 5},
	}

	rings := []Gear{
		{"Damage+1", 25, 1, 0},
		{"Damage+2", 50, 2, 0},
		{"Damage+3", 100, 3, 0},
		{"Defense+1", 20, 0, 1},
		{"Defense+2", 40, 0, 2},
		{"Defense+3", 80, 0, 3},
	}

	boss := Character{life: 104, attack: 8, defense: 1}
	hero := Character{life: 100, attack: 0, defense: 0}

	result1, result2 := run(weapons, armors, rings, hero, boss)
	fmt.Println("Result1 = ", result1)
	fmt.Println("Result2 = ", result2)
}

func isvictory(originalHero Character, originalBoss Character) bool {

	hero := Character{originalHero.life, originalHero.attack, originalHero.defense}
	boss := Character{originalBoss.life, originalBoss.attack, originalBoss.defense}
	for {
		boss.life -= hero.attack - boss.defense
		if boss.life <= 0 {
			return true
		}

		damageToHero := boss.attack - hero.defense
		if damageToHero < 1 {
			damageToHero = 1
		}
		hero.life -= damageToHero
		if hero.life <= 0 {
			return false
		}
	}
}

func getPairOfRings(rings []Gear) []Gear {
	pairs := []Gear{}
	for i := 0; i < len(rings)-1; i++ {
		for j := i + 1; j < len(rings); j++ {
			cumulatedName := fmt.Sprintf("%s%s", rings[i].name, rings[j].name)
			cumulatedCost := rings[i].cost + rings[j].cost
			cumulatedDamage := rings[i].damage + rings[j].damage
			cumulatedArmor := rings[i].armor + rings[j].armor
			cumulatedRings := Gear{cumulatedName, cumulatedCost, cumulatedDamage, cumulatedArmor}
			pairs = append(pairs, cumulatedRings)
		}
	}
	return pairs
}

func run(weapons []Gear, armors []Gear, rings []Gear, hero Character, boss Character) (int, int) {
	cheapest := 100000
	mostExpensive := 0
	emptyGear := Gear{"empty", 0, 0, 0}
	possibleWeapons := []Gear{}
	possibleWeapons = append(possibleWeapons, weapons...)
	possibleArmors := []Gear{emptyGear}
	possibleArmors = append(possibleArmors, armors...)
	possibleRings := []Gear{emptyGear}
	possibleRings = append(possibleRings, rings...)
	possibleRings = append(possibleRings, getPairOfRings(rings)...)

	for _, weapon := range possibleWeapons {
		for _, armor := range possibleArmors {
			for _, ring := range possibleRings {
				cost := weapon.cost + armor.cost + ring.cost
				hero.attack = weapon.damage + armor.damage + ring.damage
				hero.defense = weapon.armor + armor.armor + ring.armor
				if isvictory(hero, boss) && cost < cheapest {
					cheapest = cost

				}
				if !isvictory(hero, boss) && cost > mostExpensive {
					mostExpensive = cost

				}
			}
		}
	}
	return cheapest, mostExpensive
}
