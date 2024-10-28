#!/usr/bin/env python3

import os
import sys
from typing import List

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))


def read_input(path: str) -> List[str]:
    with open(path, 'r') as f:
        lines = f.read().splitlines()
    return lines


def parse(lines):
    distances = {}
    for line in lines:
        match line.split(" "):
            case [city1, "to", city2, "=", distance]:
                distances.setdefault(city1,{})[city2] = int(distance)
                distances.setdefault(city2,{})[city1] = int(distance)
            case _:
                print("ERROR", line)
                sys.exit(1)
    return distances


def permutations(elements):
    if len(elements) <= 1:
        yield elements
        return
    for perm in permutations(elements[1:]):
        for i in range(len(elements)):
            yield perm[:i] + elements[0:1] + perm[i:]


def run(distances):
    path_proposals = list(permutations(list(distances.keys())))
    possible_path_lenghs = []
    for path in path_proposals:
        path_lengh = 0
        for i in range(len(path)-1):
            path_lengh += distances[path[i]][path[i+1]]
        # print(path, path_lengh)
        possible_path_lenghs.append(path_lengh)
    return min(possible_path_lenghs), max(possible_path_lenghs)


#INPUT = f"{SCRIPT_DIR}/input_test.txt"
INPUT = f"{SCRIPT_DIR}/input.txt"
distances = parse(read_input(INPUT))
result1, result2 = run(distances)
print("Result1 = ", result1)
print("Result2 = ", result2)
