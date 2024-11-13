#!/usr/bin/env python3

import os
import random

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))


def read_input(path: str) -> str:
    with open(path, 'r') as f:
        data = f.read()
    return data


def parse(data):
    [replacements, molecule] = data.split("\n\n")

    replacement_map = {}
    for line in replacements.split("\n"):
        match line.split(" "):
            case [source, "=>", dest]:
                if not source in replacement_map:
                    replacement_map[source] = []
                replacement_map[source].append(dest)
    return replacement_map, molecule


def run1(replacements, molecule):
    possibles = set()
    for old, news in replacements.items():
        index = 0
        while True:
            index = molecule.find(old, index)
            if index == -1:
                break
            for new in news:
                new_molecule = molecule[:index] + new + molecule[index+len(old):]
                possibles.add(new_molecule)
            index += 1
    return possibles, len(possibles)



#INPUT = f"{SCRIPT_DIR}/input_test.txt"
INPUT = f"{SCRIPT_DIR}/input.txt"
data = read_input(INPUT)
replacements, molecule = parse(data)
_, result1 = run1(replacements, molecule)
print("Result1 = ", result1)

# =========================================


def parse_reverse(data):
    [replacements, molecule] = data.split("\n\n")

    reverse_replacement_map = {}
    for line in replacements.split("\n"):
        match line.split(" "):
            case [source, "=>", dest]:
                if not dest in reverse_replacement_map:
                    reverse_replacement_map[dest] = []
                reverse_replacement_map[dest].append(source)
    return reverse_replacement_map, molecule


CACHE = set()

def run2(replacements, molecule):
    results = []
    for _ in range(3):   # Let's do the computation 5 times to ensure to converge towards best solution
        steps = 0
        current = molecule
        while current != "e":
            steps += 1
            possible_next, _ = run1(replacements, current)
            if ("e" in current and len(current) != 1) or (len(possible_next) == 0):
                steps = 0
                current = molecule
            else:
                current = random.choice(list(possible_next))
                # print(current)
        results.append(steps)
    return min(results)

   
reverse_replacements, molecule = parse_reverse(data)
result2 = run2(reverse_replacements, molecule)
print("Result2 = ", result2)
