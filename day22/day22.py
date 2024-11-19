#!/usr/bin/env python3

BOSS_HEALTH = 51
BOSS_DAMAGE = 9
HERO_HEALTH = 50
HERO_MANA = 500


class Character():

    def __init__(self, health):
        self.health = health
        self.damage = BOSS_DAMAGE
        self.poison = 0
        self.shield = 0
        self.recharge = 0
        self.mana = HERO_MANA
        self.mana_spent = 0

    def __repr__(self):
        return (f"[H: {self.health}, M: {self.mana}]")

    def copy(self):
        new = Character(self.health)
        new.damage = self.damage
        new.poison = self.poison
        new.shield = self.shield
        new.recharge = self.recharge
        new.mana = self.mana
        new.mana_spent = self.mana_spent
        return new
    
    def start_turn(self):
        if self.poison > 0:
            self.poison -= 1
            self.health -= 3

        if self.shield > 0:
            self.shield -= 1

        if self.recharge > 0:
            self.recharge -= 1
            self.mana += 101


def run(hero, boss, best=0, chain=[], turn="hero", difficulty="normal"):

    hero.start_turn()
    boss.start_turn()
    if boss.health <= 0:
        if hero.mana_spent < best or best == 0:
            best = hero.mana_spent
            return best

    if turn == "hero":
        if difficulty == "hard":    # 2nd star
            hero.health -= 1
            if hero.health <= 0:
                return best
        for spell in ["missile", "drain", "shield", "poison", "recharge"]:
            current_hero = hero.copy()
            current_boss = boss.copy()
            new_chain = chain + [spell]
            match spell:
                case "missile":
                    if current_hero.mana < 53:
                        continue
                    current_hero.mana -= 53
                    current_hero.mana_spent += 53
                    current_boss.health -= 4
                case "drain":
                    if current_hero.mana < 73:
                        continue
                    current_hero.mana -= 73
                    current_hero.mana_spent += 73
                    current_hero.health += 2
                    current_boss.health -= 2
                case "poison":
                    if current_hero.mana < 173 or current_boss.poison > 0:
                        continue
                    current_hero.mana -= 173
                    current_hero.mana_spent += 173
                    current_boss.poison += 6
                case "recharge":
                    if current_hero.mana < 229 or current_hero.recharge > 0:
                        continue
                    current_hero.mana -= 229
                    current_hero.mana_spent += 229
                    current_hero.recharge += 5
                case "shield":
                    if current_hero.mana < 113 or current_hero.shield > 0:
                        continue
                    current_hero.mana -= 113
                    current_hero.mana_spent += 113
                    current_hero.shield += 6
            if current_hero.mana_spent > best and best > 0:
                return best
            if current_boss.health <= 0:
                if current_hero.mana_spent < best or best == 0:
                    best = current_hero.mana_spent
                    return best
            best = run(current_hero, current_boss, best, new_chain, turn="boss", difficulty=difficulty)

    if turn == "boss":
        if hero.shield > 0:
            hero.health -= boss.damage - 7
        else:
            hero.health -= boss.damage
        if hero.health <= 0:
            return best
        best = run(hero, boss, best, chain, turn="hero", difficulty=difficulty)
        
    return best


hero = Character(HERO_HEALTH)
boss = Character(BOSS_HEALTH)
result1 = run(hero, boss)
print("Result1 = ", result1)

hero = Character(HERO_HEALTH)
boss = Character(BOSS_HEALTH)
result2 = run(hero, boss, difficulty="hard")
print("Result2 = ", result2)