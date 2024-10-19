#!/usr/bin/env python3

import os
import hashlib
from typing import List

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))


def read_input(path: str) -> str:
    with open(path, 'r') as f:
        str = f.read()
    return str

def md5(string):
    return hashlib.md5(string.encode('utf-8')).hexdigest()

def run1(secret):
    number = 0
    while True:
        hash = md5(secret + str(number))
        if hash.startswith("00000"):
            return number
        number += 1


#INPUT = f"{SCRIPT_DIR}/input_test.txt"
INPUT = f"{SCRIPT_DIR}/input.txt"
data = read_input(INPUT)
result1 = run1(data)
print("Result1 = ", result1)

# =========================================


def run2(secret):
    number = 0
    while True:
        hash = md5(secret + str(number))
        if hash.startswith("000000"):
            return number
        number += 1


#INPUT = f"{SCRIPT_DIR}/input_test.txt"
data = read_input(INPUT)
result2 = run2(data)
print("Result2 = ", result2)
