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


def run1(lines: List[str]):
    result = 0
    for line in lines:
        dimensions = line.split("x")
        l = int(dimensions[0])
        w = int(dimensions[1])
        h = int(dimensions[2])
        need = 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)
        result += need
    return result

#INPUT = f"{SCRIPT_DIR}/input_test.txt"
INPUT = f"{SCRIPT_DIR}/input.txt"
data = read_input(INPUT)
result1 = run1(data)
print("Result1 = ", result1)

# =========================================

# INPUT = f"{SCRIPT_DIR}/input_test.txt"


def run2(lines: List[str]):
    result = 0
    for line in lines:
        dimensions = line.split("x")
        l = int(dimensions[0])
        w = int(dimensions[1])
        h = int(dimensions[2])
        need = l*w*h + min(l+w, w+h, h+l)*2
        result += need
    return result


data = read_input(INPUT)
result2 = run2(data)
print("Result2 = ", result2)
