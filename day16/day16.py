#!/usr/bin/env python3

import os
from typing import List

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))


def read_input(path: str) -> List[str]:
    with open(path, 'r') as f:
        lines = f.read().splitlines()
    return lines


def parse(lines):
    aunts = []
    for line in lines:
        name, description = line.split(": ", 1)
        aunt = {"name": name}
        for prop in description.split(", "):
            match prop.split(": "):
                case [property, value]:
                    aunt[property] = int(value)
        aunts.append(aunt)
    return aunts


MYSTERY_AUNT_PROPERTIES = {
    "children": 3,
    "cats": 7,
    "samoyeds": 2,
    "pomeranians": 3,
    "akitas": 0,
    "vizslas": 0,
    "goldfish": 5,
    "trees": 3,
    "cars": 2,
    "perfumes": 1,
}


def run1(aunts):
    possible_aunts = [ a for a in aunts]  # deep copy
    for k, v in MYSTERY_AUNT_PROPERTIES.items():
        new_possibles = []
        for aunt in possible_aunts:
            if k in aunt and aunt[k] != v:
                pass
            else:
                new_possibles.append(aunt)
        possible_aunts = new_possibles
    if len(possible_aunts) == 1:
        return possible_aunts[0]["name"]


INPUT = f"{SCRIPT_DIR}/input.txt"
lines = read_input(INPUT)
aunts = parse(lines)
result1 = run1(aunts)
print("Result1 = ", result1)

# =========================================


def run2(aunts):
    possible_aunts = [ a for a in aunts]  # deep copy
    for k, v in MYSTERY_AUNT_PROPERTIES.items():
        new_possibles = []
        for aunt in possible_aunts:
            if k in aunt:
                if (k == "cats" or k == "trees"):
                    if aunt[k] <= v:
                        continue
                elif (k == "pomeranians" or k == "goldfish"):
                    if aunt[k] >= v:
                        continue
                elif aunt[k] != v:
                    continue
                new_possibles.append(aunt)
            else:
                new_possibles.append(aunt)
        possible_aunts = new_possibles
    if len(possible_aunts) == 1:
        return possible_aunts[0]["name"]

result2 = run2(aunts)
print("Result2 = ", result2)
