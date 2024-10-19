#!/usr/bin/env python3

import os
from typing import List

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))
# INPUT = f"{SCRIPT_DIR}/input_test.txt"
INPUT = f"{SCRIPT_DIR}/input.txt"


def read_input(path: str) -> str:
    with open(path, 'r') as f:
        str = f.read()
    return str



def run1(str):
    floor = 0
    for c in str:
        if c == '(':
            floor +=1
        if c == ')':
            floor -=1
    return floor

data = read_input(INPUT)
result1 = run1(data)
print("Result1 = ", result1)

# =========================================

# INPUT = f"{SCRIPT_DIR}/input_test.txt"


def run2(str):
    floor = 0
    index = 0
    for c in str:
        if c == '(':
            floor +=1
        if c == ')':
            floor -=1
        index += 1
        if floor == -1:
            return index


data = read_input(INPUT)
result2 = run2(data)
print("Result2 = ", result2)
