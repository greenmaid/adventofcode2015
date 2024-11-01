#!/usr/bin/env python3

import json
import os
from typing import List

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))


def read_input(path: str) -> str:
    with open(path, 'r') as f:
        str = f.read()
    return str


def run1(data):
    sum = 0
    red_elems = ""
    for c in data:
        if c in "-0123456789":
            red_elems += c
        elif red_elems != "" and red_elems != "-":
            sum += int(red_elems)
            red_elems = ""
        else:
            red_elems = ""
    return sum


#INPUT = f"{SCRIPT_DIR}/input_test.txt"
INPUT = f"{SCRIPT_DIR}/input.txt"
data = read_input(INPUT)
result1 = run1(data)
print("Result1 = ", result1)

# =========================================


def extract_red(json_obj, red_elems=[]):
    if isinstance(json_obj, dict):
        for _, v in json_obj.items():
            if v == "red":
                red_elems.append(json_obj)
                return red_elems
        for _, v in json_obj.items():
            extract_red(v, red_elems)
    if isinstance(json_obj, list):
        for elem in json_obj:
            extract_red(elem, red_elems)
    return red_elems


def run2(json_data):
    red = extract_red(json_data)
    return run1(str(red))


with open(INPUT, 'r') as file:
    json_data = json.load(file)

red = run2(json_data)
print("Result2 = ", result1 - red)
