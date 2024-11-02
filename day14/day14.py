#!/usr/bin/env python3

import os
from typing import List

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))


def read_input(path: str) -> List[str]:
    with open(path, 'r') as f:
        lines = f.read().splitlines()
    return lines


def parse(data):
    reindeers = {}
    for reindeer in data:
        match reindeer[:-1].split(" "):
            case [name, 'can', 'fly', speed, 'km/s', 'for', endurance, 'seconds,', 'but', 'then', 'must', 'rest', 'for', rest_time, 'seconds']:
                reindeers[name] = {"speed": int(speed), "endurance": int(endurance), "rest_time": int(rest_time)}
    return reindeers

def update_scores_part2(race):
    best = max([r["distance"] for _ , r in race.items()])
    for name, data in race.items():
        if data["distance"] == best:
            data["score_part2"] += 1


def run(reindeers, race_time):
    race = {}
    for r in reindeers.keys():
        race[r] = {"state": "flying", "since": 0, "distance": 0, "score_part2": 0}
    for i in range(race_time):
        for name, data in race.items():
            data["since"] += 1
            match data["state"]:
                case "flying":
                    data["distance"] += reindeers[name]["speed"]
                    if data["since"] == reindeers[name]["endurance"]:
                        data["state"] = "resting"
                        data["since"] = 0
                case "resting":
                    if data["since"] == reindeers[name]["rest_time"]:
                        data["state"] = "flying"
                        data["since"] = 0
                case _:
                    print("ERROR")
        update_scores_part2(race)
    return max([r["distance"] for _ , r in race.items()]), max([r["score_part2"] for _ , r in race.items()])
    

#INPUT = f"{SCRIPT_DIR}/input_test.txt"
INPUT = f"{SCRIPT_DIR}/input.txt"
data = read_input(INPUT)
reindeers = parse(data)
result1, result2 = run(reindeers, 2503)
print("Result1 = ", result1)
print("Result2 = ", result2)
