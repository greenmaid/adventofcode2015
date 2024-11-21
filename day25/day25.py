#!/usr/bin/env python3

import sys

START = 20151125

TARGET_ROW = 2947
TARGET_COLUMN = 3029

def next_code(code):
    return (code * 252533 ) % 33554393


def run1(start):
    x = 2
    y = 0
    code = start
    while True:
        x, y = y, x
        x = 1
        while True:
            code = next_code(code)

            if x == TARGET_COLUMN and y == TARGET_ROW:
                return code

            y -= 1
            x += 1

            if y == 0:
                break



result1 = run1(START)
print("Result1 = ", result1)

# =========================================



def run2():
    pass


result2 = run2()
print("Result2 = ", result2)
