#!/usr/bin/env python3

import os
from typing import List


def look_and_say(phrase):
    index = 0
    result = ""
    while index < len(phrase):
        char = phrase[index]
        count = 0
        # print(phrase, index, count, char)
        while index+count < len(phrase) and phrase[index + count] == char:
            count += 1
        result += str(count) + char
        index += count
    return result


def run(puzzle_input, loop_count):
    phrase = puzzle_input
    for _ in range(loop_count):
        phrase = look_and_say(phrase)
    return len(phrase)


puzzle_input = "1113122113"
result1 = run(puzzle_input, 40)
print("Result1 = ", result1)

result2 = run(puzzle_input, 50)
print("Result2 = ", result2)
