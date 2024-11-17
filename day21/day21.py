#!/usr/bin/env python3

import os
import itertools

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))


def read_input(path: str) -> str:
    with open(path, 'r') as f:
        str = f.read()
    return str


def parse_stuff(data):
    weapons = {}
    armors = {}
    rings = {}
    [w_descrip, a_descrip, r_descrip] = data.split("\n\n")

    for e in w_descrip.split("\n")[1:]:
        match list(filter(len, e.split(" "))):
            case [name, cost, damage, armor]:
                weapons[name] = {
                    "cost": int(cost),
                    "damage": int(damage),
                    "armor": int(armor),
                    }

    for e in a_descrip.split("\n")[1:]:
        match list(filter(len, e.split(" "))):
            case [name, cost, damage, armor]:
                armors[name] = {
                    "cost": int(cost),
                    "damage": int(damage),
                    "armor": int(armor),
                    }
                
    for e in r_descrip.split("\n")[1:]:
        match list(filter(len, e.split(" "))):
            case [name1, name2, cost, damage, armor]:
                rings[name1+name2] = {
                    "cost": int(cost),
                    "damage": int(damage),
                    "armor": int(armor),
                    }
                                
    return weapons, armors, rings


class Character():
    def __repr__(self):
        return (f"[H: {self.health}, D: {self.damage}, A: {self.armor}]")


class Boss(Character):
    def __init__(self):
        self.health = 104
        self.damage = 8
        self.armor = 1


class Hero(Character):
    def __init__(self):
        self.health = 100
        self.damage = 0
        self.armor = 0


def hero_wins(hero, boss):
    """
    Returns True if hero wins the fight
    """
    while True:
        hero_attack_damage = hero.damage - boss.armor
        if hero_attack_damage < 1:
            hero_attack_damage = 1
        boss.health -= hero_attack_damage

        if boss.health <= 0:
            return True

        boss_attack_damage = boss.damage - hero.armor
        if boss_attack_damage < 1:
            boss_attack_damage = 1
        hero.health -= boss_attack_damage

        if hero.health <= 0:
            return False


def run(weapons, armors, rings):

    weapon_choices = { k:v for k,v in weapons.items() }   # deep_copy

    armor_choices = { k:v for k,v in armors.items() }   # deep_copy
    armor_choices["none"] = {"cost": 0, "damage": 0, "armor": 0}

    ring_choices = { k:v for k,v in rings.items() }   # deep_copy
    ring_choices["none"] = {"cost": 0, "damage": 0, "armor": 0}
    for ring1, ring2 in itertools.combinations(rings, 2):
        ring_choices[f"{ring1}/{ring2}"] = {
            "cost": rings[ring1]["cost"] + rings[ring2]["cost"], 
            "damage": rings[ring1]["damage"] + rings[ring2]["damage"], 
            "armor": rings[ring1]["armor"] + rings[ring2]["armor"], 
            }

    # FOR RUN1    
    cheapest_stuff = ("", "", "")
    best_cost = 1000000

    # FOR RUN2
    more_expensive_stuff = ("", "", "")
    worst_cost = 0

    for weapon_name, weapon_stats in weapon_choices.items():
        for armor_name, armor_stats in armor_choices.items():
            for ring_name, ring_stats in ring_choices.items():

                stuff = (weapon_name, armor_name, ring_name)
                cost = weapon_stats["cost"] + armor_stats["cost"] + ring_stats["cost"]

                hero = Hero()
                boss = Boss()
                hero.damage += weapon_stats["damage"] + armor_stats["damage"] + ring_stats["damage"]
                hero.armor += weapon_stats["armor"] + armor_stats["armor"] + ring_stats["armor"]

                if hero_wins(hero, boss):   # run1
                    if cost < best_cost:
                        best_cost = cost
                        cheapest_stuff = stuff

                else:                       # run 2
                    if cost > worst_cost:
                        worst_cost = cost
                        more_expensive_stuff = stuff

    return cheapest_stuff, best_cost, more_expensive_stuff, worst_cost


INPUT = f"{SCRIPT_DIR}/input.txt"
data = read_input(INPUT)
weapons, armors, rings = parse_stuff(data)
stuff1, cost1, stuff2, cost2 = run(weapons, armors, rings)
print("Result1 = ", cost1, "with", stuff1)
print("Result2 = ", cost2, "with", stuff2)
