#!/usr/bin/env python3

import os
from typing import List

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))


def read_input(path: str) -> List[str]:
    with open(path, 'r') as f:
        lines = f.read().splitlines()
    return [int(l) for l in lines]


def test_capaciy(containers, bucket, target, possibles):
    next_containers = [ k for k in containers ]
    current = next_containers.pop()
    bucket_with_current = bucket + [current]
    if sum(bucket_with_current) == target:
        possibles.append(bucket_with_current)
    elif sum(bucket_with_current) < target and next_containers:
        test_capaciy(next_containers, bucket_with_current, target, possibles)
    if next_containers:
        test_capaciy(next_containers, bucket, target, possibles)


def run1(containers, target):
    possibles = []
    test_capaciy(containers, [], target, possibles)
    return possibles, len(possibles)
    

#INPUT, TARGET = f"{SCRIPT_DIR}/input_test.txt", 25
INPUT, TARGET = f"{SCRIPT_DIR}/input.txt", 150
containers = read_input(INPUT)
possibles, result1 = run1(containers, TARGET)
print("Result1 = ", result1)

# =========================================


def run2(possibles):
    minimum = len(possibles[0])
    min_count = 0
    for p in possibles:
        if len(p) < minimum:
            minimum = len(p)
            min_count = 1
        elif len(p) == minimum:
            min_count += 1
    return minimum, min_count


_, result2 = run2(possibles)
print("Result2 = ", result2)
