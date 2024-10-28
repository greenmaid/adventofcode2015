#!/usr/bin/env python3

import os
from typing import List

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))


def read_input(path: str) -> List[str]:
    with open(path, 'r') as f:
        lines = f.read().splitlines()
    return lines


def parse(lines):
  wires = {}
  for line in lines:
        match line.split(" "):
            case [val1, op, val2, "->", result]:
                wires[result] = [op, val1, val2]
            case ["NOT", val, "->", result]:
                wires[result] = ["NOT", val]
            case [val, "->", result]:
                wires[result] = val
  return wires


def evaluate(value, wires):

    if value in CACHE:
        return CACHE[value]

    if value in wires:

        definition = wires[value]
        try:
            return int(definition)
        except:
            pass

        match definition[0]:
            case "AND":
                result = evaluate(definition[1], wires) & evaluate(definition[2], wires)
            case "OR":
                result = evaluate(definition[1], wires) | evaluate(definition[2], wires) 
            case "LSHIFT":
                result = evaluate(definition[1], wires) << evaluate(definition[2], wires)
            case "RSHIFT":
                result = evaluate(definition[1], wires) >> evaluate(definition[2], wires)
            case "NOT":
                result = ~ evaluate(definition[1], wires)
            case _:
                result = evaluate(definition, wires)
        
        result = result & 0xffff
        CACHE[value] = result
        return result

    return int(value)



def run(wires):
    a = evaluate("a", wires)
    return a

# INPUT = f"{SCRIPT_DIR}/input_test.txt"

INPUT = f"{SCRIPT_DIR}/input.txt"
data = read_input(INPUT)
wires = parse(data)
CACHE = {}
result1 = run(wires)
print("Result1 = ", result1)

# =========================================

CACHE = {}
wires['b'] = result1
result2 = run(wires)
print("Result2 = ", result2)
