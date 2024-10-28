#!/usr/bin/env python3

import os
from typing import List

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))


def read_input(path: str) -> List[str]:
    with open(path, 'r') as f:
        lines = f.read().splitlines()
    return lines


def run1(lines: List[str]):
    printed_char_count = 0
    memory_char_count = 0
    for line in lines:
        printed_char_count += len(line)
        i = 1
        memory = 0
        while i < len(line) -1:
            if line[i] == "\\" and line[i+1] == "x":
                memory += 1
                i += 4
                continue
            if line[i] == "\\":
                memory += 1
                i += 2
                continue
            else:
                memory += 1
                i += 1
        memory_char_count += memory

    return printed_char_count, memory_char_count


#INPUT = f"{SCRIPT_DIR}/input_test.txt"
INPUT = f"{SCRIPT_DIR}/input.txt"
data = read_input(INPUT)
printed_char_count, memory_char_count = run1(data)
print("Result1 = ", printed_char_count - memory_char_count)

# =========================================


def run2(lines: List[str]):
    total_char_count = 0
    for line in lines:
        i = 0
        char_count = 2
        while i < len(line):
            if line[i] in ["\\", "\"", "\'"]:
                char_count += 2
                i += 1
            else:
                char_count += 1
                i += 1
        total_char_count += char_count
    return total_char_count

encoded = run2(data)
print("Result2 = ", encoded - printed_char_count)
