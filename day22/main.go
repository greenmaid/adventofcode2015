package main

import (
	"fmt"
)

const DAY = "22"

const BOSS_HEALTH = 51
const BOSS_ATTACK = 9
const HERO_HEALTH = 50
const HERO_MANA = 500

type Character struct {
	life, attack, armor, mana, shield, poison, recharge, manaSpent int
}

func main() {

	hero := Character{life: HERO_HEALTH, mana: HERO_MANA}
	boss := Character{life: BOSS_HEALTH, attack: BOSS_ATTACK}

	result1 := step1(hero, boss)
	fmt.Println("Result1 = ", result1)

	result2 := step2(hero, boss)
	fmt.Println("Result2 = ", result2)
}

func startTurn(char *Character) {
	if char.poison > 0 {
		char.poison--
		char.life -= 3
	}
	if char.recharge > 0 {
		char.recharge--
		char.mana += 101
	}
	if char.shield > 0 {
		char.shield--
		char.armor = 7
	} else {
		char.armor = 0
	}
}

func nextTurn(originalHero Character, originalBoss Character, turn string, best int, spellChain []string, difficulty string) int {

	hero := originalHero
	startTurn(&hero)
	boss := originalBoss
	startTurn(&boss)

	if boss.life <= 0 {
		// win
		if hero.manaSpent < best {
			best = hero.manaSpent
		}
		return best
	}

	if turn == "hero" {
		if difficulty == "hard" {
			hero.life -= 1
			if hero.life <= 0 {
				//loose
				return best
			}
		}
		if originalHero.mana < 53 {
			// loose out of mana
			return best
		}

		heroSavedStatus := hero
		bossSavedStatus := boss

		for _, spell := range []string{"missile", "shield", "drain", "poison", "recharge"} {

			hero = heroSavedStatus
			boss = bossSavedStatus

			switch spell {
			case "missile":
				cost := 53
				if hero.mana < cost {
					continue
				}
				hero.mana -= cost
				hero.manaSpent += cost
				boss.life -= 4
			case "drain":
				cost := 73
				if hero.mana < cost {
					continue
				}
				hero.mana -= cost
				hero.manaSpent += cost
				hero.life += 2
				boss.life -= 2
			case "shield":
				cost := 113
				if hero.shield > 0 || hero.mana < cost {
					continue
				}
				hero.mana -= cost
				hero.manaSpent += cost
				hero.shield = 6
			case "poison":
				cost := 173
				if boss.poison > 0 || hero.mana < cost {
					continue
				}
				hero.mana -= cost
				hero.manaSpent += cost
				boss.poison = 6
			case "recharge":
				cost := 229
				if hero.recharge > 0 || hero.mana < cost {
					continue
				}
				hero.mana -= cost
				hero.manaSpent += cost
				hero.recharge = 5
			}

			if hero.manaSpent >= best {
				return best
			}

			if boss.life <= 0 {
				// win
				if hero.manaSpent < best {
					best = hero.manaSpent
				}

			} else {
				best = nextTurn(hero, boss, "boss", best, append(spellChain, spell), difficulty)
			}
		}

	}
	if turn == "boss" {

		hero.life -= boss.attack - hero.armor
		if hero.life <= 0 {
			// loose
			return best
		} else {
			best = nextTurn(hero, boss, "hero", best, spellChain, difficulty)
		}
	}
	return best
}

func step1(hero Character, boss Character) int {
	newHero := hero
	newBoss := boss
	result := nextTurn(newHero, newBoss, "hero", 1000000, []string{}, "normal")
	return result
}

func step2(hero Character, boss Character) int {
	newHero := hero
	newBoss := boss
	result := nextTurn(newHero, newBoss, "hero", 1000000, []string{}, "hard")
	return result
}
