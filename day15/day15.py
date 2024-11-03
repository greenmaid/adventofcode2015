#!/usr/bin/env python3

import os
from typing import List

SCRIPT_DIR = os.path.dirname(os.path.realpath(__file__))


def read_input(path: str) -> List[str]:
    with open(path, 'r') as f:
        lines = f.read().splitlines()
    return lines


def parse(data):
    ingredients ={}
    for d in data:
        match d.replace(",", "").replace(":", "").split(" "):
            case [name, 'capacity', capacity, 'durability', durability, 'flavor', flavor, 'texture', texture, 'calories', calories]:
                ingredients[name] = {
                    'capacity': int(capacity), 
                    'durability': int(durability), 
                    'flavor': int(flavor), 
                    'texture': int(texture), 
                    'calories': int(calories),
                }
    return ingredients


def count_score(recipe):
    score = 1
    for k, v in recipe.items():
        if k in ['capacity', 'durability', 'flavor', 'texture']:
            if v <= 0:
                return 0
            score *= v
    return score


def add_ingredient(ingredient, ingredients, recipe):
    for type, value in ingredients[ingredient].items():
        if not type in recipe:
            recipe[type] = 0
        recipe[type] += value
    if not ingredient in recipe:
        recipe[ingredient] = 0
    recipe[ingredient] += 1


def run1(ingredients):
    recipe = {}
    for ingredient, _ in ingredients.items():
        add_ingredient(ingredient, ingredients, recipe)
        
    for i in range(100 - len(ingredients)):
        i += 1 + len(ingredients)
        max = ("", -1)
        for ingredient, addings in ingredients.items():
            try_recipe = {k: v for k, v in recipe.items()}
            add_ingredient(ingredient, ingredients, try_recipe)
            try_score = count_score(try_recipe)
            if try_score > max[1]:
                max = (ingredient, try_score)
        add_ingredient(max[0], ingredients, recipe)
    return recipe, count_score(recipe)



#INPUT = f"{SCRIPT_DIR}/input_test.txt"
INPUT = f"{SCRIPT_DIR}/input.txt"
data = read_input(INPUT)
ingredients = parse(data)
recipe1, result1 = run1(ingredients)
print("Result1 = ", result1)

# =========================================

def remove_ingredient(ingredient, ingredients, recipe):
    for type, value in ingredients[ingredient].items():
        recipe[type] -= value
    recipe[ingredient] -= 1


def replace_ingredient(ingredient1, ingredient2, ingredients, recipe):
    if recipe[ingredient1] == 0:
        return False
    remove_ingredient(ingredient1, ingredients, recipe)
    add_ingredient(ingredient2, ingredients, recipe)


def run2(recipe, ingredients):
    ingredient_list = list(ingredients.keys())
    possible_pairs = [(a, b) for a in ingredient_list for b in ingredient_list if a != b]
    possible_double_pairs = [(a, b) for a in possible_pairs for b in possible_pairs]
    current_cal = recipe["calories"]
    while current_cal != 500:
        best = ([], 0, current_cal)
        for pair in possible_pairs:
            try_recipe = {k: v for k, v in recipe.items()}
            replace_ingredient(pair[0], pair[1], ingredients, try_recipe)
            try_cal = try_recipe["calories"]
            try_score = count_score(try_recipe)
            if (abs(try_cal-500) < abs(best[2]-500) and try_score >= best[1] and try_score > 0):
                best = (pair, try_score, try_cal)
        if best[0] != []:
            replace_ingredient(best[0][0], best[0][1], ingredients, recipe)

        else:
            for double_pair in possible_double_pairs:
                try_recipe = {k: v for k, v in recipe.items()}
                replace_ingredient(double_pair[0][0], double_pair[0][1], ingredients, try_recipe)
                replace_ingredient(double_pair[1][0], double_pair[1][1], ingredients, try_recipe)
                try_cal = try_recipe["calories"]
                try_score = count_score(try_recipe)
                if (abs(try_cal-500) <= abs(best[2]-500) and try_score >= best[1] and try_score > 0):
                    best = (double_pair, try_score, try_cal)
            replace_ingredient(best[0][0][0], best[0][0][1], ingredients, recipe) 
            replace_ingredient(best[0][1][0], best[0][1][1], ingredients, recipe)
        current_cal = recipe["calories"]

    while True:
        current_score = count_score(recipe)
        best = ([], current_score, current_cal)
        for pair in possible_pairs:
            try_recipe = {k: v for k, v in recipe.items()}
            replace_ingredient(pair[0], pair[1], ingredients, try_recipe)
            try_cal = try_recipe["calories"]
            try_score = count_score(try_recipe)
            #print(pair,"calories:", best[2], "->", try_cal, "| score: ", best[1], "->", try_score)
            if (abs(try_cal-500) == abs(best[2]-500) and try_score > best[1]):
                best = (pair, try_score, try_cal)
        if best[0] != []:
            replace_ingredient(best[0][0], best[0][1], ingredients, recipe)
        else:
            break

    return recipe, count_score(recipe)


recipe2, result2 = run2(recipe1, ingredients)
print("Result2 = ", result2)
