#!/usr/bin/env python3

import os
from typing import List

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))


def read_input_into_table(path: str) -> List[List[str]]:
    with open(path, "r") as f:
        lines = f.read().splitlines()
    map = []
    for line in lines:
        map.append([s for s in line])
    return map


def display(map):
    [print(i) for i in map]
    print("")
    print("")


def generate_next_turn(map):
    max = len(map)   # assuming we have a square
    new_map = [['.'] * max for _ in range(max)]
    for y in range(max):
        for x in range(max):
            count_on = 0
            if x > 0 and y > 0 and map[x-1][y-1] == '#':
                count_on += 1
            if x > 0 and map[x-1][y] == '#':
                count_on += 1
            if x > 0 and y < max-1 and map[x-1][y+1] == '#':
                count_on += 1
            if y > 0 and map[x][y-1] == '#':
                count_on += 1
            if y < max-1 and map[x][y+1] == '#':
                count_on += 1
            if x < max-1 and map[x+1][y] == '#':
                count_on += 1
            if x < max-1 and y > 0 and map[x+1][y-1] == '#':
                count_on += 1
            if x < max-1 and y < max-1 and map[x+1][y+1] == '#':
                count_on += 1

            if map[x][y] == '#' and (count_on == 2 or count_on == 3):
                new_map[x][y] = "#"
            if map[x][y] != '#' and count_on == 3:
                new_map[x][y] = "#"
    return new_map


def count_result(map):
    result = 0
    for i in map:
        for j in i:
            if j == '#':
                result += 1
    return result


def run1(data, turns):
    table = [ i for i in data]  # deep copy
    for _ in range(turns):
        table = generate_next_turn(table)
    return count_result(table)


#INPUT, TURNS = f"{SCRIPT_DIR}/input_test.txt", 5
INPUT, TURNS = f"{SCRIPT_DIR}/input.txt", 100
data = read_input_into_table(INPUT)
result1 = run1(data, TURNS)
print("Result1 = ", result1)

# =========================================



def run2(data, turns):
    table = [ i for i in data]  # deep copy
    max = len(table)
    table[0][0] = '#'
    table[0][max-1] = '#'
    table[max-1][0] = '#'
    table[max-1][max-1] = '#'
    for _ in range(turns):
        table = generate_next_turn(table)
        table[0][0] = '#'
        table[0][max-1] = '#'
        table[max-1][0] = '#'
        table[max-1][max-1] = '#'
    return count_result(table)


result2 = run2(data, TURNS)
print("Result2 = ", result2)
