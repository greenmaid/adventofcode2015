#!/usr/bin/env python3

import os
from typing import List

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))


def read_input(path: str) -> List[str]:
    with open(path, 'r') as f:
        lines = f.read().splitlines()
    return lines


def read_inputinto_table(path: str) -> List[List[str]]:
    with open(path, "r") as f:
        lines = f.read().splitlines()
    map = []
    for line in lines:
        map.append([s for s in line])
    return map


def is_good(line):
    double_letters = False
    vowels = 0
    if "ab" in line or "cd" in line or "pq" in line or "xy" in line:
        return False
    line = line + "*"
    for i in range(len(line)-1):
        if line[i] == line[i+1]:
            double_letters = True
        if line[i] in "aeiou":
            vowels += 1
    if double_letters and vowels >= 3:
        return True
    return False

    
def run1(lines: List[str]):
    count = 0
    for line in lines:
        if is_good(line):
            count += 1
    return count
    
INPUT = f"{SCRIPT_DIR}/input.txt"
data = read_input(INPUT)

result1 = run1(data)
print("Result1 = ", result1)

# =========================================

def is_good2(line):
    rule1 = False
    rule2 = False
    for i in range(len(line)-2):
        if line[i:i+2] in line[i+2:]:
            rule1 = True
        if line[i] == line[i+2]:
            rule2 = True
        if rule1 and rule2:
            return True
    return False


def run2(lines: List[str]):
    count = 0
    for line in lines:
        if is_good2(line):
            count += 1
    return count


data = read_input(INPUT)
result2 = run2(data)
print("Result2 = ", result2)
