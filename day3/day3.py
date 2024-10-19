#!/usr/bin/env python3

import os
from typing import List

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))


def read_input(path: str) -> str:
    with open(path, 'r') as f:
        str = f.read()
    return str


def run1(input):
    current = (0, 0)
    visited = {current}
    for c in input:
        if c == ">":
            current = (current[0], current[1]+1)
        if c == "<":
            current = (current[0], current[1]-1)
        if c == "^":
            current = (current[0]+1, current[1])
        if c == "v":
            current = (current[0]-1, current[1])
        visited.add(current)
    return len(visited)


#INPUT = f"{SCRIPT_DIR}/input_test.txt"
INPUT = f"{SCRIPT_DIR}/input.txt"
data = read_input(INPUT)
result1 = run1(data)
print("Result1 = ", result1)

# =========================================


class Santa:
    def __init__(self):
        self.location = (0, 0)
        self.visited = {self.location}

    def go_north(self):
        self.location = (self.location[0]+1, self.location[1])
        self.visited.add(self.location)
    def go_south(self):
        self.location = (self.location[0]-1, self.location[1])
        self.visited.add(self.location)
    def go_west(self):
        self.location = (self.location[0], self.location[1]-1)
        self.visited.add(self.location)
    def go_east(self):
        self.location = (self.location[0], self.location[1]+1)
        self.visited.add(self.location)
    


def run2(input):
    santa = Santa()
    robo_santa = Santa()
    index = 0
    for c in input:
        index += 1
        if index % 2 == 1:
            current_moving = santa
        else:
            current_moving = robo_santa

        if c == ">":
            current_moving.go_east()
        if c == "<":
            current_moving.go_west()
        if c == "^":
            current_moving.go_north()
        if c == "v":
            current_moving.go_south()
        
    visited = santa.visited.union(robo_santa.visited)
    return len(visited)


#INPUT = f"{SCRIPT_DIR}/input_test.txt"
data = read_input(INPUT)
result2 = run2(data)
print("Result2 = ", result2)
