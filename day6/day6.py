#!/usr/bin/env python3

import os
from typing import List

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))


def read_input(path: str) -> List[str]:
    with open(path, 'r') as f:
        lines = f.read().splitlines()
    return lines


def parse_line(line):
    match line.split():
        case ["turn", action, upper_corner, "through", bottom_corner] | [action, upper_corner, "through", bottom_corner]:
            [x1, y1] = upper_corner.split(",")
            [x2, y2] = bottom_corner.split(",")
            x1 = int(x1); x2 = int(x2); 
            y1 = int(y1); y2 = int(y2); 
        case _:
            print("ERROR", line)
    return action, x1, y1, x2, y2


def count_board(board):
    result = 0
    for column in board:
        for value in column:
            result += value
    return result


def run1(lines: List[str]):
    board = [ [0]*1000 for _ in range(1000) ]
    for line in lines:
        action, x1, y1, x2, y2 = parse_line(line)
        for y in range(y1, y2+1):
            for x in range(x1, x2+1):
                if action == "on":
                    board[x][y] = 1
                if action == "off":
                    board[x][y] = 0
                if action == "toggle":
                    if board[x][y] == 0:
                        board[x][y] = 1
                    else:
                        board[x][y] = 0

    return count_board(board)


INPUT = f"{SCRIPT_DIR}/input.txt"
data = read_input(INPUT)
result1 = run1(data)
print("Result1 = ", result1)

# =========================================


def run2(lines: List[str]):
    board = [ [0]*1000 for _ in range(1000) ]
    for line in lines:
        action, x1, y1, x2, y2 = parse_line(line)
        for y in range(y1, y2+1):
            for x in range(x1, x2+1):
                if action == "on":
                    board[x][y] += 1
                if action == "off" and board[x][y] > 0:
                    board[x][y] -= 1
                if action == "toggle":
                    board[x][y] += 2
    
    return count_board(board)

# INPUT = f"{SCRIPT_DIR}/input_test.txt"
data = read_input(INPUT)
result2 = run2(data)
print("Result2 = ", result2)
