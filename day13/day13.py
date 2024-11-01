#!/usr/bin/env python3

import os
from typing import List

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))


def read_input(path: str) -> List[str]:
    with open(path, 'r') as f:
        lines = f.read().splitlines()
    return lines


def parse(data):
    happiness_map = {}
    for happiness_case in data:
        match happiness_case[:-1].split(" "):
            case [person, 'would', gain_lose, value, 'happiness', 'units', 'by', 'sitting', 'next', 'to', other]:
                value = int(value)
                if gain_lose == "lose":
                    value = -1 * value
                if not person in happiness_map:
                    happiness_map[person] = {}
                happiness_map[person][other] = value
    return happiness_map


def permutations(elements):
    if len(elements) <= 1:
        yield elements
        return
    for perm in permutations(elements[1:]):
        for i in range(len(elements)):
            yield perm[:i] + elements[0:1] + perm[i:]


def count_layout_happiness(layout, happiness_map):
    layout.append(layout[0])
    sum = 0
    for i in range(len(layout)-1):
        sum += happiness_map[layout[i]][layout[i+1]]
        sum += happiness_map[layout[i+1]][layout[i]]
    return sum


def run(happiness_map):
    people = list(happiness_map.keys())
    happiness_possibilities = []
    for table_layout in list(permutations(people[1:])):
        table_layout.append(people[0])
        count = count_layout_happiness(table_layout, happiness_map)
        happiness_possibilities.append(count)
    return max(happiness_possibilities)


#INPUT = f"{SCRIPT_DIR}/input_test.txt"
INPUT = f"{SCRIPT_DIR}/input.txt"
data = read_input(INPUT)
happiness_map = parse(data)
result1 = run(happiness_map)
print("Result1 = ", result1)

# =========================================


def add_me_in_happiness_map(happiness_map):
    happiness_map["me"] = {}
    for person in happiness_map.keys():
        if person != "me":
            happiness_map[person]["me"] = 0
            happiness_map["me"][person] = 0
    return happiness_map


happiness_map2 = add_me_in_happiness_map(happiness_map)
result2 = run(happiness_map2)
print("Result2 = ", result2)
