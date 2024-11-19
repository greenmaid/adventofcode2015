#!/usr/bin/env python3

import os
import math
from itertools import chain, combinations

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))


def read_input(path):
    with open(path, 'r') as f:
        lines = f.read().splitlines()
    return [int(l) for l in lines]


def all_subsets(ss):
    return chain(*map(lambda x: combinations(ss, x), range(len(ss))))


def get_quantum_entanglement(packages):
    return math.prod(packages)


def check_remaining_can_by_divided_in_balanced(package_group, packages):
    group_expected_weight = sum(package_group)
    remaining = [ k  for k in packages ]
    for e in package_group:
        remaining.remove(e)
    return any(sum(g) == group_expected_weight for g in all_subsets(remaining))


def run1(packages):
    total_weight = sum(packages)
    group_expected_weight = int(total_weight / 3)

    best = get_quantum_entanglement(packages)
    least_package_count = 0

    for i in range(len(packages)-2):
        for package_group in  combinations(packages, i+1):
            if sum(package_group) == group_expected_weight:
                if check_remaining_can_by_divided_in_balanced(package_group, packages):
                    least_package_count = len(package_group)
                    qe = get_quantum_entanglement(package_group)
                    if qe < best:
                        best = qe
        if least_package_count > 0:
            break
    return best


INPUT = f"{SCRIPT_DIR}/input_test.txt"
INPUT = f"{SCRIPT_DIR}/input.txt"
packages = read_input(INPUT)
result1 = run1(packages)
print("Result1 = ", result1)

# =========================================


def check_remaining_can_by_divided_in_3_balanced(package_group, packages):
    group_expected_weight = sum(package_group)
    remaining = [ k  for k in packages ]
    for e in package_group:
        remaining.remove(e)
    for g in all_subsets(remaining):
        if sum(g) == group_expected_weight:
            if check_remaining_can_by_divided_in_balanced(g, remaining):
                return True
    return False


def run2(packages):
    total_weight = sum(packages)
    group_expected_weight = int(total_weight / 4)

    best = get_quantum_entanglement(packages)
    least_package_count = 0

    for i in range(len(packages)-2):
        for package_group in  combinations(packages, i+1):
            if sum(package_group) == group_expected_weight:
                if check_remaining_can_by_divided_in_3_balanced(package_group, packages):
                    least_package_count = len(package_group)
                    qe = get_quantum_entanglement(package_group)
                    if qe < best:
                        best = qe
        if least_package_count > 0:
            break
    return best


packages = read_input(INPUT)
result2 = run2(packages)
print("Result2 = ", result2)
