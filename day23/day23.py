#!/usr/bin/env python3

import os
from typing import List

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))


def read_input(path: str) -> List[str]:
    with open(path, 'r') as f:
        lines = f.read().splitlines()
    return lines


def evaluate(number, registers):
    if number in registers:
        return registers[number]
    return int(number)


def run(program: List[str], registers):
    current_offset = 0
    while current_offset < len(program) and current_offset >= 0:
        match program[current_offset].split(" "):
            case ["hlf", r]:
                registers[r] = int(evaluate(r, registers) / 2)
                current_offset += 1
            case ["tpl", r]:
                registers[r] = evaluate(r, registers) * 3
                current_offset += 1
            case ["inc", r]:
                registers[r] += 1
                current_offset += 1
            case ["jmp", off]:
                offset_jmp = evaluate(off, registers)
                current_offset += offset_jmp
            case ["jmp", o]:
                offset_jmp = evaluate(o, registers)
                current_offset += offset_jmp
            case ["jie", r, o]:
                r = r[:-1]  # eliminates ","
                if evaluate(r, registers) % 2 == 0:
                    offset_jmp = evaluate(o, registers)
                    current_offset += offset_jmp
                else:
                    current_offset += 1
            case ["jio", r, o]:
                r = r[:-1]  # eliminates ","
                if evaluate(r, registers) == 1:
                    offset_jmp = evaluate(o, registers)
                    current_offset += offset_jmp
                else:
                    current_offset += 1
            case _:
                print("Unkonwn instruction", program[current_offset])

    return registers


INPUT = f"{SCRIPT_DIR}/input.txt"
program = read_input(INPUT)
registers = {
        "a": 0,
        "b": 0,
    }
result1 = run(program, registers)
print("Result1 = ", result1["b"])

# =========================================
registers = {
        "a": 1,
        "b": 0,
    }
result2 = run(program, registers)
print("Result2 = ", result2["b"])
